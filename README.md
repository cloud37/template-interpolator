# Template Interpolator

The Template Interpolator tool allows you to process text templates by using environment variables and custom delimiters, utilizing the `github.com/Masterminds/sprig/v3` library to enhance template functionality with additional functions. This tool is particularly useful in scenarios where dynamically generated text needs to be created based on environment variables.

## Installation and Building

This project uses a Makefile for simplifying the building process. To compile the binary for your operating system, use the following command:

```bash
make build
```

This will generate the executable `interpol` in the `./bin/darwin/` or `./bin/linux/` directory, depending on your operating system.

## Usage

After building the tool, you can use it to process template files. The basic syntax is as follows:

```bash
./bin/<your_os>/interpol [options] filename
```

Where `<your_os>` should be replaced with `darwin` or `linux`, depending on your operating system, and `filename` is the path to the template file you wish to process.

### Options

- `-brace` or `-b`: Use curly braces (`{{`, `}}`) as delimiters. This is the default setting and does not need to be explicitly specified.
- `-square` or `-s`: Use square brackets (`[[`, `]]`) as delimiters.

### Using the Makefile for Rendering

You can also use the Makefile to render templates directly. This is done by specifying the template file name after the `render-brace/` or `render-square/` targets:

- For brace delimiters: `make render-brace/env-brace.json`
- For square delimiters: `make render-square/env-square.json`

### Sprig Library Usage

The tool incorporates the `github.com/Masterminds/sprig/v3` library, providing a wide range of template functions. For example, to format a date within your template, you could use:

```plaintext
{{ now | date "Monday, Jan 2, 2006" }}
```

Refer to the [Sprig Documentation](http://masterminds.github.io/sprig/) for more information on available functions.

## Makefile Targets

The `Makefile` includes several targets for convenience:

- `build`: Compiles the tool for your OS.
- `clean`: Removes compiled binaries.
- `render-brace/%`: Renders a file with brace delimiters.
- `render-square/%`: Renders a file with square delimiters.

## Contributing

Contributions are welcome! Please create an issue or pull request if you find a bug, suggest an improvement, or want to add a new feature.

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.
