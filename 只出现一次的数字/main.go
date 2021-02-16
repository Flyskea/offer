package main

import (
	"fmt"
	"sort"
	"time"
)

// FindNumsAppearOnce ...
func FindNumsAppearOnce(nums []int) []int {
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

func main() {
	s := FindNumsAppearOnce([]int{1, 2, 3, 4, 5, 1, 2, 7, 8, 9, 1, 7, 9, 10, 3})
	now := time.Now()
	sort.Ints(s)
	time.Sleep(1 * time.Second)
	end := time.Since(now)
	fmt.Println(end.Nanoseconds())
	fmt.Println(s)
}
