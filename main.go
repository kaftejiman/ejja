package main

import (
	"fmt"
	"log"

	"github.com/docopt/docopt.go"
	modules "github.com/kaftejiman/ejja/modules"
	"github.com/kaftejiman/ejja/utils"
)

var arguments struct {
	Run      bool     `docopt:"run"`
	List     bool     `docopt:"list"`
	Project  string   `docopt:"-p,--project"`
	Module   string   `docopt:"-m,--module"`
	Function []string `docopt:"-f,--function"`
	Verbose  bool     `docopt:"--verbose"`
}

func init() {

	var (
		usage = `
Golang source code level obfuscator. 
Usage:
	ejja run --project <path> [--module <module> [--function <functions>]]
	ejja list
	ejja -h | --help
	ejja --version

	Options:
	run                                     Runs the selected module.
	list                                    Lists available modules.
	-p --project <path>                     Absolute path to your Golang project. 
	-m --module <module>                    The obfuscation module to apply on target project.
	-f --function <function>                Target function or functions for obfuscation. Optional.
	-h --help                               Shows this screen.
	--verbose                               Shows details.
	--version                               Shows version.`
	)

	opts, err := docopt.ParseArgs(usage, nil, utils.Version)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = opts.Bind(&arguments)
	if err != nil {
		log.Fatalln(err)
		return
	}

}

func main() {

	switch {
	case arguments.List:
		modules.List()
		break
	case arguments.Run:
		fmt.Println("project:", arguments.Project, " module:", arguments.Module, " function:", arguments.Function)
		modules.Run(arguments.Module, arguments.Project)
		break
	}

	fmt.Println("main finished")
	/*
		src, err := ioutil.ReadFile("samples/fib.go")
		fmt.Print(string(src))
		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, "demo", src, parser.ParseComments)
		var out ast.Node
		if err != nil {
			panic(err)
		}

		visitorFunc := func(n ast.Node) bool {
			funcDecl, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}
			out = funcDecl.Body
			fmt.Printf("Found func at: %d:%d\n",
				fset.Position(funcDecl.Pos()).Line,
				fset.Position(funcDecl.Pos()).Column,
			)
			return true
		}

		ast.Inspect(node, visitorFunc)
		ast.Fprint(os.Stdout, fset, out, nil)
		//printer.Fprint(os.Stdout, fset, node)*/

}
