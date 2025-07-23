// MIN HEAP INT
// ИСПОЛЬЗОВАНИЕ
// minHeap = MinHeap{} || minHeap = MinHeap(slice)
// heap.Init(&minHeap)
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

func (h *MinHeap) PushItem(x int) {
	heap.Push(h, x)
}

func (h *MinHeap) PopItem() int {
	return heap.Pop(h).(int)
}

// MAX HEAP INT
// ИСПОЛЬЗОВАНИЕ
// minHeap = MinHeap{} || minHeap = MinHeap(slice)
// heap.Init(&minHeap)
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
	// Создаем MaxHeap для Person по возрасту
	personHeap := NewHeap[Person](func(a, b Person) bool {
		return a.Age > b.Age
	})
	heap.Init(personHeap)

	// Добавляем элементы
	personHeap.Push(Person{"Alice", 30})

	// Достаем элементы
	personHeap.PopItem()
}

// ===========================================================
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
