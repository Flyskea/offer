package main

import "fmt"

func isNumeric(str string) bool {
	if str == "" {
		return false
	}
	if str[0] == '+' || str[0] == '-' {
		str = str[1:]
	}
	if str == "" {
		return false
	}
	numeric := true
	scanDigits(&str)
	if str != "" {
		if str[0] == '.' {
			str = str[1:]
			scanDigits(&str)
			if str == "" {
				return true
			}
			if str[0] == 'e' || str[0] == 'E' {
				numeric = isExponential(&str)
			}
		} else if str[0] == 'e' || str[0] == 'E' {
			numeric = isExponential(&str)
		} else {
			numeric = false
		}
	}
	return numeric && str == ""
}

func scanDigits(str *string) {
	for *str != "" && (*str)[0] >= '0' && (*str)[0] <= '9' {
		*str = (*str)[1:]
	}
}

func isExponential(str *string) bool {
	if !((*str)[0] == 'e' || (*str)[0] == 'E') {
		return false
	}

	(*str) = (*str)[1:]
	if *str == "" {
		return false
	}
	if (*str)[0] == '+' || (*str)[0] == '-' {
		*str = (*str)[1:]
	}
	if *str == "" {
		return false
	}
	scanDigits(str)
	return *str == ""
}

func main() {
	var char byte
	fmt.Println(isNumeric("3.14e2"))
	fmt.Println(char)
}
