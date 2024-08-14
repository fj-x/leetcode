func findLeastNumOfUniqueInts(arr []int, k int) int {
	numbers := make(map[int]int)
	for _, v := range arr {
		numbers[v]++
	}
    if k == 0 {
        return len(numbers)
    } 

	keys := make([]int, 0, len(numbers))
	for key := range numbers {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return numbers[keys[i]] < numbers[keys[j]]
	})

	for _, y := range keys {

		for numbers[y] > 0 {
			numbers[y]--
			k--
			if k == 0 {
				break
			}
		}
		if k == 0 {
			break
		}
	}

	sum := 0
	for _, num := range numbers {
		if num > 0 {
			sum++
		}
	}

	return sum
}
