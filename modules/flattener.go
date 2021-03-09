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

func (m *FlattenerModule) run(project string, rewrite bool, functions ...string) {
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

	out := ""
	whileLabel := utils.UniqueID()
	switchVariable := utils.UniqueID()
	entry := utils.UniqueID()
	exit := utils.UniqueID()
	var levels levels
	levels.tabs = 1

	var breaks breaks
	var continues continues
	fmt.Printf("\n[+] Emitting transformed function..\n\n")

	out = out + fmt.Sprintf("%s{\n", collection.FunctionSig)
	out = out + utils.ReturnAssignments(collection)
	out = out + fmt.Sprintf("%svar %s string\n", utils.GetTabs(levels.tabs), switchVariable)
	out = out + fmt.Sprintf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, entry)
	out = out + fmt.Sprintf("%sfor %s != \"%s\" {\n", utils.GetTabs(levels.tabs), switchVariable, exit)
	levels.tabs++
	out = out + fmt.Sprintf("%sswitch(%s){\n", utils.GetTabs(levels.tabs), switchVariable)

	levels.label = append(levels.label, whileLabel)
	levels.variable = append(levels.variable, switchVariable)

	out = out + transformBlock(collection.Listing, entry, exit, &levels, &breaks, &continues)
	levels.label = append(levels.label[:0], levels.label[1:]...)
	out = out + fmt.Sprintf("%s}\n", utils.GetTabs(levels.tabs))
	levels.tabs--
	out = out + fmt.Sprintf("%s}\n", utils.GetTabs(levels.tabs))
	levels.tabs--
	out = out + fmt.Sprintf("%s}\n", utils.GetTabs(levels.tabs))
	levels.tabs--
	fmt.Printf("%s", out)

}

func transformBlock(stmts []ast.Stmt, entry string, exit string, levels *levels, breaks *breaks, continues *continues) string {

	out := ""
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
		case "BlockStmt":
			out = out + transformBlock((stmts[i]).(*ast.BlockStmt).List, entry, partExit, levels, breaks, continues)
			break
		case "IfStmt":
			out = out + transformIf(stmts[i], entry, partExit, *levels)
			break
		case "SwitchStmt":
			out = out + transformSwitch(stmts[i], entry, partExit, *levels)
			break
		case "ForStmt":
			out = out + transformFor(stmts[i], entry, partExit, *levels, *breaks, *continues)
			break
		case "RangeStmt":
			out = out + transformRange(stmts[i], entry, partExit, *levels)
			break
		case "BranchStmt":
			out = out + transformBranch(stmts[i], entry, partExit, *levels)
			break
		case "ExprStmt":
			out = out + transformExpr(stmts[i], entry, partExit, *levels)
			break
		case "ReturnStmt":
			out = out + transformReturn(stmts[i], entry, partExit, *levels)
			break
		case "AssignStmt":
			out = out + transformAssignment(stmts[i], entry, partExit, *levels)
			break
		case "EmptyStmt":
			break
		default:
			fmt.Println("not implemented:")
			fmt.Printf(utils.FormatNode(stmts[i]) + "\n")
			fmt.Printf(utils.GetNodeType(stmts[i]) + "\n")
			break
		}
		entry = partExit

	}
	return out
}

func transformExpr(stmt ast.Stmt, entry string, exit string, levels levels) string {

	out := ""
	currStmt := stmt.(*ast.ExprStmt)
	switchVariable := levels.variable[0]

	// emit transformed code
	out = out + fmt.Sprintf("%scase \"%s\":\n", utils.GetTabs(levels.tabs), entry)
	levels.tabs++
	out = out + fmt.Sprintf("%s%s\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt))
	out = out + fmt.Sprintf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, exit)
	out = out + fmt.Sprintf("%sbreak\n", utils.GetTabs(levels.tabs))
	levels.tabs = levels.tabs - 2

	return out

}

// wrong check again
func transformIf(stmt ast.Stmt, entry string, exit string, levels levels) string {

	out := ""
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
	out = out + fmt.Sprintf("%scase \"%s\": \n", utils.GetTabs(levels.tabs), entry)
	levels.tabs++
	out = out + fmt.Sprintf("%sif (%s) {\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt.Cond))
	levels.tabs++
	out = out + fmt.Sprintf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, thenEntry)
	levels.tabs--
	out = out + fmt.Sprintf("%s}else{\n", utils.GetTabs(levels.tabs))
	levels.tabs++
	out = out + fmt.Sprintf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, elseEntry)
	levels.tabs--
	out = out + fmt.Sprintf("%s}\n", utils.GetTabs(levels.tabs))
	out = out + fmt.Sprintf("%sbreak\n", utils.GetTabs(levels.tabs))
	levels.tabs--
	out = out + transformBlock(currStmt.Body.List, thenEntry, exit, &levels, nil, nil)

	if hasElse {
		out = out + transformBlock(currStmt.Else.(*ast.BlockStmt).List, elseEntry, exit, &levels, nil, nil)
	}

	return out

}

