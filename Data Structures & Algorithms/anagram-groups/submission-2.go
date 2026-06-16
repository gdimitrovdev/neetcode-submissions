func groupAnagrams(strs []string) [][]string {
    hashMap := make(map[string][]string)
    var ans [][]string

    for _, str := range strs {
        runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedStr := string(runes)

        hashMap[sortedStr] = append(hashMap[sortedStr], str)
    }

    for _, v := range hashMap {
        ans = append(ans, v)
    }

    return ans
}
