package main

import "sort"

// IsContinuous ..
func IsContinuous(numbers []int) bool {
	length := len(numbers)
	if numbers == nil || length == 0 {
		return false
	}
	sort.Ints(numbers)
	numOfZero, numOfGap := 0, 0
	for i := 0; i < length-1; i++ {
		if numbers[i] == 0 {
			numOfZero++
			continue
		}
		if numbers[i] == numbers[i+1] {
			return false
		}
		numOfGap += numbers[i+1] - numbers[i] - 1
	}
	if numOfGap > numOfZero {
		return false
	}
	return true
}
