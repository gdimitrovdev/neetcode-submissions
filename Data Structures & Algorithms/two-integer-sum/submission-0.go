func twoSum(nums []int, target int) []int {
    hashMap := make(map[int]int)

	for j, v := range nums {
		i, ok := hashMap[target - v]
		if ok {
			return []int{i, j}
		} else {
			hashMap[v] = j
		}
	}

	return []int{0, 0}
}
