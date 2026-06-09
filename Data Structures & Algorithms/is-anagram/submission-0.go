func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dict := make(map[int32]int)
	for _, i := range s {
		dict[i]++
	}
	for _, i := range t {
		dict[i]--
	}

	for _, i := range dict {
		if i != 0 {
			return false
		}
	}
	return true
}