package main

import "fmt"

func NumberOf1(n int32) int32 {
	count := 0
	for n != 0 {
		count++
		n &= (n - 1)
	}
	return int32(count)
}

func main() {
	fmt.Println(NumberOf1(-2147483648))
}
