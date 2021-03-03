package main

/**
 *
 * @param num int整型一维数组
 * @param size int整型
 * @return int整型一维数组
 */
// 双端队列
func maxInWindows(num []int, size int) []int {
	var res []int
	if len(num) < size || size < 1 {
		return nil
	}
	index := []int{}
	for i := 0; i < size; i++ {
		for len(index) > 0 && num[i] >= num[index[len(index)-1]] {
			index = index[:len(index)-1]
		}
		index = append(index, i)
	}
	for i := size; i < len(num); i++ {
		res = append(res, num[index[0]])
		for len(index) > 0 && num[i] >= num[index[len(index)-1]] {
			index = index[:len(index)-1]
		}
		if len(index) > 0 && index[0] <= (i-size) {
			index = index[1:]
		}
	}
	res = append(res, num[index[0]])
	return res
}
