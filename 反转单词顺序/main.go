package main

import "fmt"

// ReverseSentence ..
func ReverseSentence(s string) string {
	i, j := 0, 0
	stack := make([]string, 0)
	for j < len(s) {
		for j < len(s) && i <= j {
			// 找到空格 将空格之前的字符作为单词入栈
			if s[j] == ' ' {
				stack = append(stack, string(s[i:j]))
				// 跳过空格
				i = j + 1
			}
			// 没有空格 到字符串末尾,将i到j的字符入栈
			if j == len(s)-1 {
				// 切片不会包含最后一个索引，所以j+1
				stack = append(stack, string(s[i:j+1]))
				i = j + 1
			}
			j++
		}
	}

	res := ""
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		if len(curr) > 0 {
			res += " " + curr
		}
		stack = stack[0 : len(stack)-1]
	}
	if len(res) > 1 {
		return res[1:]
	}
	return res
}

func main() {
	fmt.Println(ReverseSentence("asd dsa"))
}
