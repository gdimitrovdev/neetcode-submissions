func removeDuplicates(nums []int) int {
	k := 0
	lastNum := -101
	l := len(nums)

	for i := 0; i < l; i++ {
		if nums[i] != lastNum {
			nums[k] = nums[i]
			lastNum = nums[k]
			k++
		}
	}

	return k
}
