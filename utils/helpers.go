package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/henrylee2cn/aster/aster"
)

// LoadDirs parses the source code of Go files under the directories and loads a new program.
func LoadDirs(dirs ...string) (*aster.Program, error) {
	p := aster.NewProgram()
	for _, dir := range dirs {
		if !filepath.IsAbs(dir) {
			dir, _ = filepath.Abs(dir)
		}
		err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
			if err != nil || !f.IsDir() {
				return nil
			}
			p.Import(path)
			return nil
		})
		if err != nil {
			fmt.Println("Error: ", err)
			return nil, err
		}
	}
	return p.Load()
}

// Validate Validates given functions removes empty functions if found exits when no function given
// returns the clean list of functions
func Validate(functions []string) []string {
	out := []string{}
	fn := []string{}
	for i := range functions {
		fn = strings.Split(functions[i], ",")
		for j := range fn {
			if fn[j] == "" {
				continue
			}
			out = append(out, fn[j])
		}
	}
	if len(out) == 0 {
		fmt.Println("Error: no functions given.")
		os.Exit(-1)
	}
	return out
}
