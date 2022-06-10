func canPermutePalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	saved := map[rune]bool{}

	for _, r := range s {
		if _, ok := saved[r]; ok {
			delete(saved, r)
		} else {
			saved[r] = true
		}
	}

	return len(saved) <= 1
}