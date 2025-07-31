package Backtracking

var (
	ans    [][]string
	cols   map[int]bool
	diags1 map[int]bool
	diags2 map[int]bool
)

func solveNQueens(n int) [][]string {
	ans = make([][]string, 0)
	cols = make(map[int]bool)
	diags1 = make(map[int]bool)
	diags2 = make(map[int]bool)

	queens := make([]int, n)
	solveNQueensDfs(queens, 0)
	return ans
}

func solveNQueensDfs(queens []int, row int) {
	if row == len(queens) {
		addAns(queens)
		return
	}
	for col := 0; col < len(queens); col++ {
		diag1 := row - col
		diag2 := row + col
		if cols[col] || diags1[diag1] || diags2[diag2] {
			continue
		}
		// 选择这一步
		queens[row] = col
		cols[col] = true
		diags1[diag1] = true
		diags2[diag2] = true
		// 搜索下一步
		solveNQueensDfs(queens, row+1)
		// 恢复现场
		queens[row] = 0
		cols[col] = false
		diags1[diag1] = false
		diags2[diag2] = false
	}
}

func addAns(queens []int) {
	n := len(queens)
	path := make([]string, 0)
	for _, row := range queens {
		tmp := make([]byte, n)
		for col := 0; col < n; col++ {
			if col == row {
				tmp[col] = 'Q'
			} else {
				tmp[col] = '.'
			}
		}
		path = append(path, string(tmp))
	}
	ans = append(ans, path)
}
