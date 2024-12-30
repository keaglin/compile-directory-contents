package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CompilerOptions struct {
	IncludePaths      []string
	ExcludePaths      []string
	IncludeExtensions []string
	ExcludeExtensions []string
	OutputFileName    string
	DirPrefix         string
	Debug             bool
}

func isMatchingExtension(filename string, extensions []string) bool {
	if len(extensions) == 0 {
		return true
	}
	ext := strings.ToLower(filepath.Ext(filename))
	for _, includeExt := range extensions {
		if ext == includeExt {
			return true
		}
	}
	return false
}

func isMatchingPath(path string, patterns []string) bool {
	if len(patterns) == 0 {
		return true
	}
	for _, pattern := range patterns {
		if matched, _ := filepath.Match(pattern, path); matched {
			return true
		}
	}
	return false
}

func compileDirectoryContent(dirPath string, options CompilerOptions) error {
	var compiledContent strings.Builder
	var filesFound int

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dirPath, path)
		if err != nil {
			return err
		}

		if options.Debug {
			fmt.Printf("Checking path: %s\n", relPath)
		}

		// Special handling for the root directory
		if relPath == "." {
			return nil // Always enter the root directory
		}

		// Check if the directory matches the prefix
		if info.IsDir() {
			if options.DirPrefix != "" {
				if strings.HasPrefix(relPath, options.DirPrefix) {
					if options.Debug {
						fmt.Printf("Matched directory: %s\n", relPath)
					}
					return nil // Continue into this directory
				}
				if options.Debug {
					fmt.Printf("Skipping directory: %s\n", relPath)
				}
				return filepath.SkipDir // Skip this directory
			}
		}

		// Check include/exclude paths
		if !isMatchingPath(relPath, options.IncludePaths) {
			if options.Debug {
				fmt.Printf("Skipping non-included path: %s\n", relPath)
			}
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if isMatchingPath(relPath, options.ExcludePaths) {
			if options.Debug {
				fmt.Printf("Skipping excluded path: %s\n", relPath)
			}
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if !info.IsDir() {
			// Check file extensions
			if !isMatchingExtension(info.Name(), options.IncludeExtensions) {
				if options.Debug {
					fmt.Printf("Skipping non-included file type: %s\n", relPath)
				}
				return nil
			}
			if isMatchingExtension(info.Name(), options.ExcludeExtensions) {
				if options.Debug {
					fmt.Printf("Skipping excluded file type: %s\n", relPath)
				}
				return nil
			}

			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			compiledContent.WriteString(fmt.Sprintf("\n--- %s start ---\n", relPath))
			compiledContent.Write(content)
			compiledContent.WriteString(fmt.Sprintf("\n--- %s end ---\n", relPath))
			filesFound++
			if options.Debug {
				fmt.Printf("Added file: %s\n", relPath)
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	if filesFound == 0 {
		return fmt.Errorf("no matching files found")
	}

	// Create parent directories if they don't exist
	outputDir := filepath.Dir(options.OutputFileName)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	return os.WriteFile(options.OutputFileName, []byte(compiledContent.String()), 0644)
}

func main() {
	// Define command-line flags
	dirPath := flag.String("dir", ".", "Directory path to compile (default: current directory)")
	outputFile := flag.String("output", "compiled_content.md", "Output file name")
	dirPrefix := flag.String("prefix", "", "Prefix for directories to include")
	includePaths := flag.String("include", "", "Comma-separated list of paths to include (supports * wildcard)")
	excludePaths := flag.String("exclude", "node_modules,.git", "Comma-separated list of paths to exclude (supports * wildcard)")
	includeExts := flag.String("include-ext", "", "Comma-separated list of file extensions to include")
	excludeExts := flag.String("exclude-ext", ".jpg,.jpeg,.png,.gif,.bmp,.tiff,.mp3,.mp4,.avi,.mov,.wmv,.flv,.wav,.webp,.webm", "Comma-separated list of file extensions to exclude")
	debug := flag.Bool("debug", false, "Enable debug logging")

	flag.Parse()

	options := CompilerOptions{
		IncludePaths:      splitAndTrim(*includePaths),
		ExcludePaths:      splitAndTrim(*excludePaths),
		IncludeExtensions: splitAndTrim(*includeExts),
		ExcludeExtensions: splitAndTrim(*excludeExts),
		OutputFileName:    *outputFile,
		DirPrefix:         *dirPrefix,
		Debug:             *debug,
	}

	// Get absolute path
	absPath, err := filepath.Abs(*dirPath)
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		return
	}

	if options.Debug {
		fmt.Printf("Starting compilation from: %s\n", absPath)
		fmt.Printf("Looking for directories with prefix: %s\n", options.DirPrefix)
		fmt.Printf("Including paths: %v\n", options.IncludePaths)
		fmt.Printf("Excluding paths: %v\n", options.ExcludePaths)
		fmt.Printf("Including file extensions: %v\n", options.IncludeExtensions)
		fmt.Printf("Excluding file extensions: %v\n", options.ExcludeExtensions)
	}

	err = compileDirectoryContent(absPath, options)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Content compiled into %s\n", options.OutputFileName)
}

func splitAndTrim(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}
	return parts
}
