package graphs

import (
	"container/heap"
)

func maxProbability(n int, edges [][]int, succProb []float64, start_node int, end_node int) float64 {

	graph := map[int][]ProbNode{}
	for i := 0; i < len(edges); i++ {
		from := edges[i][0]
		to := edges[i][1]
		prob := succProb[i]

		graph[from] = append(graph[from], ProbNode{to, prob})
		graph[to] = append(graph[to], ProbNode{from, prob})
	}

	probabilities := map[int]float64{
		start_node: 1.0,
	}
	maxHeap := NewHeap[ProbNode](func(a, b ProbNode) bool {
		return a.Prob > b.Prob
	})
	heap.Init(maxHeap)
	maxHeap.PushItem(ProbNode{start_node, 1.0})

	for len(maxHeap.data) > 0 {
		node := maxHeap.PopItem()
		if node.Val == end_node {
			return node.Prob
		}

		for _, ngh := range graph[node.Val] {
			newProb := node.Prob * ngh.Prob
			if p, ok := probabilities[ngh.Val]; !ok || p < newProb {
				probabilities[ngh.Val] = newProb
				maxHeap.PushItem(ProbNode{ngh.Val, newProb})
			}
		}
	}

	return 0.0
}

type ProbNode struct {
	Val  int
	Prob float64
}

type Heap[T any] struct {
	data []T
	less func(T, T) bool
}

func NewHeap[T any](less func(T, T) bool) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0),
		less: less,
	}
}

func (h Heap[T]) Len() int           { return len(h.data) }
func (h Heap[T]) Less(i, j int) bool { return h.less(h.data[i], h.data[j]) }
func (h Heap[T]) Swap(i, j int)      { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *Heap[T]) Push(x interface{}) {
	h.data = append(h.data, x.(T))
}

func (h *Heap[T]) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func (h *Heap[T]) PushItem(x T) {
	heap.Push(h, x)
}

func (h *Heap[T]) PopItem() T {
	return heap.Pop(h).(T)
}

type Set[T comparable] struct {
	elements map[T]struct{}
}

func (s *Set[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *Set[T]) Has(element T) bool {
	_, ok := s.elements[element]
	return ok
}

func (s *Set[T]) Remove(element T) {
	delete(s.elements, element)
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}
