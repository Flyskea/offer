package main

import "strings"

// Reverse 反转字符串
func Reverse(str []rune) {
	i, j := 0, len(str)-1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
}

// LeftRotateString1 反转三次
func LeftRotateString1(str string, n int) string {
	if str == "" {
		return ""
	}
	if n == 0 {
		return str
	}
	s := []rune(str)
	Reverse(s[:n])
	Reverse(s[n:])
	Reverse(s)
	return string(s)
}

// LeftRotateString 字符串遍历拼接
func LeftRotateString(str string, n int) string {
	res := strings.Builder{}
	length := len(str) + n
	/*
	* 先将n以及后面的字符加入，随后将n前面的字符加入 完成左移
	* %取余 简化过程
	 */
	for i := n; i < length; i++ {
		res.WriteByte(str[i%len(str)])
	}
	return res.String()
}
