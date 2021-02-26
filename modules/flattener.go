package modules

import (
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"os"

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
	targetFuncs := utils.FindFunctions(project, functions)

	for i := range targetFuncs {
		fmt.Printf("[+] Found function `%s` in `%s`, flattening..\n\n", targetFuncs[i].Name.Name, targetFuncs[i].Name.Name)
	}

}

func flattenBlock(node ast.Node) {

	whileLabel := utils.UniqueID()
	switchVariable := utils.UniqueID()
	entry := utils.UniqueID()
	exit := utils.UniqueID()

	resTransformedBlock := transformBlock(node, switchVariable, whileLabel, entry, exit)
	fmt.Printf(`
var %s string
%s = %s
while(%s != %s){
	switch(%s){
		%s
	}
\}
	
`, switchVariable, switchVariable, entry, switchVariable, exit, switchVariable, resTransformedBlock)
}

func transformBlock(node ast.Node, switchVariable string, whileLabel string, entry string, exit string) string {
	blockParts := returnBlocks(node)
	fset := token.NewFileSet()

	for i := range blockParts {
		printer.Fprint(os.Stdout, fset, blockParts[i])
	}
	return "something"
}

func returnBlocks(node ast.Node) []ast.DeclStmt {

	out := []ast.DeclStmt{}
	ast.Inspect(node, func(n ast.Node) bool {
		fset := token.NewFileSet()
		ret, ok := n.(*ast.DeclStmt)
		if ok {
			fmt.Println("\nfound block")
			printer.Fprint(os.Stdout, fset, ret)
			out = append(out, *ret)
			return true
		}

		return true
	})
	return out
}

func findDecl(node ast.Node) {

	ast.Inspect(node, func(n ast.Node) bool {
		// Find Return Statements
		fset := token.NewFileSet()
		ret, ok := n.(*ast.AssignStmt)
		if ok {
			//fmt.Printf("gendecl statement found on line %d:\n\t", fset.Position(ret.Pos()).Line)
			printer.Fprint(os.Stdout, fset, ret)
			return true
		}
		return true
	})
}

/*
idea:
1- parse target function
2- populate collection of Stmt stacks
3- pretty print according to template (flattening) by popping from collection
*/
