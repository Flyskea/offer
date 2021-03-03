package main

import "fmt"

func movingCount(threshold int, rows int, cols int) int {
	visited := make([]bool, rows*cols)
	count := movingCountHelper(threshold, rows, cols, 0, 0, visited)
	return count
}

func movingCountHelper(threshold, rows, cols, row, col int, visited []bool) int {
	count := 0
	if row >= 0 && row < rows && col >= 0 && col < cols && getDigitSum(row)+getDigitSum(col) <= threshold && !visited[row*cols+col] {
		visited[row*cols+col] = true
		count = 1 + movingCountHelper(threshold, rows, cols, row+1, col, visited) + movingCountHelper(threshold, rows, cols, row-1, col, visited) +
			movingCountHelper(threshold, rows, cols, row, col+1, visited) + movingCountHelper(threshold, rows, cols, row, col-1, visited)
	}
	return count
}

func getDigitSum(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func main() {
	fmt.Println(movingCount(5, 10, 10))
}
