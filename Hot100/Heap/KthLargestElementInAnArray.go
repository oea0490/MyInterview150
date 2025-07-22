package Heap

var (
	size int
	heap []int
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	size = n
	heap = make([]int, n+1)

	for i := 0; i < n; i++ {
		heap[i+1] = nums[i]
	}

	for i := n / 2; i >= 1; i-- {
		maxHeapify(i)
	}

	for i := 1; i < k; i++ {
		heap[1] = heap[size]
		size--
		maxHeapify(1)
	}

	return heap[1]
}

func maxHeapify(u int) {
	t := u
	if u*2 <= size && heap[u*2] > heap[t] {
		t = u * 2
	}
	if u*2+1 <= size && heap[u*2+1] > heap[t] {
		t = u*2 + 1
	}
	if t != u {
		swap(t, u)
		maxHeapify(t)
	}
}

func swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}
