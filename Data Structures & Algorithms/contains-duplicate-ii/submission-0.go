func containsNearbyDuplicate(nums []int, k int) bool {
	counts := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if i > k {
			counts[nums[i - k - 1]]--
		}
		if counts[nums[i]] > 0 {
			return true
		}
		counts[nums[i]]++
	}

	return false
}
