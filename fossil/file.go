package fossil

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
)

type File struct {
	Package string
	Name    string
	Path    string
	Funcs   []Function
}

// Create a new File structure.
func newFile(pkg string, path string) (File, error) {
	f := File{
		Package: pkg,
		Name:    filepath.Base(path),
		Path:    path,
	}

	funcs, err := f.GetFunctions()

	if err != nil {
		return File{}, err
	}

	f.Funcs = funcs
	return f, nil
}

func (f *File) GetFunctions() ([]Function, error) {
	var funcs []Function
	fileSet := token.NewFileSet()

	astFile, err := parser.ParseFile(fileSet, f.Path, nil, 0)
	if err != nil {
		return []Function{}, err
	}

	for _, d := range astFile.Decls {
		if fn, isFunction := d.(*ast.FuncDecl); isFunction {
			funcs = append(funcs, Function{
				Name:       fn.Name.Name,
				LineNumber: fileSet.Position(fn.Pos()).Line,
			})
		}
	}
	return funcs, nil
}
