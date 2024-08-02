func isPalindrome(x int) bool {
	// less than  0 cant be a palindrom
	if x < 0 {
		return false
	}
	// single digit is a palindrom
	if x < 10 {
		return true
	}
	// if number end to 0 it is not a palindrom
	if x%10 == 0 {
		return false
	}

	reverted := 0
	for x > 0 {
        // take the last namber from x and add to reverted
		reverted = reverted*10 + x%10
        // remove last number from x
		x = x / 10
        // if x is palindrom we will iterate only half of digits
		if reverted == x || reverted == x/10 {
			return true
		}
	}
    return false
}
