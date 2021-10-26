package fossil

type Fossil struct {
	Files []File
}

// Add a File structure to the Fossil.
func (s *Fossil) AddFile(pkg string, path string) error {
	file, err := newFile(pkg, path)

	if err != nil {
		return err
	}

	s.Files = append(s.Files, file)
	return nil
}

// Query our memory structure by package name and return a list of files residing in the package.
func (s *Fossil) QueryPackageName(pkg string) []File {
	var files []File

	for _, f := range s.Files {
		if f.Package == pkg {
			files = append(files, f)
		}
	}

	return files
}

// Query our memory structure by function name and return a list of files containing a function with that name.
func (s *Fossil) QueryFunctionName(fn string) []File {
	var files []File

	for _, f := range s.Files {
		for _, x := range f.Funcs {
			if x.Name == fn {
				files = append(files, f)
			}
		}
	}

	return files
}

// Return an array of arrays where the files are sorted by package into specific arrays.
func (s *Fossil) QueryAll() [][]File {
	var files [][]File
	var pkgNames []string

	// Populate our package array so we have the package names indexed.
	for _, f := range s.Files {
		if len(pkgNames) == 0 {
			pkgNames = append(pkgNames, f.Package)
			continue
		}
		for _, p := range pkgNames {
			if p == f.Package {
				continue
			} else {
				pkgNames = append(pkgNames, f.Package)
			}
		}
	}

	for _, p := range pkgNames {
		var matches []File
		for _, f := range s.Files {
			if f.Package == p {
				matches = append(matches, f)
			}
		}
		files = append(files, matches)
	}
	return files
}
