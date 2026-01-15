package main

import "fmt"

func main() {
	set := NewSet[int]()
	set.Add(1)
	has := set.Has(1)

	fmt.Println(has)
	set.Remove(1)

	fmt.Println(has)

}

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

func (s Set[T]) Has(element T) bool {
	_, ok := s[element]
	return ok
}

func (s Set[T]) Remove(element T) {
	delete(s, element)
}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}
