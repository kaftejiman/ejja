package modules

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/henrylee2cn/aster/aster"
	"github.com/kaftejiman/ejja/utils"
)

type summary struct {
	basicObj     int `default:"0"`
	arrayObj     int `default:"0"`
	sliceObj     int `default:"0"`
	structObj    int `default:"0"`
	pointerObj   int `default:"0"`
	tupleObj     int `default:"0"`
	signatureObj int `default:"0"`
	interfaceObj int `default:"0"`
	mapObj       int `default:"0"`
	chanObj      int `default:"0"`
}

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
	fmt.Printf("Name: %s\n", "analyser")
	fmt.Printf("Usage: %s\n", "ejja --project=\"example/project\" --module=\"analyser\"")
	fmt.Printf("Description: %s\n", `Runs an analysis on the target project's codebase, returns summary of object analysis.`)
}

func (m *AnalyserModule) run(project string, functions ...string) {
	//fmt.Print("analyser is running with project path ", project, " functions are : ", functions, "\n")
	generateSummary(project)
}

// https://stackoverflow.com/a/56778921/15186713
func setField(field reflect.Value, defaultVal string) error {

	if !field.CanSet() {
		return fmt.Errorf("Can't set value\n")
	}

	switch field.Kind() {

	case reflect.Int:
		if val, err := strconv.ParseInt(defaultVal, 10, 64); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}
	case reflect.String:
		field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
	}

	return nil
}

// Set https://stackoverflow.com/a/56778921/15186713
func Set(ptr interface{}, tag string) error {
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return fmt.Errorf("Not a pointer")
	}

	v := reflect.ValueOf(ptr).Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		if defaultVal := t.Field(i).Tag.Get(tag); defaultVal != "-" {
			if err := setField(v.Field(i), defaultVal); err != nil {
				return err
			}

		}
	}
	return nil
}

func generateSummary(project string) bool {
	return objectSummary(project)

}

func objectSummary(project string) bool {
	var myprogram *aster.Program
	var summ summary
	var err error
	Set(summ, "default")
	myprogram, err = utils.LoadDirs(project)

	myprogram.Inspect(func(f aster.Facade) bool {
		switch f.TypKind() {
		case aster.Basic:
			summ.basicObj++
			break
		case aster.Array:
			summ.arrayObj++
			break
		case aster.Slice:
			summ.sliceObj++
			break
		case aster.Struct:
			summ.structObj++
			break
		case aster.Pointer:
			summ.pointerObj++
			break
		case aster.Tuple:
			summ.tupleObj++
			break
		case aster.Signature:
			summ.signatureObj++
			break
		case aster.Interface:
			summ.interfaceObj++
			break
		case aster.Map:
			summ.mapObj++
			break
		case aster.Chan:
			summ.chanObj++
			break
		default:
			fmt.Println("error generate summary")
			break
		}
		return true
	})

	if err != nil {
		fmt.Println("error: generateSummary")
	}

	fmt.Printf(`
		=== Object Summary ===

	basic objects        : %d
	array objects        : %d
	slice objects        : %d
	struct objects       : %d
	pointer objects      : %d
	tuple objects        : %d
	signature objects    : %d
	interface objects    : %d
	map objects          : %d
	chan objects         : %d

	`, summ.basicObj, summ.arrayObj, summ.sliceObj,
		summ.structObj, summ.pointerObj, summ.tupleObj,
		summ.signatureObj, summ.interfaceObj, summ.mapObj,
		summ.chanObj)

	return true
}
