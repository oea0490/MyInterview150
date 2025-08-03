package Matrix

func rotate(matrix [][]int) {
	n := len(matrix)
	row := (n + 1) / 2
	col := n / 2
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = tmp
		}
	}
}
