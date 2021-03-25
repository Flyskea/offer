package main

func InversePairs(data []int) int {
	if len(data) <= 1 {
		return 0
	}
	tmpData := make([]int, len(data))
	return countPairs(data, tmpData, 0, len(data)-1) % 1000000007
}

func countPairs(data, tmpData []int, left, right int) int {
	if left >= right {
		return 0
	}

	mid := left + (right-left)/2
	leftCount := countPairs(data, tmpData, left, mid)
	rightCount := countPairs(data, tmpData, mid+1, right)
	mergeCount := mergeAndCount(data, tmpData, left, right)

	return leftCount + rightCount + mergeCount
}

func mergeAndCount(data, tmpData []int, left, right int) int {
	if left >= right {
		return 0
	}

	var mergeCount int

	tmpIndex := left
	mid := left + (right-left)/2
	start1, start2 := left, mid+1
	for start1 <= mid && start2 <= right {
		if data[start2] <= data[start1] {
			mergeCount += mid - start1 + 1

			tmpData[tmpIndex] = data[start2]
			start2++
		} else {
			tmpData[tmpIndex] = data[start1]
			start1++
		}
		tmpIndex++
	}
	for start1 <= mid {
		tmpData[tmpIndex] = data[start1]
		start1++
		tmpIndex++
	}
	for start2 <= right {
		tmpData[tmpIndex] = data[start2]
		start2++
		tmpIndex++
	}

	for i := left; i <= right; i++ {
		data[i] = tmpData[i]
	}

	return mergeCount
}
