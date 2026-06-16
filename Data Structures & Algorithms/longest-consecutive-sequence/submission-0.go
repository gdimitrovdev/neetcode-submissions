func longestConsecutive(nums []int) int {
	var seen map[int]bool = make(map[int]bool)
	var seq map[int]int = make(map[int]int)
	var lookingFor map[int]int = make(map[int]int)

	for _, num := range nums {
		_, ok1 := seen[num]
		if ok1 {
			continue
		}

		seen[num] = true

		startOfSeq, ok2 := lookingFor[num]
		if ok2 {
			endSecondSeq, ok3 := seq[num+1]

			if ok3 {
				seq[startOfSeq] = endSecondSeq
				delete(lookingFor, num)
				delete(seq, num+1)
				lookingFor[endSecondSeq + 1] = startOfSeq
			} else {
				seq[startOfSeq] = num
				delete(lookingFor, num)
				lookingFor[num+1] = startOfSeq
			}

			continue
		}

		endSeq, ok4 := seq[num+1]
		if ok4 {
			seq[num] = endSeq
			delete(seq, num+1)
			lookingFor[endSeq + 1] = num

			continue
		}

		seq[num] = num
		lookingFor[num+1] = num
	}

	max := 0
	for k, v := range seq {
		if v - k + 1 > max {
			max = v - k + 1
		}
	}

	return max
}
