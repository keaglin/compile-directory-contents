package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type CompilerOptions struct {
	ExcludedPaths  []string
	OutputFileName string
	DirPrefix      string
	Debug          bool
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

		// Skip excluded paths
		for _, excludedPath := range options.ExcludedPaths {
			if strings.HasPrefix(relPath, excludedPath) {
				if info.IsDir() {
					if options.Debug {
						fmt.Printf("Skipping excluded directory: %s\n", relPath)
					}
					return filepath.SkipDir
				}
				if options.Debug {
					fmt.Printf("Skipping excluded file: %s\n", relPath)
				}
				return nil
			}
		}

		if !info.IsDir() {
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
		return fmt.Errorf("no matching files found in directories with prefix '%s'", options.DirPrefix)
	}

	return os.WriteFile(options.OutputFileName, []byte(compiledContent.String()), 0644)
}

func main() {
	// Define command-line flags
	dirPath := flag.String("dir", ".", "Directory path to compile (default: current directory)")
	outputFile := flag.String("output", "compiled_content.md", "Output file name")
	dirPrefix := flag.String("prefix", "", "Prefix for directories to include")
	excludePaths := flag.String("exclude", "node_modules,.git", "Comma-separated list of paths to exclude")
	debug := flag.Bool("debug", false, "Enable debug logging")

	flag.Parse()

	options := CompilerOptions{
		ExcludedPaths:  strings.Split(*excludePaths, ","),
		OutputFileName: *outputFile,
		DirPrefix:      *dirPrefix,
		Debug:          *debug,
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
	}

	err = compileDirectoryContent(absPath, options)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Content compiled into %s\n", options.OutputFileName)
}
