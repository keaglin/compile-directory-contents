# Directory Content Compiler

## Overview

The Directory Content Compiler is a Go-based tool that walks through a directory structure, compiles the content of files into a single document, and provides various customization options. It's particularly useful for creating documentation, performing code reviews, or getting an overview of a project's structure and content.

## Features

- Compile content from multiple files into a single document
- Filter directories based on prefix
- Exclude specific paths
- Respect .gitignore rules
- Customizable output file name
- Command-line interface for easy use and integration
- File content breakpoints for improved readability

## Installation

### Prerequisites

Ensure you have Go installed on your system. If not, download and install it from [golang.org](https://golang.org/).

### Installing Dependencies

Install the required dependency:

```
go get github.com/sabhiram/go-gitignore
```

### Compiling the Executable

1. Clone this repository or download the `directory_compiler.go` file.

2. Navigate to the directory containing `directory_compiler.go`.

3. Compile the script into an executable:

   ```
   go build -o dircompile directory_compiler.go
   ```

   This will create an executable named `dircompile` (or `dircompile.exe` on Windows).

4. (Optional) Move the executable to a directory in your system's PATH to run it from anywhere. For example:

   On Linux/macOS:
   ```
   sudo mv dircompile /usr/local/bin/
   ```

   On Windows, you can create a new directory for your custom executables, add it to your PATH, and move the executable there.

## Usage

If you've added the executable to your PATH, you can run it from anywhere:

```
dircompile [options]
```

Otherwise, run it from the directory where it's located:

```
./dircompile [options]
```

### Options

- `-dir`: Specify the directory to compile (default: current directory)
- `-output`: Set the output file name (default: "compiled_content.md")
- `-prefix`: Filter directories by prefix
- `-exclude`: Comma-separated list of paths to exclude (default: "node_modules")
- `-gitignore`: Use .gitignore rules (default: true)

### Examples

1. Compile the current directory:
   ```
   dircompile
   ```

2. Compile a specific directory:
   ```
   dircompile -dir="/path/to/your/directory"
   ```

3. Compile only directories starting with "email-":
   ```
   dircompile -prefix="email-"
   ```

4. Specify a custom output file:
   ```
   dircompile -output="my_compilation.txt"
   ```

5. Exclude specific directories:
   ```
   dircompile -exclude="node_modules,.git,vendor"
   ```

6. Disable .gitignore rules:
   ```
   dircompile -gitignore=false
   ```

7. Combine multiple options:
   ```
   dircompile -dir="/path/to/your/directory" -prefix="email-" -output="email_services.md" -exclude="node_modules,.git,vendor" -gitignore=true
   ```

## Output

The tool generates a single file (default: `compiled_content.md`) containing the content of all compiled files. Each file's content is preceded by a header indicating its relative path within the compiled directory.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)
