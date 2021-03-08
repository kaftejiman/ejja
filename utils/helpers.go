package utils

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	goastutils "golang.org/x/tools/go/ast/astutil"

	"github.com/henrylee2cn/aster/aster"
	"github.com/rs/xid"
)

// StatementCollection is a collection of all statements in target function as stacks
type StatementCollection struct {
	FunctionSig     string
	Listing         []ast.Stmt
	FuncStack       []ast.FuncDecl
	AssignDeclStack []ast.Stmt
	AssignStack     []ast.Stmt
	ExprStack       []ast.Stmt
	IfStack         []ast.Stmt
	BadStack        []ast.Stmt
	DeclStack       []ast.Stmt
	EmptyStack      []ast.Stmt
	LabeledStack    []ast.Stmt
	SendStack       []ast.Stmt
	IncDecStack     []ast.Stmt
	GoStack         []ast.Stmt
	DeferStack      []ast.Stmt
	ReturnStack     []ast.Stmt
	BranchStack     []ast.Stmt
	BlockStack      []ast.Stmt
	SwitchStack     []ast.Stmt
	TypeSwitchStack []ast.Stmt
	CommStack       []ast.Stmt
	SelectStack     []ast.Stmt
	ForStack        []ast.Stmt
	RangeStack      []ast.Stmt
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
			fmt.Println("Error: ", err)
			return nil, err
		}
	}
	return p.Load()
}

