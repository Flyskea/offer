package main

import (
	"fmt"
)

func firstK(data []int, k, start, end int) int {
	if start > end {
		return -1
	}
	middleIndex := (start + end) / 2
	middleData := data[middleIndex]
	if middleData == k {
		if (middleIndex > 0 && data[middleIndex-1] != k) || middleIndex == 0 {
			return middleIndex
		} else {
			end = middleIndex - 1
		}
	} else if middleData > k {
		end = middleIndex - 1
	} else {
		start = middleIndex + 1
	}
	return firstK(data, k, start, end)
}

func lastK(data []int, k, start, end int) int {
	if start > end {
		return -1
	}
	middleIndex := (start + end) / 2
	middleData := data[middleIndex]
	if middleData == k {
		if (middleIndex < len(data)-1 && data[middleIndex+1] != k) || middleIndex == len(data)-1 {
			return middleIndex
		} else {
			start = middleIndex + 1
		}
	} else if middleData < k {
		start = middleIndex + 1
	} else {
		end = middleIndex - 1
	}
	return lastK(data, k, start, end)
}

// GetNumberOfK get k
func GetNumberOfK(data []int, k int) int {
	var number int
	if len(data) > 0 {
		first := firstK(data, k, 0, len(data)-1)
		last := lastK(data, k, 0, len(data)-1)
		fmt.Println(first, last)
		if first > -1 && last > -1 {
			number = last - first + 1
		}
	}
	return number
}

func main() {
	fmt.Println(GetNumberOfK([]int{1, 2, 3, 3, 3, 3, 4, 5}, 3))
}
