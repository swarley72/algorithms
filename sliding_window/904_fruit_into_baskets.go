func totalFruit(fruits []int) int {
	begin := 0
	// тут храним фрукты и их количество
	windowState := make(map[int]int)
	result := 0

	for end := range len(fruits) {
		windowState[fruits[end]]++

		// если фруктов стало больше чем надо
		for len(windowState) > 2 {
			// вычитаем один фрукт (он точно будет т.к. end быстрее begin)
			windowState[fruits[begin]]--
			// если счетчик ноль, удаляем из мапы
			if windowState[fruits[begin]] == 0 {
				delete(windowState, fruits[begin])
			}
			// сдвигаем рамку каждый раз
			begin++
		}
		result = max(result, end-begin+1)
	}
	return result
}