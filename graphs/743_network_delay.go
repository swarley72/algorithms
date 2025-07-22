import "container/heap"

func networkDelayTime(times [][]int, n int, k int) int {
	graph := map[int][]GraphNode{}

	for _, net := range times {
		source := net[0]
		target := net[1]
		time := net[2]
		graph[source] = append(graph[source], GraphNode{target, time})
	}

	minHeap := NewMinHeap[GraphNode](func(a, b GraphNode) bool {
		return a.Time < b.Time
	})
	heap.Init(minHeap)
	heap.Push(minHeap, GraphNode{k, 0})

	distances := map[int]int{
		k: 0,
	}

	for len(minHeap.data) > 0 {
		item := heap.Pop(minHeap).(GraphNode)
		node := item.Node
		time := item.Time

		if t, ok := distances[node]; ok && time > t {
			continue
		}

		for _, ngh := range graph[node] {
			nextNode := ngh.Node
			newTime := ngh.Time + time

			if t, ok := distances[nextNode]; !ok || newTime < t {
				distances[nextNode] = newTime
				heap.Push(minHeap, GraphNode{nextNode, newTime})
			}
		}

	}

	if len(distances) != n {
		return -1
	}

	result := 0
	for _, v := range distances {
		result = max(result, v)
	}
	return result
}

type GraphNode struct {
	Node, Time int
}

type MinHeap[T any] struct {
	data []T
	less func(T, T) bool
}

func NewMinHeap[T any](less func(T, T) bool) *MinHeap[T] {
	return &MinHeap[T]{
		data: make([]T, 0),
		less: less,
	}
}

func (h MinHeap[T]) Len() int           { return len(h.data) }
func (h MinHeap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h MinHeap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *MinHeap[T]) Push(x interface{}) {
	h.data = append(h.data, x.(T))
}

func (h *MinHeap[T]) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}
