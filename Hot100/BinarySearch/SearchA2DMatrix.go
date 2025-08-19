package BinarySearch

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	// 第一次二分查找,找到target所在的行
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] > target {
			right = mid - 1
		} else if matrix[mid][0] < target {
			left = mid + 1
		}
	}
	row := right
	if row < 0 || row >= n {
		return false
	}
	// 第二次二分查找,找到target所在的列
	left, right = 0, m-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] > target {
			right = mid - 1
		} else if matrix[row][mid] < target {
			left = mid + 1
		}
	}
	return false
}
