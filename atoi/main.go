package main

import "fmt"

// StrToInt string to int
func StrToInt(str string) int {
	if str == "" {
		return 0
	}
	i := 0
	for str[i] == ' ' {
		i++
	}
	flag := false
	if str[i] == '+' {
		i++
	} else if str[i] == '-' {
		flag = true
		i++
	}
	res := 0
	for ; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			res = (int(str[i]) - '0') + res*10
		} else {
			return 0
		}
	}
	if flag {
		res = 0 - res
	}
	return res
}

func main() {
	fmt.Println(StrToInt("  -900"))
}
