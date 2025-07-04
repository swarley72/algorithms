func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for idx, n := range nums {
		if prevIdx, ok := hashMap[target-n]; ok {
			return []int{prevIdx, idx}
		}
		hashMap[n] = idx
	}

	return nil
}
