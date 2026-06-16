func hasDuplicate(nums []int) bool {
    var hashMap map[int]bool = make(map[int]bool)

    for _, v := range nums {
        _, ok := hashMap[v]
        if ok {
            return true
        }
        hashMap[v] = true
    }

    return false
}
