package modules

import "fmt"

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
	fmt.Print("this is flattener manifestation\n")
}

func (m *FlattenerModule) run(project ...string) {
	fmt.Print("flattener is running with project path\n", project[0])
}
