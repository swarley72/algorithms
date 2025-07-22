// MIN HEAP INT
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

// MIN HEAP STRUCT

// ПРИМЕР ================================================
type Person struct {
	Name string
	Age  int
}

func main() {
	// Создаем MinHeap для Person по возрасту
	personHeap := NewMinHeap[Person](func(a, b Person) bool {
		return a.Age < b.Age
	})
	heap.Init(personHeap)

	// Добавляем элементы
	heap.Push(personHeap, Person{"Alice", 30})
}

// ===========================================================
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