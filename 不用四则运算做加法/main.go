package main

import "fmt"

// Add ..
func Add(num1 int, num2 int) int {
	sum, carry := 0, 0
	for num2 != 0 {
		sum = num1 ^ num2
		carry = (num1 & num2) << 1
		num1 = sum
		num2 = carry
	}
	return num1
}

func main() {
	fmt.Println(Add(1000, 200))
}
