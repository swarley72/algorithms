type Deque[T any] struct {
	items []T
}

func (d *Deque[T]) Append(item T) {
	d.items = append(d.items, item)
}

func (d *Deque[T]) AppendLeft(item T) {
	d.items = append([]T{item}, d.items...)
}

func (d *Deque[T]) Pop() T {
	pop := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return pop
}

func (d *Deque[T]) PopLeft() T {
	pop := d.items[0]
	d.items = d.items[1:]
	return pop
}