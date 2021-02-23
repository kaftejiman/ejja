package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/henrylee2cn/aster/aster"
)

func readBodyifFibonnaci(fa aster.Facade) bool {

	if fa.Name() != "FibonacciRecursion" {
		return true
	}
	body, _ := fa.Body()
	fmt.Println("name:", fa.Name(), "body:", body)

	//fa.CoverBody(body)

	/*for i := fa.NumFields() - 1; i >= 0; i-- {
		field := fa.Field(i)
		if !field.Exported() {
			continue
		}
		field.Tags().Set(&aster.Tag{
			Key:     "json",
			Name:    goutil.SnakeString(field.Name()),
			Options: []string{"omitempty"},
		})
	}*/

	return true
}

// LoadDirs parses the source code of Go files under the directories and loads a new program.
func LoadDirs(dirs ...string) (*aster.Program, error) {
	p := aster.NewProgram()
	for _, dir := range dirs {
		if !filepath.IsAbs(dir) {
			dir, _ = filepath.Abs(dir)
		}
		err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			if err != nil || !f.IsDir() {
				return nil
			}
			p.Import(path)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return p.Load()
}

func TestHelper(project string) {
	prog, _ := LoadDirs(project)
	prog.Inspect(readBodyifFibonnaci)

	_ = prog.Rewrite()

}
