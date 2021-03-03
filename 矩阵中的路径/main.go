package main

func hasPath(matrix string, rows int, cols int, str string) bool {
	if matrix == "" || rows < 1 || cols < 1 || str == "" {
		return false
	}
	visited := make([]bool, rows*cols)
	pathLen := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if hasPathHelper(matrix, rows, cols, row, col, str, &pathLen, visited) {
				return true
			}
		}
	}
	return false
}

func hasPathHelper(matrix string, rows, cols, row, col int, str string, pathLen *int, visited []bool) bool {
	if *pathLen >= len(str) {
		return true
	}
	hasPath := true
	if row >= 0 && row < rows && col >= 0 && col < cols && matrix[row*cols+col] == str[*pathLen] && !visited[row*cols+col] {
		*pathLen++
		visited[row*cols+col] = true
		hasPath = hasPathHelper(matrix, rows, cols, row, col-1, str, pathLen, visited) || hasPathHelper(matrix, rows, cols, row-1, col, str, pathLen, visited) ||
			hasPathHelper(matrix, rows, cols, row, col+1, str, pathLen, visited) || hasPathHelper(matrix, rows, cols, row+1, col, str, pathLen, visited)
		if !hasPath {
			*pathLen--
			visited[row*cols+col] = false
		}
	}
	return hasPath
}