// Validate given functions removes empty functions if found exits when no function given
// returns the clean list of functions
// TODO check for errornous lists
func Validate(functions []string) []string {
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

// UniqueID generates a unique identifier
func UniqueID() string {
	id := xid.New()
	return id.String()
}

// GetNodeType returns a the given node's type
func GetNodeType(node ast.Node) string {
	val := reflect.ValueOf(node).Elem()
	return val.Type().Name()
}

// FormatNode returs the node as a string
func FormatNode(node ast.Node) string {
	buf := new(bytes.Buffer)
	_ = format.Node(buf, token.NewFileSet(), node)
	return buf.String()
}

// GetTabs returns a string composing of wanted tabs
func GetTabs(index int) string {
	tabs := ""
	for i := 0; i < index; i++ {
		tabs = tabs + "\t"
	}
	return tabs
}

// recursiveDeclAssignments finds declarations or declarative assignment statements recursively.
// func recursiveDeclAssignments(function *ast.FuncDecl, collection StatementCollection) StatementCollection {

// 	ast.Inspect(function, func(n ast.Node) bool {
// 		assignment, ok := n.(*ast.AssignStmt)
// 		decl, okk := n.(*ast.DeclStmt)

// 		// if declarative statement
// 		if ok {
// 			if assignment.Tok.String() == ":=" {
// 				collection.AssignDeclStack = append(collection.AssignDeclStack, assignment)
// 			}

// 		}

// 		// if declaration
// 		if okk {
// 			collection.AssignDeclStack = append(collection.AssignDeclStack, decl)
// 		}

// 		return true
// 	})

// 	return collection
// }

// ReturnAssignments returns the assignment statements as a string
func ReturnAssignments(collection StatementCollection) string {

	out := ""
	for i := range collection.AssignDeclStack {
		out = out + fmt.Sprintf("%s"+FormatNode(collection.AssignDeclStack[i])+"\n", GetTabs(1))
	}

	return out
}

// calibrateFuncion removes assignment statements from the function node tree, returns calibrated function
func calibrateFuncion(function *ast.FuncDecl, collection StatementCollection) (*ast.FuncDecl, StatementCollection) {

	goastutils.Apply(function, func(cr *goastutils.Cursor) bool {

		assignment, ok := cr.Node().(*ast.AssignStmt)
		decl, okk := cr.Node().(*ast.DeclStmt)

		if okk {
			collection.AssignDeclStack = append(collection.AssignDeclStack, decl)
			cr.Delete()
		}

		if ok {
			if assignment.Tok.String() == ":=" {
				collection.AssignDeclStack = append(collection.AssignDeclStack, assignment)
				//cr.Delete()
				// TODO: fix me cant delete me when I have no parent
			}
		}
		return true
	}, nil)

	return function, collection
}

// formatSignature returns function's signature
func formatSignature(funcType *ast.FuncType, funcIdent *ast.Ident) string {
	out := ""
	ident := FormatNode(funcIdent)
	ftype := FormatNode(funcType)
	index := 4
	out = out + ftype[:index] + " " + ident + ftype[index:]
	return out
}

// FindFunctions returns a list of *ast.FuncDecl matching given functions in a given folder path
// use verbose flag for printing found functions files.
func FindFunctions(project string, functions []string, verbose bool) []*ast.FuncDecl {
	functions = Validate(functions)
	targetFuncs := []*ast.FuncDecl{}
	for i := range functions {
		fn := findFunction(project, functions[i], verbose)
		targetFuncs = append(targetFuncs, fn)
	}
	return targetFuncs
}

// findFunction finds a function in a given project folder, returns the node of the given function
func findFunction(project string, function string, verbose bool) *ast.FuncDecl {

	var out *ast.FuncDecl
	fset := token.NewFileSet()
	packages, _ := parser.ParseDir(fset, project, nil, parser.AllErrors)

	for i := range packages {

		for _, file := range packages[i].Files {
			ast.Inspect(file, func(n ast.Node) bool {
				fn, ok := n.(*ast.FuncDecl)
				if ok {
					if fn.Name.Name == function {
						if verbose {
							fmt.Printf("[+] Found function `%s` in `%s.go` ..\n", fn.Name.Name, file.Name.Name)
						}
						out = fn
					}
				}
				return true
			})
		}
	}
	return out
}

// ParseFunctions returns a list of collections of statments of the given functions if found.
// verbose prints found file names
// TODO: error handling
func ParseFunctions(project string, functions []string, verbose bool) []StatementCollection {
	targetFuncs := FindFunctions(project, functions, verbose)
	out := []StatementCollection{}
	for i := range targetFuncs {
		out = append(out, parseFunction(targetFuncs[i], formatSignature(targetFuncs[i].Type, targetFuncs[i].Name)))
	}
	return out
}

func returnListing(fn *ast.FuncDecl) []ast.Stmt {
	return fn.Body.List
}

// parseFunction fills the statements collection with the statements found on target function recursively.
func parseFunction(fn *ast.FuncDecl, signature string) StatementCollection {

	collection := StatementCollection{}
	collection.FunctionSig = signature
	var ok bool

	ast.Inspect(fn, func(n ast.Node) bool {

		_, ok = n.(*ast.AssignStmt)
		if ok {
			collection.AssignStack = append(collection.AssignStack, n.(*ast.AssignStmt))
			return true
		}

		_, ok = n.(*ast.ExprStmt)
		if ok {
			collection.AssignStack = append(collection.ExprStack, n.(*ast.ExprStmt))
			return true
		}

		_, ok = n.(*ast.IfStmt)
		if ok {
			collection.IfStack = append(collection.IfStack, n.(*ast.IfStmt))
			return true
		}

		_, ok = n.(*ast.BadStmt)
		if ok {
			collection.BadStack = append(collection.BadStack, n.(*ast.BadStmt))
			return true
		}

		_, ok = n.(*ast.DeclStmt)
		if ok {
			collection.DeclStack = append(collection.DeclStack, n.(*ast.DeclStmt))
			return true
		}

		_, ok = n.(*ast.EmptyStmt)
		if ok {
			collection.EmptyStack = append(collection.EmptyStack, n.(*ast.EmptyStmt))
			return true
		}

		_, ok = n.(*ast.LabeledStmt)
		if ok {
			collection.LabeledStack = append(collection.LabeledStack, n.(*ast.LabeledStmt))
			return true
		}

		_, ok = n.(*ast.SendStmt)
		if ok {
			collection.SendStack = append(collection.SendStack, n.(*ast.SendStmt))
			return true
		}

		_, ok = n.(*ast.IncDecStmt)
		if ok {
			collection.IncDecStack = append(collection.IncDecStack, n.(*ast.IncDecStmt))
			return true
		}

		_, ok = n.(*ast.GoStmt)
		if ok {
			collection.GoStack = append(collection.GoStack, n.(*ast.GoStmt))
			return true
		}

		_, ok = n.(*ast.DeferStmt)
		if ok {
			collection.DeferStack = append(collection.DeferStack, n.(*ast.DeferStmt))
			return true
		}

		_, ok = n.(*ast.ReturnStmt)
		if ok {
			collection.ReturnStack = append(collection.ReturnStack, n.(*ast.ReturnStmt))
			return true
		}

		_, ok = n.(*ast.BranchStmt)
		if ok {
			collection.BranchStack = append(collection.BranchStack, n.(*ast.BranchStmt))
			return true
		}

		_, ok = n.(*ast.BlockStmt)
		if ok {
			collection.BlockStack = append(collection.BlockStack, n.(*ast.BlockStmt))
			return true
		}

		_, ok = n.(*ast.SwitchStmt)
		if ok {
			collection.SwitchStack = append(collection.SwitchStack, n.(*ast.SwitchStmt))
			return true
		}

		_, ok = n.(*ast.TypeSwitchStmt)
		if ok {
			collection.TypeSwitchStack = append(collection.TypeSwitchStack, n.(*ast.TypeSwitchStmt))
			return true
		}

		_, ok = n.(*ast.SelectStmt)
		if ok {
			collection.SelectStack = append(collection.SelectStack, n.(*ast.SelectStmt))
			return true
		}

		_, ok = n.(*ast.ForStmt)
		if ok {
			collection.ForStack = append(collection.ForStack, n.(*ast.ForStmt))
			return true
		}

		_, ok = n.(*ast.RangeStmt)
		if ok {
			collection.RangeStack = append(collection.RangeStack, n.(*ast.RangeStmt))
			return true
		}

		return true
	})

	collection.Listing = returnListing(fn)
	fn, collection = calibrateFuncion(fn, collection)
	return collection
}
