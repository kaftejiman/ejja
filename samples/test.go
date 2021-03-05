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

// func sort(numbers []int) []int {

// 	var c119pgbm9cj20m0udp00 string
// 	i := len(numbers)
// 	j := 1
// 	intermediate := numbers[j]
// 	c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp0g"
// 	for c119pgbm9cj20m0udp00 != "c119pgbm9cj20m0udp10" {
// 		switch c119pgbm9cj20m0udp00 {
// 		case "c119pgbm9cj20m0udp0g":
// 			c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp2g"
// 			break
// 		case "c119pgbm9cj20m0udp2g":
// 			if i > 0 {
// 				c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp20"
// 			} else {
// 				c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp1g"
// 			}
// 			break
// 		case "c119pgbm9cj20m0udp30":
// 			i--
// 			c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp2g"
// 			break
// 		case "c119pgbm9cj20m0udp20":
// 			c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp40"
// 			break
// 		case "c119pgbm9cj20m0udp40":
// 			if j < i {
// 				c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp3g"
// 			} else {
// 				c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp0g"
// 			}
// 			break
// 		case "c119pgbm9cj20m0udp4g":
// 			j++
// 			c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp40"
// 			break
// 		case "c119pgbm9cj20m0udp3g":
// 			if numbers[j-1] > numbers[j] {
// 				c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp50"
// 			} else {
// 				c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp20"
// 			}
// 			break
// 		case "c119pgbm9cj20m0udp50":
// 			c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp5g"
// 			break
// 		case "c119pgbm9cj20m0udp5g":
// 			numbers[j] = numbers[j-1]
// 			c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp60"
// 			break
// 		case "c119pgbm9cj20m0udp60":
// 			numbers[j-1] = intermediate
// 			c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp20"
// 			break
// 		case "c119pgbm9cj20m0udp1g":
// 			c119pgbm9cj20m0udp00 = "c119pgbm9cj20m0udp10"
// 			break
// 		}
// 	}
// 	return numbers
// }

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
	if 1 > 2 {
		fmt.Println("no")
	} else {
		fmt.Println("yes")
	}

	// for statement
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// call expression
	fmt.Println(fibonacci(30))
}
