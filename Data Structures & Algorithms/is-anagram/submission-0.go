func isAnagram(s string, t string) bool {
	ls, lt := len(s), len(t)
	if ls != lt {
		return false
	}

	hashMap := make(map[rune]int)

	for _, v := range s {
		_, ok := hashMap[v]
		if !ok {
			hashMap[v] = 1
		} else {
			hashMap[v] += 1
		}
	}

	for _, v := range t {
		_, ok := hashMap[v]
		if !ok {
			return false
		} else {
			hashMap[v] -= 1
			if hashMap[v] < 0 {
				return false
			}
		}
	}

	return true
}
