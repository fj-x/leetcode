func romanToInt(s string) int {
	var result, previous, current int
	roman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for i := len(s) - 1; i >= 0; i-- {
		current = roman[s[i]]
		if current < previous {
			result -= current
		} else {
			result += current
		}
		previous = current
	}
	return result
}
