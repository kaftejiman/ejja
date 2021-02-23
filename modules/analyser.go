package modules

import "fmt"

// AnalyserModule structure
type AnalyserModule struct {
	*template
}

func newAnalyserModule() Module {
	analyser := &AnalyserModule{}
	template := newTemplate(analyser)
	analyser.template = template
	analyser.template.name = "analyser"
	return analyser
}

func (*AnalyserModule) manifest() {
	fmt.Print("this is analyser manifestation \n")
}

func (m *AnalyserModule) run(project string, functions ...string) {
	fmt.Print("analyser is running with project path ", project, "functions are : ", functions, "\n")
}
