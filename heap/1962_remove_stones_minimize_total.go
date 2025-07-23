import (
	"container/heap"
	"math"
)

func minStoneSum(piles []int, k int) int {
	maxHeap := MaxHeap(piles)
	heap.Init(&maxHeap)

	for i := 0; i < k; i++ {
		n := maxHeap.PopItem()
		n = int(math.Ceil(float64(n) / 2.0))
		if n != 0 {
			maxHeap.PushItem(n)
		}
	}
	return sum(maxHeap)
}

func sum(nums []int) int {
	count := 0
	for _, n := range nums {
		count += n
	}
	return count
}

// MAX HEAP INT
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
