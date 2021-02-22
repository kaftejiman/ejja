package main

import (
	"fmt"
	"log"

	"github.com/docopt/docopt.go"
)

var arguments struct {
	Obfuscate bool     `docopt:"obfuscate"`
	Analyse   bool     `docopt:"analyse"`
	List      bool     `docopt:"list"`
	Project   string   `docopt:"-p,--project"`
	Module    string   `docopt:"-m,--module"`
	Function  []string `docopt:"-f,--function"`
	Verbose   bool     `docopt:"--verbose"`
}

func init() {

	var (
		usage = `

	Golang source code level obfuscator. 
	Usage:
	ejja obfuscate  --project <path_to_project> [--module <obfuscation_module_name>] [--function <target_function>]
	ejja analyse    --project <path_to_project>
	ejja list
	ejja -h | --help
	ejja --version

	Options:
	obfuscate                                   Applies the selected obfuscation module.
	analyse                                     Runs project wide source code level analysis.
	list                                        Lists available obfuscation modules.
	-p --project <path_to_project>              Absolute path to your Golang project. 
	-m --module <obfuscation_module_name>       The obfuscation module to apply on target project.
	-f --function <target_function>             Target function/functions for obfuscation. Optional.
	-h --help                                   Shows this screen.
	--verbose                                   Shows details.
	--version                                   Shows version.`
	)

	opts, err := docopt.ParseArgs(usage, nil, "0.1.0")
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
	case arguments.Analyse:
		fmt.Println("project: ", arguments.Project)
		// delegate to project analyser

	case arguments.List:
		fmt.Println("some modules..")
		// list obfuscation modules available

	case arguments.Obfuscate:
		fmt.Println("project:", arguments.Project, " module: ", arguments.Module, " function: ", arguments.Function)
		// delegate to corresponding obfuscation module with optional args

	}

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
