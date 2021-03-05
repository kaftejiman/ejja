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
	levels.tabs = 1

	var breaks breaks
	var continues continues
	fmt.Printf("[+] Emitting body of the transformed function..\n\n")
	collection = utils.ReturnAssignments(collection)

	fmt.Printf("%svar %s string\n", utils.GetTabs(levels.tabs), switchVariable)
	fmt.Printf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, entry)
	fmt.Printf("%sfor %s != \"%s\" {\n", utils.GetTabs(levels.tabs), switchVariable, exit)
	levels.tabs++
	fmt.Printf("%sswitch(%s){\n", utils.GetTabs(levels.tabs), switchVariable)

	levels.label = append(levels.label, whileLabel)
	levels.variable = append(levels.variable, switchVariable)

	transformBlock(collection.Listing, entry, exit, &levels, &breaks, &continues)
	levels.label = append(levels.label[:0], levels.label[1:]...)
	fmt.Printf("%s}\n", utils.GetTabs(levels.tabs))
	levels.tabs--
	fmt.Printf("%s}\n", utils.GetTabs(levels.tabs))
	levels.tabs--

}

func transformBlock(stmts []ast.Stmt, entry string, exit string, levels *levels, breaks *breaks, continues *continues) string {

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
		case "IfStmt":
			transformIf(stmts[i], entry, partExit, *levels)
			break
		case "SwitchStmt":
			transformSwitch(stmts[i], entry, partExit, *levels)
			break
		case "ForStmt":
			transformFor(stmts[i], entry, partExit, *levels, *breaks, *continues)
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

	// emit transformed code
	fmt.Printf("%scase \"%s\":\n", utils.GetTabs(levels.tabs), entry)
	levels.tabs++
	fmt.Printf("%s%s\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt))
	fmt.Printf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, exit)
	fmt.Printf("%sbreak\n", utils.GetTabs(levels.tabs))
	levels.tabs = levels.tabs - 2

}

func transformIf(stmt ast.Stmt, entry string, exit string, levels levels) {

	currStmt := stmt.(*ast.IfStmt)
	switchVariable := levels.variable[0]
	thenEntry := utils.UniqueID()
	elseEntry := exit
	hasElse := false

	if currStmt.Else != nil {
		hasElse = true
	}

	if hasElse {
		elseEntry = utils.UniqueID()
	}

	// emit transformed code
	fmt.Printf("%scase \"%s\": \n", utils.GetTabs(levels.tabs), entry)
	levels.tabs++
	fmt.Printf("%sif (%s) {\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt.Cond))
	levels.tabs++
	fmt.Printf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, thenEntry)
	levels.tabs--
	fmt.Printf("%s}else{\n", utils.GetTabs(levels.tabs))
	levels.tabs++
	fmt.Printf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, elseEntry)
	levels.tabs--
	fmt.Printf("%s}\n", utils.GetTabs(levels.tabs))
	fmt.Printf("%sbreak\n", utils.GetTabs(levels.tabs))
	levels.tabs--
	transformBlock(currStmt.Body.List, thenEntry, exit, &levels, nil, nil)
	if hasElse {
		transformBlock(currStmt.Else.(*ast.BlockStmt).List, elseEntry, exit, &levels, nil, nil)
	}

}

func transformFor(stmt ast.Stmt, entry string, exit string, levels levels, breaks breaks, continues continues) {
	currStmt := stmt.(*ast.ForStmt)
	switchVariable := levels.variable[0]
	bodyEntry := utils.UniqueID()
	testEntry := utils.UniqueID()
	incEntry := utils.UniqueID()

	// emit transformed code
	// case entry
	fmt.Printf("%scase \"%s\":\n", utils.GetTabs(levels.tabs), entry)
	levels.tabs++
	fmt.Printf("%s%s\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt.Init))
	fmt.Printf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, testEntry)
	fmt.Printf("%sbreak\n", utils.GetTabs(levels.tabs))
	levels.tabs--
	// case test cond
	fmt.Printf("%scase \"%s\":\n", utils.GetTabs(levels.tabs), testEntry)
	levels.tabs++
	fmt.Printf("%sif %s {\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt.Cond))
	levels.tabs++
	fmt.Printf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, bodyEntry)
	levels.tabs--
	fmt.Printf("%s}else{\n", utils.GetTabs(levels.tabs))
	levels.tabs++
	fmt.Printf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, exit)
	levels.tabs--
	fmt.Printf("%s}\n", utils.GetTabs(levels.tabs))
	fmt.Printf("%sbreak\n", utils.GetTabs(levels.tabs))
	levels.tabs--
	// case incrementation
	fmt.Printf("%scase \"%s\":\n", utils.GetTabs(levels.tabs), incEntry)
	levels.tabs++
	fmt.Printf("%s%s\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt.Post))
	fmt.Printf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, testEntry)
	fmt.Printf("%sbreak\n", utils.GetTabs(levels.tabs))
	levels.tabs--

	breaks.entry = append(breaks.entry, exit)
	continues.entry = append(continues.entry, entry)
	transformBlock(currStmt.Body.List, bodyEntry, entry, &levels, &breaks, &continues)

}

func transformBranch(stmt ast.Stmt, entry string, exit string, levels levels) {
	//currStmt := stmt.(*ast.BranchStmt)
}

func transformRange(stmt ast.Stmt, entry string, exit string, levels levels) {
	//currStmt := stmt.(*ast.RangeStmt)
}

func transformSwitch(stmt ast.Stmt, entry string, exit string, levels levels) {
	//currStmt := stmt.(*ast.SwitchStmt)
}
