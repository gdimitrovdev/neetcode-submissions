func productExceptSelf(nums []int) []int {
	var total int = 1
	var zeroes int = 0
	var totalNoZeroes int = 1

	for _, num := range nums {
		if num == 0 {
			zeroes++
		} else {
			totalNoZeroes *= num
		}
		if zeroes > 1 {
			totalNoZeroes = 0
		}

		total *= num
	}

	var ans []int;
	for _, num := range nums {
		if num == 0 {
			ans = append(ans, totalNoZeroes)
		} else {
			ans = append(ans, total / num)
		}
	}

	return ans
}
