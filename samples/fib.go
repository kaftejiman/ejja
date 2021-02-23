package test

import (
	"fmt"
	"strconv"
)

// FibonacciRecursion recursive fibonacci
func FibonacciRecursion(n int) int {
	{
		if n <= 1 {
			return n
		}
		return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
	}

}

func main() {
	for i := 0; i <= 9; i++ {
		fmt.Print(strconv.Itoa(FibonacciRecursion(i)) + "")
	}
}

/*
target

int fac(int x) {
  int tmp ;
  unsigned long next ;
  next = 4;
  while (1) {
    switch (next) {
    case 4:
    if (x == 1) {
      next = 3;
    } else {
      next = 2;
    }
    break;
    case 3: return (1); break;
    case 2: tmp = fac(x - 1); next = 1; break;
    case 1: return (x * tmp); break;
    }
  }
}
*/
