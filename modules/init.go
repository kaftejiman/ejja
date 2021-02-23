package modules

import "fmt"

// Module interface of modules
type Module interface {
	run(project ...string)
	manifest()
}

type template struct {
	implement
	name    string
	project string
}

type implement interface {
	manifest()
	run(project ...string)
}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

func (t *template) manifest() {
	fmt.Print("This is a manifestation of a module:\n")
}

func (t *template) run(project ...string) {
	t.project = project[0]
	fmt.Print("loaded module\n")
	t.implement.run()
	fmt.Print("finished running\n")
}

// Run specified module with optional project argumnent
func Run(name string, project ...string) {
	var module Module
	switch name {
	case "analyser":
		module = newAnalyserModule()
		module.run(project...)
		break
	case "flattener":
		module = newFlattenerModule()
		module.run(project...)
		break
	}

}

// Manifest specified module
func Manifest(name string) {
	var module Module
	switch name {
	case "analyser":
		module = newAnalyserModule()
		break
	case "flattener":
		module = newFlattenerModule()
		break
	}

	module.manifest()

}

// List available modules
func List() {
	fmt.Println("list of available modules")
}
