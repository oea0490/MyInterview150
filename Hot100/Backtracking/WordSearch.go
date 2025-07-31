package Backtracking

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && existDfs(i, j, board, word, visited, 0) {
				return true
			}
		}
	}
	return false
}

func existDfs(i, j int, board [][]byte, word string, visited [][]bool, index int) bool {
	if visited[i][j] || board[i][j] != word[index] {
		return false
	}
	if index == len(word)-1 {
		return true
	}

	visited[i][j] = true
	for _, dir := range dirs {
		newI, newJ := i+dir[0], j+dir[1]
		if newI >= 0 && newI < len(board) && newJ >= 0 && newJ < len(board[0]) {
			if existDfs(newI, newJ, board, word, visited, index+1) {
				return true
			}
		}
	}
	visited[i][j] = false
	return false
}
