func sortColors(nums []int) {
    counter0, counter1, counter2 := 0, 0, 0

	for _, v := range nums {
		if v == 0 {
			counter0++
		} else if v == 1 {
			counter1++
		} else {
			counter2++
		}
	}

	i := 0

	for i < counter0 {
		nums[i] = 0
		i++
	}
	for i < counter0 + counter1 {
		nums[i] = 1
		i++
	}
	for i < len(nums) {
		nums[i] = 2
		i++
	}
}
