func isPalindrome(s string) bool {
	var l, r int = 0, len(s) - 1

	for l < r {
		if !(s[l] >= 'a' && s[l] <= 'z') && !(s[l] >= 'A' && s[l] <= 'Z') && !(s[l] >= '0' && s[l] <= '9') {
			l++
			continue
		}

		if !(s[r] >= 'a' && s[r] <= 'z') && !(s[r] >= 'A' && s[r] <= 'Z') && !(s[l] >= '0' && s[l] <= '9') {
			r--
			continue
		}

		if s[l] >= 'a' && s[l] <= 'z' {
			if s[r] != s[l] && s[r] - 'A' + 'a' != s[l] {
				return false
			}
		}

		if s[l] >= 'A' && s[l] <= 'Z' {
			if s[r] != s[l] && s[r] - 'a' + 'A' != s[l] {
				return false
			}
		}

		if s[l] >= '0' && s[l] <= '9' && s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
