package main

import "fmt"

// FindContinuousSequence ..
func FindContinuousSequence(sum int) [][]int {
	if sum < 3 {
		return nil
	}
	res := make([][]int, 0)
	small := 1
	big := 2
	// 如果small大于middle，已经不可能有连续序列
	middle := (1 + sum) / 2
	// 当前总和
	cursum := small + big
	for small < middle {
		if cursum == sum {
			ans := []int{}
			for i := small; i <= big; i++ {
				ans = append(ans, i)
			}
			res = append(res, ans)
		}
		for cursum > sum && small < middle {
			cursum -= small
			small++
			if cursum == sum {
				ans := []int{}
				for i := small; i <= big; i++ {
					ans = append(ans, i)
				}
				res = append(res, ans)
			}
		}
		big++
		cursum += big
	}
	return res
}

func main() {
	fmt.Println(FindContinuousSequence(9))
}
