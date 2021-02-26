package modules

import (
	"fmt"
	"go/ast"

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

	var collections []utils.StatementCollection
	var collectionElement utils.StatementCollection
	collections = utils.ParseFunctions(project, functions, true)

	//fmt.Println(collections)

	for i := range collections {
		collectionElement = collections[i]
		flattenCollection(collectionElement)
	}

}

type levels struct {
	variable []string
	label    []string
}

type breaks struct {
	level []string
	entry []string
}

type continues struct {
	level []string
	entry []string
}

func flattenCollection(collection utils.StatementCollection) {

	whileLabel := utils.UniqueID()
	switchVariable := utils.UniqueID()
	entry := utils.UniqueID()
	exit := utils.UniqueID()
	var levels levels
	var breaks breaks
	var continues continues

	varDeclarations := utils.ReturnAssignments(collection)
	// if len(collection.ReturnStack) != 0 {

	// }

	//resTransformedBlock := transformBlock(stmts, switchVariable, whileLabel, entry, exit)
	fmt.Printf(`
%s
var %s string
%s = %s
while(%s != %s){
	switch(%s){`,
		varDeclarations, switchVariable, switchVariable, entry, switchVariable, exit, switchVariable)

	levels.label = append(levels.label, whileLabel)
	levels.variable = append(levels.variable, switchVariable)

}

func transformBlock(stmts []ast.Stmt, switchVariable string, whileLabel string, entry string, exit string) string {

	return "something"
}

/*func findDecl(node ast.Node) {

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
}*/

/*
idea:
1- parse target function
2- populate collection of Stmt stacks
3- pretty print according to template (flattening) by popping from collection
*/
