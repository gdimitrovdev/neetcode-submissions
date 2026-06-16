func groupAnagrams(strs []string) [][]string {
    hashMap := make(map[string][]string)
    var ans [][]string

    for _, str := range strs {
        split := strings.Split(str, "")
        sort.Strings(split)
        sortedStr := strings.Join(split, "")

        hashMap[sortedStr] = append(hashMap[sortedStr], str)
    }

    for _, v := range hashMap {
        ans = append(ans, v)
    }

    return ans
}
