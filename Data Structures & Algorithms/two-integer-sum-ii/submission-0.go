func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers) - 1
	ans := []int{}

	for l < r {
		if numbers[l] + numbers[r] > target {
			r--
			continue
		} else if numbers[l] + numbers[r] < target {
			l++
			continue
		} else {
			ans = append(ans, l + 1)
			ans = append(ans, r + 1)
			return ans
		}
	}

	return ans
}
