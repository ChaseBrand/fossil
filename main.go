package main

import (
	"fmt"
	"os"

	"github.com/ChaseBrand/fossil/fossil"
)

// This example for fossil only works in the binary's working directory.
// This could be easily extended, but for the sake of simplicity for the assignment I decided to just run it like this.

func main() {

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting the working directory")
		os.Exit(1)
	}

	// Iterate all files and subdirectories in our working directory for .go files.
	files, err := fossil.CollectFiles(cwd)
	if err != nil {
		fmt.Println("error collecting files from the directory")
		os.Exit(1)
	}

	// Instantiate a "Fossil" type- the master structure that contains all of the data about the project we're indexing.
	var fsl fossil.Fossil

	for _, f := range files {
		// For each file earlier indexed, collect the package name of the file.
		pkgName, err := fossil.GetPackageName(f)
		if err != nil {
			fmt.Printf("error getting package name %v\n", err)
			continue
		}

		// Using the package name and the absolute location of the file, start adding files to the Fossil.
		err = fsl.AddFile(pkgName, f)
		if err != nil {
			fmt.Printf("error adding file to fossil : %v\n", err)
			continue
		}
	}

	// Example of querying for all files in the "fossil" package.
	fossilPkgFiles := fsl.QueryPackageName("fossil")
	// And again for the "main" package (this file).
	mainPkgFiles := fsl.QueryPackageName("main")

	// Example of querying for all files containing a function named "QueryFunctionName".
	fnNameQuery := fsl.QueryFunctionName("QueryFunctionName")
	// Again with "MAIN" as a parameter. The input is intentionally case-sensitive, so this will return nothing.
	mainNameQuery := fsl.QueryFunctionName("MAIN")

	fmt.Println(fossilPkgFiles)
	fmt.Println(mainPkgFiles)
	fmt.Println("---")
	fmt.Println(fnNameQuery)
	fmt.Println(mainNameQuery)

	// Return all stored files in an array of arrays, where each different package is in its own array.
	allFiles := fsl.QueryAll()
	fmt.Println("---")
	fmt.Println(allFiles)
}
