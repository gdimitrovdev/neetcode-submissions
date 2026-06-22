func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	ans := [][]int{}
	if length < 4 {
		return ans
	}

	sort.Ints(nums)

	for p1 := 0; p1 < length - 3; p1++ {
		if p1 > 0 && nums[p1] == nums[p1 - 1] {
			continue
		}

		for p2 := p1 + 1; p2 < length - 2; p2++ {
			if p2 > p1 + 1 && nums[p2] == nums[p2 - 1] {
				continue
			}

			l, r := p2 + 1, length - 1
			for l < r {
				if nums[p1] + nums[p2] + nums[l] + nums[r] < target {
					l++
					continue
				} else if nums[p1] + nums[p2] + nums[l] + nums[r] > target {
					r--
					continue
				} else {
					ans = append(ans, []int{nums[p1], nums[p2], nums[l], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l - 1] {
						l++
					}
				}
			}
		}
	}

	return ans
}
