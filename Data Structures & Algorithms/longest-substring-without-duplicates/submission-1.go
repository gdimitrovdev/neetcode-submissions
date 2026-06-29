func lengthOfLongestSubstring(s string) int {
    counts := make(map[byte]int)
    maxLen := 0
    startPointer := 0

    for i := 0; i < len(s); i++ {
        if counts[s[i]] == 0 {
            counts[s[i]]++
        } else {
            maxLen = max(maxLen, i - startPointer)
            for counts[s[i]] > 0 {
                counts[s[startPointer]]--
                startPointer++
            }
            counts[s[i]]++
        }
    }

    maxLen = max(maxLen, len(s) - startPointer)

    return maxLen
}
