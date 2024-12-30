# Directory Content Compiler

## Overview

The Directory Content Compiler is a Go-based tool that walks through a directory structure, compiles the content of files into a single document, and provides various customization options. It's particularly useful for creating documentation, performing code reviews, or getting an overview of a project's structure and content.

## Features

- Compile content from multiple files into a single document
- Filter directories based on prefix
- Include or exclude specific paths (with wildcard support)
- Include or exclude specific file extensions
- Customizable output file name
- Command-line interface for easy use and integration
- Debug logging for troubleshooting

## Installation

### Prerequisites

Ensure you have Go installed on your system. If not, download and install it from [golang.org](https://golang.org/).

### Compiling the Executable

1. Clone this repository or download the `directory_compiler.go` file.

2. Navigate to the directory containing `directory_compiler.go`.

3. Compile the script into an executable:

   ```
   go build -o dircompile directory_compiler.go
   ```

   This will create an executable named `dircompile` (or `dircompile.exe` on Windows).

4. (Optional) Move the executable to a directory in your system's PATH to run it from anywhere.

## Usage

Run the compiler using the following command:

```
./dircompile [options]
```

### Options

- `-dir`: Specify the directory to compile (default: current directory)
- `-output`: Set the output file name (default: "compiled_content.md")
- `-prefix`: Filter directories by prefix
- `-include`: Comma-separated list of paths to include (supports * wildcard)
- `-exclude`: Comma-separated list of paths to exclude (supports * wildcard, default: "node_modules,.git")
- `-include-ext`: Comma-separated list of file extensions to include
- `-exclude-ext`: Comma-separated list of file extensions to exclude (default: common binary and multimedia extensions)
- `-debug`: Enable debug logging (default: false)

### Examples

1. Compile the current directory:
   ```
   ./dircompile
   ```

2. Compile a specific directory:
   ```
   ./dircompile -dir="/path/to/your/directory"
   ```

3. Compile only directories starting with "email-":
   ```
   ./dircompile -prefix="email-"
   ```

4. Include only specific paths:
   ```
   ./dircompile -include="*/src/*,*/lib/*"
   ```

5. Exclude specific paths:
   ```
   ./dircompile -exclude="*/test/*,*/vendor/*"
   ```

6. Include only specific file types:
   ```
   ./dircompile -include-ext=".go,.js,.ts"
   ```

7. Exclude specific file types:
   ```
   ./dircompile -exclude-ext=".jpg,.png,.pdf"
   ```

8. Combine multiple options:
   ```
   ./dircompile -dir="/path/to/your/directory" -prefix="email-" -include="*/src/*" -exclude="*/test/*" -include-ext=".go,.js" -exclude-ext=".jpg,.png" -output="email_services.md" -debug
   ```

## Output

The tool generates a single file (default: `compiled_content.md`) containing the content of all compiled files. Each file's content is wrapped with start and end markers that include the relative file path:

```
--- path/to/filename.ext start ---
...
file contents
...
--- path/to/filename.ext end ---
```

This format provides clear separation between files and makes it easy to identify the source of each piece of content.

## Debugging

If you encounter any issues or unexpected behavior, you can use the `-debug` flag to enable detailed logging of the compilation process. This will show you which paths are being checked, which are being included or excluded, and which files are being added to the compilation.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)
