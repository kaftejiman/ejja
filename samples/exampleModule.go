package modules

import (
	"fmt"

	"github.com/kaftejiman/ejja/utils"
)

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
	fmt.Printf("Name: %s\n", "samplemodule")
	fmt.Printf("Usage: %s\n", "ejja --project=\"example/project\" --module=\"samplemodule\"")
	fmt.Printf("Description: %s\n", `Sample module: Prints number of statement in target function.`)
}

func (m *AnalyserModule) run(project string, functions ...string) {
	fmt.Print("Sample module is running with project path %s functions are : %s \n", project, functions)

	var collections []utils.StatementCollection
	var collectionElement utils.StatementCollection
	collections = utils.ParseFunctions(project, functions, true)

	for i := range collections {
		collectionElement = collections[i]
		fmt.Printf("\n[+] Number of statement is %d", len(collectionElement.Listing))
	}

	fmt.Println("\n[+] Done.")
}

func doSth() {

}
