type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var ans string = ""

	for _, str := range strs {
		ans += strconv.Itoa(len(str))
		ans += "#"
		ans += str
	} 
	return ans
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		i = j + 1
		res = append(res, encoded[i:i+length])
		i += length
	}
	return res
}
