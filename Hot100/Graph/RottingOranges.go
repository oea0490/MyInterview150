package Graph

func orangesRotting(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	time := -1
	fresh := 0
	queue := make([][]int, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	if len(queue) == 0 {
		if fresh == 0 {
			return 0
		}
		return -1
	}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			for _, d := range direction {
				x := cur[0] + d[0]
				y := cur[1] + d[1]
				if x >= 0 && x < row && y >= 0 && y < col && grid[x][y] == 1 {
					grid[x][y] = 2
					fresh--
					queue = append(queue, []int{x, y})
				}
			}
		}
		time++
	}
	if fresh != 0 {
		return -1
	}
	return time
}
