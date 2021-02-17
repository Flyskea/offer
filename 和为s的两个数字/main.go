package main

func FindNumbersWithSum(array []int, sum int) []int {
	if array == nil || len(array) < 2 {
		return nil
	}
	ahead := len(array) - 1
	behind := 0
	for ahead > behind {
		// 当前总和
		cursum := array[ahead] + array[behind]
		if cursum == sum {
			return []int{array[behind], array[ahead]}
		} else if cursum > sum {
			ahead--
		} else {
			behind++
		}
	}
	return nil
}
