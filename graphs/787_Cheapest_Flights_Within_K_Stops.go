
import "container/heap"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	graph := map[int][]Ticket{}

	for _, dest := range flights {
		from := dest[0]
		to := dest[1]
		price := dest[2]
		graph[from] = append(graph[from], Ticket{to, 0, price})
	}

	stops := map[int]int{
		src: 0,
	}

	minHeap := NewHeap[Ticket](func(a, b Ticket) bool {
		return a.Price < b.Price
	})
	heap.Init(minHeap)
	minHeap.PushItem(Ticket{src, 0, 0})

	for len(minHeap.data) > 0 {
		ticket := minHeap.PopItem()

		// потому что если k остановок то мы еще в рамказ дозволенного
		if ticket.Stops > k+1 {
			continue
		}
		if ticket.Dst == dst {
			return ticket.Price
		}
		// обновляем остановки ЗАЧЕМ???
		stops[ticket.Dst] = ticket.Stops

		for _, ngh := range graph[ticket.Dst] {
			newPrice := ngh.Price + ticket.Price
			newStops := ticket.Stops + 1
			// ЗАЧЕМ ПРОВЕРЯТЬ?
			if s, ok := stops[ngh.Dst]; !ok || newStops < s {
				minHeap.PushItem(Ticket{ngh.Dst, newStops, newPrice})
			}
		}
	}

	return -1
}

type Ticket struct {
	Dst   int
	Stops int
	Price int
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
