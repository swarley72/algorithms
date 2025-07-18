import "sort"

func meetingRooms2(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	minHeap := MinHeap{}
	heap.Push(&minHeap, intervals[0][1])

	for i := 1; i < len(intervals); i++ {
		meeting := intervals[i]
		if minHeap[0] <= meeting[0] {
			heap.Pop(&minHeap)
		}
		heap.Push(&minHeap, meeting[1])
	}
	return len(minHeap)
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