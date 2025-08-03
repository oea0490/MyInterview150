package Matrix

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])
	x, y := 0, col-1
	for x < row && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else if matrix[x][y] < target {
			x++
		}
	}
	return false
}
