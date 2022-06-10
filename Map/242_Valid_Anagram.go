func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := map[byte]int{}
	tMap := map[byte]int{}

	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
		tMap[t[i]]++
	}

	for i := 0; i < len(s); i++ {
		if sMap[s[i]] != tMap[s[i]] {
			return false
		}
	}
	return true
}