package main

import "strconv"

// 哈希

// FindNumsAppearOnce1 ..
func FindNumsAppearOnce1(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	hash := map[int]int{}
	for _, v := range nums {
		if _, ok := hash[v]; !ok {
			hash[v] = 1
		} else {
			hash[v]++
		}
	}
	res := []int{}
	for k, v := range hash {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

// 位运算

// IsBits verify whether the bit is 1
func IsBits(num int, indexBit uint) bool {
	num >>= int(indexBit)
	return (num&1 == 1)
}

// FindFirstBitIsOne find the first bit which is 1
func FindFirstBitIsOne(num int) uint {
	indexBit := 0
	for (num&1 == 0) && (indexBit < strconv.IntSize) {
		num >>= 1
		indexBit++
	}
	return uint(indexBit)
}

// FindNumsAppearOnce ..
func FindNumsAppearOnce(nums []int) []int {
	length := len(nums)
	if nums == nil || length < 2 {
		return nil
	}
	resultExclusiveOR := 0
	for i := 0; i < length; i++ {
		resultExclusiveOR ^= nums[i]
	}
	indexOfOne := FindFirstBitIsOne(resultExclusiveOR)
	num1, num2 := 0, 0
	for i := 0; i < length; i++ {
		if IsBits(nums[i], indexOfOne) {
			num1 ^= nums[i]
		} else {
			num2 ^= nums[i]
		}
	}
	return []int{num1, num2}
}
