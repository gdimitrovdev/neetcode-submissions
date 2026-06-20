func mergeAlternately(word1 string, word2 string) string {
	ans := ""
	i := 0
	len1 := len(word1)
	len2 := len(word2)

	for i < len1 && i < len2 {
		ans += string([]byte{word1[i]})
		ans += string([]byte{word2[i]})
		i++
	}

	if i < len1 {
		ans += word1[i:]
	}
	if i < len2 {
		ans += word2[i:]
	}

	return ans
}
