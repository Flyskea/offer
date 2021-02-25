package main

func duplicate(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	for _, v := range nums {
		if v < 0 || v > len(nums)-1 {
			return -1
		}
	}
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			tmp := nums[i]
			nums[i] = nums[tmp]
			nums[tmp] = tmp
		}
	}
	return -1
}
