package another_test

import "fmt"

func testIf(a int) int {
	if a == 0 {
		return 1
	}
	return 10
}

func main() {
	fmt.Println(testIf(0))
}
