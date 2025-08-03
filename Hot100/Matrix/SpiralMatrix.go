package Matrix

var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	row := len(matrix)
	col := len(matrix[0])
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	x, y, curDir := 0, 0, 0
	for !visited[x][y] {
		visited[x][y] = true
		ans = append(ans, matrix[x][y])
		// 更新下一步位置
		nextX, nextY := x+dir[curDir][0], y+dir[curDir][1]
		// 判断下一步位置是否越界或者位置是否被访问过
		if nextX < 0 || nextX >= row || nextY < 0 || nextY >= col || visited[nextX][nextY] {
			curDir = (curDir + 1) % 4
			nextX, nextY = x+dir[curDir][0], y+dir[curDir][1]
		}
		// 如果更新方向后的位置越界或者位置被访问过，则停止
		if nextX < 0 || nextX >= row || nextY < 0 || nextY >= col || visited[nextX][nextY] {
			break
		}
		x, y = nextX, nextY
	}
	return ans
}
