package main

import "fmt"

// 所以无论 \textit{numbers}[\textit{high}]numbers[high] 是不是最小值，都有一个它的「替代品」\textit{numbers}[\textit{pivot}]numbers[pivot]，因此我们可以忽略二分查找区间的右端点。

func minArray(rotateArray []int) int {
	n := len(rotateArray)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	for l <= r {
		mid := (r-l)/2 + l
		if rotateArray[mid] > rotateArray[r] {
			l = mid + 1
		} else if rotateArray[mid] == rotateArray[r] {
			r--
		} else {
			r = mid
		}
	}
	return rotateArray[l]
}

func minNumberInRotateArray(rotateArray []int) int {
	left := 0
	right := len(rotateArray) - 1
	for left < right {
		if rotateArray[left] < rotateArray[right] {
			return rotateArray[left]
		}
		mid := left + (right-left)>>1
		if rotateArray[mid] >= rotateArray[left] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return rotateArray[left]
}

func main() {
	fmt.Println(minNumberInRotateArray([]int{1, 0, 1, 1, 1}))
}
