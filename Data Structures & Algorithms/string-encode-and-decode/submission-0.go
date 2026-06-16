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
	var currNumStr string = ""
	var ans []string;

	i := 0

	for i < len(encoded) {
		if encoded[i] >= '0' && encoded[i] <= '9' {
			currNumStr += string(encoded[i])
		} else if encoded[i] == '#' {
			currNum, _ := strconv.Atoi(currNumStr)
			currNumStr = ""
			i++
			currStr := ""

			for j := 0; j < currNum; j++ {
				currStr += string(encoded[i])
				i++
			}
			i--

			ans = append(ans, currStr)
		}

		i++
	}

	return ans
}
