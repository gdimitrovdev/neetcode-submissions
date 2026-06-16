func groupAnagrams(strs []string) [][]string {
    hashMap := make(map[string][]string)
    var ans [][]string

    for _, str := range strs {
        split := strings.Split(str, "")
        sort.Strings(split)
        sortedStr := strings.Join(split, "")

        _, ok := hashMap[sortedStr]
        if ok {
            l := hashMap[sortedStr]
            l = append(l, str)
            hashMap[sortedStr] = l
        } else {
            hashMap[sortedStr] = []string{str}
        }
    }

    for _, v := range hashMap {
        ans = append(ans, v)
    }

    return ans
}
