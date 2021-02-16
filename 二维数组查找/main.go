package main

func Find(target int, array [][]int) bool {
	var colmuns int = len(array[0])
	var rows int = len(array)
	if array != nil && rows > 0 && colmuns > 0 {
		row := 0
		colmun := colmuns - 1
		for row < rows && colmun > 0 {
			if array[row][colmun] == target {
				return true
			} else if array[row][colmun] < target {
				row++
			} else {
				colmun--
			}
		}
	}
	return false
}

func main() {
	Find(1, [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}})
}
