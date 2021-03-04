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

	for i := range collections {
		collectionElement = collections[i]
		flattenCollection(collectionElement)
	}

	fmt.Println("\n[+] Done.")

}

type levels struct {
	variable []string
	label    []string
	tabs     int
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
	//var breaks breaks
	//var continues continues
	collection = utils.ReturnAssignments(collection)

	fmt.Printf(`
var %s string
%s = "%s"
for %s != "%s" {
	switch(%s){`,
		switchVariable, switchVariable, entry, switchVariable, exit, switchVariable)
	fmt.Println("")
	levels.label = append(levels.label, whileLabel)
	levels.variable = append(levels.variable, switchVariable)
	levels.tabs = 1
	transformBlock(collection.Listing, entry, exit, &levels)
	levels.label = append(levels.label[:0], levels.label[1:]...)
	fmt.Println("\n\t}")
	fmt.Println("\n}")

}

func transformBlock(stmts []ast.Stmt, entry string, exit string, levels *levels) string {

	for i := range stmts {

		partExit := ""
		// exit setup
		if stmts[i] == stmts[len(stmts)-1] {
			// if last element
			partExit = exit
		} else {
			partExit = utils.UniqueID()
		}

		switch utils.GetNodeType(stmts[i]) {
		// case "Block":
		// 	transformBlock(stmts[i], entry, partExit)
		// 	break
		case "IfStmt":
			transformIf(stmts[i], entry, partExit, *levels)
			break
		case "SwitchStmt":
			transformSwitch(stmts[i], entry, partExit, *levels)
			break
		case "ForStmt":
			transformFor(stmts[i], entry, partExit, *levels)
			break
		case "RangeStmt":
			transformRange(stmts[i], entry, partExit, *levels)
			break
		case "BranchStmt":
			transformBranch(stmts[i], entry, partExit, *levels)
			break
		case "ExprStmt":
			transformExpr(stmts[i], entry, partExit, *levels)
			break
		default:
			fmt.Println("not implemented:")
			fmt.Printf(utils.FormatNode(stmts[i]) + "\n")
			break
		}
		entry = partExit

	}
	return "something"
}

func transformExpr(stmt ast.Stmt, entry string, exit string, levels levels) {

	currStmt := stmt.(*ast.ExprStmt)
	switchVariable := levels.variable[0]
	// setup tabs
	tabs := ""
	for i := 0; i < levels.tabs; i++ {
		tabs = tabs + "\t"
	}

	// emit transformed code
	fmt.Printf("%scase \"%s\": \n", tabs, entry)
	fmt.Printf("%s\t %s\n", tabs, utils.FormatNode(currStmt))
	fmt.Printf("%s\t %s = \"%s\" \n", tabs, switchVariable, exit)
	fmt.Printf("%s\t break\n", tabs)

}

func transformIf(stmt ast.Stmt, entry string, exit string, levels levels) {

	currStmt := stmt.(*ast.IfStmt)
	switchVariable := levels.variable[0]
	thenEntry := utils.UniqueID()
	elseEntry := exit
	hasElse := false

	// setup tabs
	tabs := ""
	for i := 0; i < levels.tabs; i++ {
		tabs = tabs + "\t"
	}

	if currStmt.Else != nil {
		hasElse = true
	}

	if hasElse {
		elseEntry = utils.UniqueID()
	}

	// emit transformed code
	fmt.Printf("%scase \"%s\": \n", tabs, entry)
	fmt.Printf("%sif (%s) {\n", tabs, utils.FormatNode(currStmt.Cond))
	fmt.Printf("%s\t %s = \"%s\"\n", tabs, switchVariable, thenEntry)
	fmt.Printf("%s}else{\n", tabs)
	fmt.Printf("%s\t %s = \"%s\"\n", tabs, switchVariable, elseEntry)
	fmt.Printf("%s}\n", tabs)
	fmt.Printf("%sbreak\n", tabs)
	transformBlock(currStmt.Body.List, thenEntry, exit, &levels)
	if hasElse {
		levels.tabs++
		transformBlock(currStmt.Else.(*ast.BlockStmt).List, elseEntry, exit, &levels)
	}

}

func transformBranch(stmt ast.Stmt, entry string, exit string, levels levels) {
	//currStmt := stmt.(*ast.BranchStmt)
}

func transformFor(stmt ast.Stmt, entry string, exit string, levels levels) {
	//currStmt := stmt.(*ast.ForStmt)
}

func transformRange(stmt ast.Stmt, entry string, exit string, levels levels) {
	//currStmt := stmt.(*ast.RangeStmt)
}

func transformSwitch(stmt ast.Stmt, entry string, exit string, levels levels) {
	//currStmt := stmt.(*ast.SwitchStmt)
}
