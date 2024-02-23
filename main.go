package main

import (
	"flag"
	"fmt"
	"github.com/Masterminds/sprig/v3" // Third-party library to extend template functionalities
	log "github.com/sirupsen/logrus"  // Logging library
	"os"
	"strings"
	"text/template" // Standard library for template processing
)

// init function sets the initial configuration for the log package.
func init() {
	log.SetLevel(log.ErrorLevel) // Only log errors
}

func main() {
	// Define flags for delimiter types with both long and short versions
	var braceFlag bool
	var squareFlag bool
	flag.BoolVar(&braceFlag, "brace", false, "")
	flag.BoolVar(&braceFlag, "b", false, "Use brace delimiters. This is the default and does not need to be explicitly specified.")
	flag.BoolVar(&squareFlag, "square", false, "")
	flag.BoolVar(&squareFlag, "s", false, "Use square delimiters.")

	// Custom usage function for the flag package
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "  %s [options] filename\n\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Options:")
		// Manually print the options for more control over formatting
		fmt.Fprintln(flag.CommandLine.Output(), "  --brace -b\tUse brace delimiters. This is the default and does not need to be explicitly specified.")
		fmt.Fprintln(flag.CommandLine.Output(), "  --square -s\tUse square delimiters.")
		fmt.Fprintln(flag.CommandLine.Output(), "\nfilename: The path to the template file to be processed.")
	}

	// Parse the flags
	flag.Parse()

	// Check if a filename argument has been provided
	if flag.NArg() < 1 {
		flag.Usage() // Display usage information
		os.Exit(0)   // Exit without error
	}
	filename := flag.Arg(0)
	render(filename, braceFlag, squareFlag) // Call render function with filename and delimiter flags
}

func render(filename string, useBrace, useSquare bool) {
	content, err := os.ReadFile(filename) // Read file content
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filename, err) // Log and exit on read error
	}

	envMap, err := envToMap() // Convert environment variables to a map
	if err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err) // Log and exit on conversion error
	}

	// Initialize template with default delimiters and sprig functions
	tmpl := template.New("tmpl").Funcs(sprig.FuncMap())
	if useBrace {
		// Default delimiters are already "{{" and "}}", so no changes are required for BRACE
	} else if useSquare {
		tmpl.Delims("[[", "]]") // Set delimiters to square brackets if -s or --square flag is used
	}

	t, err := tmpl.Parse(string(content)) // Parse the file content as a template
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err) // Log and exit on parse error
	}

	if err = t.Execute(os.Stdout, envMap); err != nil { // Execute the template with environment variables
		log.Fatalf("Failed to execute template: %v", err) // Log and exit on execution error
	}
}

func envToMap() (map[string]string, error) {
	envMap := make(map[string]string)
	for _, v := range os.Environ() {
		split := strings.SplitN(v, "=", 2) // Split environment variable into key and value
		if len(split) != 2 {
			return nil, fmt.Errorf("invalid environment variable: %s", v) // Return an error for invalid variables
		}
		envMap[split[0]] = split[1] // Add to map
	}
	return envMap, nil
}
