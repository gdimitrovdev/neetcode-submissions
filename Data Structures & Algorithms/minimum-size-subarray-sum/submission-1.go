func minSubArrayLen(target int, nums []int) int {
    l, r, shortest, currSum := 0, 0, 0, 0

    for r < len(nums) {
        if currSum < target {
            currSum += nums[r]
            r++
        } else {
            if r - l < shortest || shortest == 0 {
                shortest = r - l
            }
            currSum -= nums[l]
            l++
        }
    }

    for l < len(nums) && currSum >= target {
        if r - l < shortest || shortest == 0 {
            shortest = r - l
        }
        currSum -= nums[l]
        l++
    }

    return shortest
}
