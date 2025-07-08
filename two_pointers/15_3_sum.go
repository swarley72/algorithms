import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		// пропускаем дубликаты
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		l := i + 1
		r := len(nums) - 1

		for l < r {
			twoSum := nums[l] + nums[r]

			if twoSum == target {
				res = append(res, []int{nums[l], nums[r], nums[i]})
				l++
				r--

				// пропускаем дубликаты слева
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				// пропускаем дубликаты справа
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if twoSum < target {
				l++
			} else {
				r--
			}
		}
	}

	return res
}
