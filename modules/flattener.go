package modules

import (
	"fmt"
	"go/ast"
	"os"
	"strings"

	aster "github.com/henrylee2cn/aster/aster"
	"github.com/kaftejiman/ejja/utils"
)

var currFunction string

// FlattenerModule structure
type FlattenerModule struct {
	*template
}

func newFlattenerModule() Module {
	flattener := &FlattenerModule{}
	template := newTemplate(flattener)
	flattener.template = template
	flattener.template.name = "flattener"
	return flattener
}

func (*FlattenerModule) manifest() {
	fmt.Printf("Name: %s\n", "flattener")
	fmt.Printf("Usage: %s\n", "ejja --project=\"example/project\" --module=\"flattener\" --function=\"main\"")
	fmt.Printf("Description: %s\n", `Flattens the target function's control flow.
	More: http://ac.inf.elte.hu/Vol_030_2009/003.pdf
	`)
}

func (m *FlattenerModule) run(project string, functions ...string) {
	functions = validate(functions)
	fmt.Print("flattener is running with project path ", project, " functions are ", functions, "\n")
	program, _ := utils.LoadDirs(project)

	for i := range functions {
		flatten(program, functions[i])
	}

}

func validate(functions []string) []string {
	out := []string{}
	fn := []string{}
	for i := range functions {
		fn = strings.Split(functions[i], ",")
		for j := range fn {
			if fn[j] == "" {
				continue
			}
			out = append(out, fn[j])
		}
	}
	if len(out) == 0 {
		fmt.Println("Error: no functions given.")
		os.Exit(-1)
	}
	return out
}

func flatten(program *aster.Program, function string) {
	currFunction = function
	program.Inspect(func(fa aster.Facade) bool {

		if fa.Name() != currFunction {
			return true
		}

		var node ast.Node = fa.Node()
		fmt.Println(node)

		return true
	})
	//_ = program.Rewrite()
}

func flattenBlock() {

}

func transformBlock() {

}
