package heapasdf

import (
	"container/heap"
	"math"
)

func lastStoneWeight(stones []int) int {

	maxHeap := MaxHeap(stones)
	heap.Init(&maxHeap)

	for len(maxHeap) > 1 {
		first := maxHeap.PopItem()
		second := maxHeap.PopItem()

		result := int(math.Abs(float64(first - second)))
		if result != 0 {
			maxHeap.PushItem(result)
		}

	}

	if len(maxHeap) == 0 {
		return 0
	}
	return maxHeap[0]
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

func (h *MaxHeap) PushItem(x int) {
	heap.Push(h, x)
}

func (h *MaxHeap) PopItem() int {
	return heap.Pop(h).(int)
}
