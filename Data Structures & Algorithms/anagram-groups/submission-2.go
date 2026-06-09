func sorts(s string) string {
    r := []rune(s)
    // В Go 1.21+ лучше использовать slices.Sort(r)
    sort.Slice(r, func(i, j int) bool {
        return r[i] < r[j]
    })
    return string(r) // Прямое приведение эффективнее Builder здесь
}

func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
        return [][]string{}
    }

    dict := make(map[string][]string)
    for _, s := range strs {
        key := sorts(s)
        dict[key] = append(dict[key], s)
    }

    res := make([][]string, 0, len(dict))
    for _, group := range dict {
        // Просто добавляем готовую группу, не нужно перебирать её циклом
        res = append(res, group)
    }

    return res
}