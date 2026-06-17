func subarraySum(nums []int, k int) int {
	var sum int = 0
	var prefixSums map[int]int = map[int]int{0: 1}
	var ans int = 0

	for _, num := range nums {
		sum += num
		
		amount, ok := prefixSums[sum - k]
		if ok {
			ans += amount
		}

		prefixSums[sum] += 1
	}

	return ans
}
