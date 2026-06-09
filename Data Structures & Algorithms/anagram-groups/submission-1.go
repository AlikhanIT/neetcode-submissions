func sorts(strs []rune) string {
	var sb strings.Builder
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	for _, el := range strs {
		sb.WriteRune(el)
	}
	return sb.String()
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) <= 1 || len(strs) >= 1000 {
		return [][]string{strs}
	}

	dict := make(map[string][]string, len(strs))
	for i := 0; i < len(strs); i++ {
		key := sorts([]rune(strs[i]))
		dict[key] = append(dict[key], strs[i])
	}

	res := make([][]string, len(dict))
	iter := 0
	for _, els := range dict {
		res[iter] = make([]string, 0, len(els))
		for _, el := range els {
			res[iter] = append(res[iter], el)
		}
		iter++
	}

	return res
}
