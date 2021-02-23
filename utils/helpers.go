package utils

import (
	"fmt"

	"github.com/kaftejiman/ejja/aster"
)

func readBody(fa aster.Facade) bool {

	fmt.Println(fa.TypKind())
	/*if fa.Name() != "FibonacciRecursion" {
		return true
	}
	body, _ := fa.Body()
	fmt.Println("name:", fa.Name(), "body:", body)

	fa.CoverBody(body)

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

// TestWalker test
func TestWalker(project string) {
	prog, _ := aster.LoadDirs("C:\\Users\\kaftejiman\\obfuscators\\ejja\\samples")
	prog.Inspect(readBody)

	_ = prog.Rewrite()
	//prog.PrintResume()

}
