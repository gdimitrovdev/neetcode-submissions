func firstMissingPositive(nums []int) int {
    var n int = len(nums)

    for i := 0; i < n; i++ {
        if nums[i] < 0 {
            nums[i] = 0
        }
    }

    for i := 0; i < n; i++ {
        num := nums[i]
        if num < 0 {
            num = -num
        }

        if num >= 1 && num < n + 1 {
            if nums[num-1] > 0 {
                nums[num-1] = -nums[num-1]
            } else if nums[num-1] == 0 {
                nums[num-1] = -(n + 1)
            }
        }
    }

    for i := 0; i < n; i++ {
        if nums[i] >= 0 {
            return i + 1
        }
    }

    return n + 1
}
