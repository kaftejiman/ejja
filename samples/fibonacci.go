package test

import "fmt"

// recursive fibonacci
func fibonacci(n int) int {
	{
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

}

func main() {

	// test sort
	a := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	fmt.Println(sort(a))

	// test if
	if 1 > 2 {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}

	// test fib
	fmt.Println(fibonacci(30))
}
