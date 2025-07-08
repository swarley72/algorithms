func moveZeroes(nums []int) {
	l := 0
	n := len(nums)
	// до первого нулевого элемента
	for l < n {
		if nums[l] == 0 {
			break
		}
		l++
	}

	// пролжаем с остановки
	for i := l; i < n; i++ {
		if nums[i] != 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
		}
	}
}

func moveZeroes2(nums []int) {
	l := 0

	// тут все в одном цикле
	// если первое число не ноль, то оно поменяется местами само с собой
	// и так будет пока r не встретит ноль
	// когда r первы раз встретит ноль тогда l уже будет стоять на этой точке
	for r := range len(nums) {
		if nums[r] != 0 {
			nums[r], nums[l] = nums[l], nums[r]
			l++
		}
	}
}