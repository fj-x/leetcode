func twoSum(nums []int, target int) []int {
	// create hash map to store processed elements
	indexMap := make(map[int]int)
	for key, val := range nums {
		// subtruct current val from turget
		search := target - val
		// check is subtracted value persist in map
		if val, ok := indexMap[search]; ok {
			return []int{val, key}
		}
		// set val as hash key and key as value
		indexMap[val] = key
	}
	return []int{0, 0}
}
