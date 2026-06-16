func sortArray(nums []int) []int {
	helper(nums, 0, len(nums) - 1)
	return nums
}

func helper(nums []int, start int, end int) {
	if start >= end {
		return
	}

	pivot := nums[end]
	firstPointer := start
	lastPointer := end - 1

	for firstPointer < lastPointer {
		if nums[firstPointer] <= pivot {
			firstPointer++
		} else {
			temp := nums[firstPointer]
			nums[firstPointer] = nums[lastPointer]
			nums[lastPointer] = temp

			lastPointer--
		}
	}

	if nums[lastPointer] > pivot {
		nums[end] = nums[lastPointer]
		nums[lastPointer] = pivot
		helper(nums, start, lastPointer - 1)
		helper(nums, lastPointer + 1, end)
	} else {
		nums[end] = nums[lastPointer + 1]
		nums[lastPointer + 1] = pivot
		helper(nums, start, lastPointer)
		helper(nums, lastPointer + 2, end)
	}
}
