package backtracking

func permute(nums []int) [][]int {
	result := [][]int{}

	set := NewSet[int]()
	backtrack([]int{}, nums, &result, set)

	return result
}

func backtrack(acc []int, nums []int, res *[][]int, set *Set[int]) {
	if len(acc) == len(nums) {
		tmp := make([]int, len(acc))
		copy(tmp, acc)
		*res = append(*res, tmp)
		return
	}

	for _, n := range nums {
		if set.Has(n) {
			continue
		}
		set.Add(n)
		acc = append(acc, n)
		backtrack(acc, nums, res, set)
		acc = acc[0 : len(acc)-1]
		set.Remove(n)
	}
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
