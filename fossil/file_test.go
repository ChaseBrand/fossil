package fossil_test

import (
	"fmt"
	"testing"

	"github.com/ChaseBrand/fossil/fossil"
)

func TestGetFunctions(t *testing.T) {
	t.Run("GetFunctions", func(t *testing.T) {
		thisFile := fossil.File{
			Package: "fossil_test",
			Name:    "file_test.go",
			Path:    "./file_test.go",
		}
		fns, err := thisFile.GetFunctions()
		if err != nil {
			fmt.Printf("GetFunctions failed: %v", err)
			t.Fail()
		}
		if len(fns) != 1 {
			fmt.Print("GetFunctions failed: returned incorrect data")
			t.Fail()
		}
	})
}
