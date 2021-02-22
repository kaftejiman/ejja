package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {

	src := []byte(`package main
	var sum int = 3 + 2`)
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "demo", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Fprint(os.Stdout, fset, node, nil)
}
