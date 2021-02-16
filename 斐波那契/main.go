package main

import "fmt"

func fibonacci(n uint64) int {
	a, b := 1, 1
	for i := 2; uint64(i) < n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(fibonacci(10))
}
