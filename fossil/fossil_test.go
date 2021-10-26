package fossil_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/ChaseBrand/fossil/fossil"
)

// Will construct a small Fossil type manually to test adding and querying. Seems to be the easiest way to do this for the time being.
var fsl fossil.Fossil

func TestAddFile(t *testing.T) {
	t.Run("AddFile", func(t *testing.T) {
		path, _ := filepath.Abs("./fossil_test.go")
		err := fsl.AddFile("fossil_test", path)
		if err != nil {
			fmt.Printf("AddFile failed: %v", err)
			t.Fail()
		}

		for _, f := range fsl.Files {
			f.GetFunctions()
		}
	})
}

func TestQueryPackageName(t *testing.T) {
	t.Run("QueryPackageName", func(t *testing.T) {
		files := fsl.QueryPackageName("fossil_test")
		if len(files) != 1 {
			fmt.Printf("QueryPackageName failed: incorrect data returned")
			t.Fail()
		}
	})
}

func TestQueryFunctionName(t *testing.T) {
	t.Run("QueryFunctionName", func(t *testing.T) {
		files := fsl.QueryFunctionName("TestQueryFunctionName")
		if len(files) != 1 {
			fmt.Printf("QueryFunctionName failed: incorrect data returned")
			t.Fail()
		}
	})
}

func TestQueryAll(t *testing.T) {
	t.Run("QueryAll", func(t *testing.T) {
		allData := fsl.QueryAll()
		if len(allData) != 1 {
			fmt.Printf("QueryAll failed: incorrect data returned")
			t.Fail()
		}
	})
}
