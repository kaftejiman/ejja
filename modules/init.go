package modules

import (
	"fmt"
	"os"
)

var module Module
var modules []string

func init() {
	modules = []string{
		"analyser",
		"flattener",
	}
}

// Module interface of modules
type Module interface {
	run(project string, functions ...string)
	manifest()
}

type template struct {
	implement
	name     string
	project  string
	function string
}

type implement interface {
	manifest()
	run(project string, functions ...string)
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func selection(name string) {

	switch name {
	case "analyser":
		module = newAnalyserModule()
		break
	case "flattener":
		module = newFlattenerModule()
		break
	default:
		fmt.Println("Error: module doesnt exist, please run `list` command for listing the available modules.")
		os.Exit(-1)
	}

}

// Run specified module with optional project argumnent
func Run(name string, project string, functions ...string) {

	selection(name)
	module.run(project, functions...)
}

// Manifest specified module
func Manifest(name string) {

	selection(name)
	module.manifest()
}

// List available modules
func List() {
	fmt.Println("_____________________________________________________________________")
	for i := range modules {
		Manifest(modules[i])
		fmt.Println("_____________________________________________________________________")
	}

}
