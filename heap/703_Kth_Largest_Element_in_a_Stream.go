type KthLargest struct {
	k       int
	minHeap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	minHeap := MinHeap(nums)
	heap.Init(&minHeap)
	for len(minHeap) > k {
		heap.Pop(&minHeap)
	}
	return KthLargest{k: k, minHeap: &minHeap}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	for len(*this.minHeap) > this.k {
		heap.Pop(this.minHeap)
	}
	return (*this.minHeap)[0]
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

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */