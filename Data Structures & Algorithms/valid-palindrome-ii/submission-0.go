func validPalindrome(s string) bool {
	var l, r int = 0, len(s) - 1
	var allowedChanges int = 1

	return helper(s, l, r, allowedChanges)
}

func helper(s string, l int, r int, allowedChanges int) bool {
	for l < r {
		if s[l] != s[r] {
			if allowedChanges < 1 {
				return false
			}

			if !(helper(s, l+1, r, allowedChanges - 1)) && !(helper(s, l, r-1, allowedChanges - 1)) {
				return false
			}
		}

		l++
		r--
	}

	return true
}
