func missingNumber(nums []int) int {
	cnt := len(nums)
    expected := 0
    actual := 0
	for i := 0; i < cnt; i++ {
	    expected += i+1
        actual += nums[i]
	}

	return expected - actual
}
