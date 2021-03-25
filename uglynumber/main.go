package main

func GetUglyNumber_Solution(index int) int {
	if index <= 1 {
		return index
	}
	// write code here
	uglyNumbers, a, b, c := make([]int, index), 0, 0, 0
	for i := range uglyNumbers {
		uglyNumbers[i] = 1
	}
	for i := 1; i < index; i++ {
		n2, n3, n5 := uglyNumbers[a]*2, uglyNumbers[b]*3, uglyNumbers[c]*5
		uglyNumbers[i] = min(n2, n3, n5)
		if uglyNumbers[i] == n2 {
			a++
		}
		if uglyNumbers[i] == n3 {
			b++
		}
		if uglyNumbers[i] == n5 {
			c++
		}
	}
	return uglyNumbers[len(uglyNumbers)-1]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}
