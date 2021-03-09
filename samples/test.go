package main

import "fmt"

// recursive fibonacci
// bodyStatement
// ifStatement
// returnStatement

func fibonacci(n int) int {
	{
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

}

// nested for/ if / return
func sort(numbers []int) []int {

	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				intermediate := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = intermediate
			}
		}
	}
	return numbers
}

func main() {

	// variable assignment
	a := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}

	// call expression
	fmt.Println(sort(a))

	// if statement
	if len(a) >= 1 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	// for statement
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// call expression
	fmt.Println(fibonacci(30))
}
