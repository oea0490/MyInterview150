package Substring

func MaxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)
	ans := make([]int, 0)
	for i, num := range nums {
		for len(queue) > 0 && num > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, num)
		if i >= k && nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		if i >= k-1 {
			ans = append(ans, queue[0])
		}
	}
	return ans
}
