func removeElement(nums []int, val int) int {
    var pointer int = 0

    for _, v := range nums {
        if v == val {
            continue
        }

        nums[pointer] = v
        pointer++
    }

    return pointer
}
