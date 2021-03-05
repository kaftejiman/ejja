package main

import (
	"log"

	"github.com/docopt/docopt.go"
	modules "github.com/kaftejiman/ejja/modules"
	"github.com/kaftejiman/ejja/utils"
)

var arguments struct {
	Run       bool   `docopt:"run"`
	List      bool   `docopt:"list"`
	Project   string `docopt:"-p,--project"`
	Module    string `docopt:"-m,--module"`
	Functions string `docopt:"-f,--functions"`
	Rewrite   bool   `docopt:"-r,--rewrite"`
	Verbose   bool   `docopt:"--verbose"`
}

func init() {

	var (
		usage = `
Golang source code level obfuscator. 
Usage:
	ejja run --project <path> --module <module> [--functions <functions> --rewrite]
	ejja list
	ejja -h | --help
	ejja --version

	Options:
	run                                     Runs the selected module.
	list                                    Lists available modules.
	-p --project <path>                     Absolute path to your Golang project. 
	-m --module <module>                    The obfuscation module to apply on target project.
	-f --functions <functions>              Target function or functions for obfuscation. Optional.
	-r --rewrite                            Rewrite result to project codebase (careful, data gets changed)
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
		modules.Run(arguments.Module, arguments.Project, arguments.Rewrite, arguments.Functions)
		break
	}

}
