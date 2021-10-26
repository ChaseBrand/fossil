package fossil_test

import (
	"fmt"
	"testing"

	"github.com/ChaseBrand/fossil/fossil"
)

func TestGetPackageName(t *testing.T) {
	t.Run("GetPackageName", func(t *testing.T) {
		pkg, err := fossil.GetPackageName("utils.go")
		if err != nil {
			fmt.Printf("GetPackageName failed: %v", err)
			t.Fail()
		}
		if pkg != "fossil" {
			fmt.Printf("GetPackageName failed: got %s, expected utils", pkg)
			t.Fail()
		}
	})
}

func TestCollectFiles(t *testing.T) {
	t.Run("CollectFiles", func(t *testing.T) {
		files, err := fossil.CollectFiles(".")
		if err != nil {
			fmt.Printf("CollectFiles failed: %v", err)
			t.Fail()
		}
		if len(files) == 0 {
			fmt.Print("CollectFiles failed: no files collected")
			t.Fail()
		}
	})
}
