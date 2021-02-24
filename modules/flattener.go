package modules

import (
	"fmt"
	"go/ast"

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
	fmt.Printf("Description: %s\n", `Flattens the target function's control flow graph.
	More: http://ac.inf.elte.hu/Vol_030_2009/003.pdf
	`)
}

func (m *FlattenerModule) run(project string, functions ...string) {
	fmt.Println("[+] Running flattener..")
	functions = utils.Validate(functions)
	program, _ := utils.LoadDirs(project)

	for i := range functions {
		flatten(program, functions[i])
	}

}

func flatten(program *aster.Program, function string) {
	currFunction = function
	program.Inspect(func(fa aster.Facade) bool {

		if fa.Name() != currFunction {
			return true
		}
		fmt.Printf("[+] Found function `%s` in `%s`, flattening..\n", currFunction, fa.File().Filename)
		var rootNode ast.Node = fa.Node()
		fmt.Println(rootNode)
		//rootNode
		return true
	})
	//_ = program.Rewrite()
}

func flattenBlock() {

}

func transformBlock() {

}
