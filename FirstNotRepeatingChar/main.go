package main

import "fmt"

func FirstNotRepeatingChar(str string) int {
	if str == "" {
		return -1
	}
	chars := make([]int, 256)
	for _, v := range str {
		chars[v]++
	}
	length := len(str)
	for i := 0; i < length; i++ {
		if chars[str[i]] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(FirstNotRepeatingChar("google"))
}