func transformFor(stmt ast.Stmt, entry string, exit string, levels levels, breaks breaks, continues continues) string {

	out := ""
	currStmt := stmt.(*ast.ForStmt)
	switchVariable := levels.variable[0]
	bodyEntry := utils.UniqueID()
	testEntry := utils.UniqueID()
	incEntry := utils.UniqueID()
	specialEntry := entry

	// emit transformed code
	// case entry
	if utils.GetNodeType(currStmt.Init) != "EmptyStmt" {
		out = out + fmt.Sprintf("%scase \"%s\":\n", utils.GetTabs(levels.tabs), entry)
		levels.tabs++

		out = out + fmt.Sprintf("%s%s\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt.Init))
		out = out + fmt.Sprintf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, testEntry)
		out = out + fmt.Sprintf("%sbreak\n", utils.GetTabs(levels.tabs))
		levels.tabs--
		specialEntry = testEntry
	}

	// case test cond
	out = out + fmt.Sprintf("%scase \"%s\":\n", utils.GetTabs(levels.tabs), specialEntry)
	levels.tabs++
	out = out + fmt.Sprintf("%sif %s {\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt.Cond))
	levels.tabs++
	out = out + fmt.Sprintf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, bodyEntry)
	levels.tabs--
	out = out + fmt.Sprintf("%s}else{\n", utils.GetTabs(levels.tabs))
	levels.tabs++
	out = out + fmt.Sprintf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, exit)
	levels.tabs--
	out = out + fmt.Sprintf("%s}\n", utils.GetTabs(levels.tabs))
	out = out + fmt.Sprintf("%sbreak\n", utils.GetTabs(levels.tabs))
	levels.tabs--

	// case incrementation
	out = out + fmt.Sprintf("%scase \"%s\":\n", utils.GetTabs(levels.tabs), incEntry)
	levels.tabs++
	out = out + fmt.Sprintf("%s%s\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt.Post))
	out = out + fmt.Sprintf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, specialEntry)
	out = out + fmt.Sprintf("%sbreak\n", utils.GetTabs(levels.tabs))
	levels.tabs--

	breaks.entry = append(breaks.entry, exit)
	continues.entry = append(continues.entry, entry)
	out = out + transformBlock(currStmt.Body.List, bodyEntry, incEntry, &levels, &breaks, &continues)

	return out

}

func transformReturn(stmt ast.Stmt, entry string, exit string, levels levels) string {

	out := ""
	currStmt := stmt.(*ast.ReturnStmt)

	out = out + fmt.Sprintf("%scase \"%s\":\n", utils.GetTabs(levels.tabs), entry)
	levels.tabs++
	out = out + fmt.Sprintf("%s%s\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt))
	out = out + fmt.Sprintf("%sbreak\n", utils.GetTabs(levels.tabs))
	levels.tabs--

	return out

}

func transformAssignment(stmt ast.Stmt, entry string, exit string, levels levels) string {

	out := ""

	if utils.GetNodeType(stmt) == "EmptyStmt" {
		return out
	}

	currStmt := stmt.(*ast.AssignStmt)
	switchVariable := levels.variable[0]

	out = out + fmt.Sprintf("%scase \"%s\":\n", utils.GetTabs(levels.tabs), entry)
	levels.tabs++
	out = out + fmt.Sprintf("%s%s\n", utils.GetTabs(levels.tabs), utils.FormatNode(currStmt))
	out = out + fmt.Sprintf("%s%s = \"%s\"\n", utils.GetTabs(levels.tabs), switchVariable, exit)
	out = out + fmt.Sprintf("%sbreak\n", utils.GetTabs(levels.tabs))
	levels.tabs--

	return out
}

func transformRange(stmt ast.Stmt, entry string, exit string, levels levels) string {
	//currStmt := stmt.(*ast.RangeStmt)
	out := ""

	return out
}

func transformBranch(stmt ast.Stmt, entry string, exit string, levels levels) string {
	//currStmt := stmt.(*ast.BranchStmt)
	out := ""

	return out
}

func transformSwitch(stmt ast.Stmt, entry string, exit string, levels levels) string {
	//currStmt := stmt.(*ast.SwitchStmt)
	out := ""

	return out
}
