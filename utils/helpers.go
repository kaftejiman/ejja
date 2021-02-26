package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/henrylee2cn/aster/aster"
	"github.com/rs/xid"
)

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

// Validate Validates given functions removes empty functions if found exits when no function given
// returns the clean list of functions
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

// UniqueID generates a unique identifier for variable assignments
func UniqueID() string {
	id := xid.New()
	return id.String()
}

// GetNodeType returns a the given node's type
func GetNodeType(node ast.Node) string {
	val := reflect.ValueOf(node).Elem()
	return val.Type().Name()
}

// FindFunctions returns a list of *ast.FuncDecl matching given functions in a given folder path
func FindFunctions(project string, functions []string) []*ast.FuncDecl {
	functions = Validate(functions)
	targetFuncs := []*ast.FuncDecl{}
	for i := range functions {
		fn := findFunction(project, functions[i])
		targetFuncs = append(targetFuncs, fn)
	}
	return targetFuncs
}

// findFunction finds a function in a given project folder, returns the node of the given function
func findFunction(project string, function string) *ast.FuncDecl {

	var out *ast.FuncDecl
	fset := token.NewFileSet()
	packages, _ := parser.ParseDir(fset, project, nil, parser.AllErrors)

	for i := range packages {

		for _, file := range packages[i].Files {
			ast.Inspect(file, func(n ast.Node) bool {

				fn, ok := n.(*ast.FuncDecl)
				if ok {
					if fn.Name.Name == function {
						// append
						out = fn
					}
				}
				return true
			})
		}
	}
	return out
}
