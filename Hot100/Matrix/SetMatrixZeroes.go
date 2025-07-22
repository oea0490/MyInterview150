package Matrix

func setZeroes(matrix [][]int) {
	row, col := len(matrix), len(matrix[0])
	row0flag, col0flag := false, false
	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			col0flag = true
			break
		}
	}
	for j := 0; j < col; j++ {
		if matrix[0][j] == 0 {
			row0flag = true
			break
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if col0flag {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
	if row0flag {
		for j := 0; j < col; j++ {
			matrix[0][j] = 0
		}
	}
}
