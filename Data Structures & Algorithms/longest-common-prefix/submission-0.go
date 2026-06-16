func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

    var shortest string = strs[0];
	var length int = len(shortest);

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < length {
			length = len(strs[i])
			shortest = strs[i]
		}
	}

	for _, str := range strs {
		for i := 0; i < length; i++ {
			if str[i] != shortest[i] {
				length = i
			}
		}
	}

	return shortest[:length]
}
