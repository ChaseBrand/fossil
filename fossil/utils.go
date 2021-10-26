package fossil

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Returns the package name of a provided file.
func GetPackageName(file string) (string, error) {
	fileSet := token.NewFileSet()

	// Parse the *.go source file, but stop after the package clause has been parsed.
	astFile, err := parser.ParseFile(fileSet, file, nil, parser.PackageClauseOnly)

	if err != nil {
		return "", err
	}

	if astFile.Name.Name == "" {
		return "", fmt.Errorf("failed to parse package name")
	}

	return astFile.Name.Name, nil
}

// Returns all *.go files (non-test) in a directory and all subdirectories.
func CollectFiles(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if !strings.Contains(path, ".go") || strings.Contains(path, "_test.go") {
			return nil
		}

		if err != nil {
			return err
		}
		absPath, err := filepath.Abs(path)

		if err != nil {
			return err
		}
		files = append(files, absPath)
		return nil
	})

	if err != nil {
		return []string{}, err
	}
	return files, nil
}
