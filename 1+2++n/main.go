package main

import "fmt"

// Sum_Solution ..
func Sum_Solution(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		// &&短路特性，如果前面一项为false，就不执行后一项了
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

func main() {
	fmt.Println(Sum_Solution(1000))
}
