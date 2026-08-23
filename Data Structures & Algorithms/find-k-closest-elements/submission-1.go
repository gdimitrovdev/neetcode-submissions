func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func findClosestElements(arr []int, k int, x int) []int {
	var l, r int = 0, k - 1

	for r + 1 < len(arr) && ((Abs(arr[r + 1] - x) < Abs(arr[l] - x)) || (arr[l] == arr[r + 1])) {
		l++
		r++
	}

	return arr[l:r+1]
}
