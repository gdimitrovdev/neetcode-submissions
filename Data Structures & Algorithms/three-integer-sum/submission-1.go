func threeSum(nums []int) [][]int {
    ans := [][]int{}
    sort.Slice(nums, func(i int, j int) bool {
        return nums[i] <= nums[j]
    })
    length := len(nums)
    
    for i := 0; i < length; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        l, r := i + 1, length - 1
        for l < r {
            if nums[l] + nums[r] < -nums[i] {
                l++
            } else if nums[l] + nums[r] > -nums[i] {
                r--
            } else {
                currAns := []int{nums[l], nums[r], nums[i]}
                l++
                r--
                ans = append(ans, currAns)
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
            }
        }
    }

    return ans
}
