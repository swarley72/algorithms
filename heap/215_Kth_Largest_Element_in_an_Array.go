
// MAX HEAP
func findKthLargest(nums []int, k int) int {
	maxHeap := MaxHeap(nums)

	heap.Init(&maxHeap)

	for i := 0; i < k-1; i++ {
		heap.Pop(&maxHeap)
	}

	return maxHeap[0]
}

// MIN HEAP
func findKthLargest(nums []int, k int) int {
	minHeap := MinHeap{}

	for i := 0; i < len(nums); i++ {
		heap.Push(&minHeap, nums[i])
		if len(minHeap) > k {
			heap.Pop(&minHeap)
		}
	}

	return minHeap[0]
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}