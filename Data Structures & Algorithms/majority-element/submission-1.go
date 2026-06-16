func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

    var currentElement int = nums[0];
	var currentCount int = 1;

	for _, v := range nums[1:] {
		if v == currentElement {
			currentCount++
		} else {
			currentCount--
		}

		if currentCount <= 0 {
			currentElement = v
		}
	}

	return currentElement
}
