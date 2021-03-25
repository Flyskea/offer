package main

import "fmt"

func replaceSpace(s string) string {
	if len(s) == 0 {
		return s
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count++
		}
	}
	originCount := len(s)
	for i := 0; i < (count * 2); i++ {
		s = s + " "
	}
	currCount := len(s) - 1
	str := []rune(s)
	for i := originCount - 1; i >= 0; i-- {
		if s[i] == ' ' {
			str[currCount] = '0'
			currCount--
			str[currCount] = '2'
			currCount--
			str[currCount] = '%'
			currCount--
		} else {
			str[currCount] = str[i]
			currCount--
		}
	}
	return string(str)
}

func main() {
	fmt.Println(replaceSpace("s string"))
	s := "asd"
	fmt.Println(s[1])
}
