func rotate(nums []int, k int) {
    length := len(nums)
    moved := map[int]bool{}

    for i := 0; i < length; i++ {
        _, ok := moved[i]
        index := i
        last_num := nums[i]

        for !ok {
            temp := nums[(index + k) % length]
            nums[(index + k) % length] = last_num
            last_num = temp
            moved[index] = true
            _, ok = moved[(index + k) % length]
            index = (index + k) % length
        }
    }
}
