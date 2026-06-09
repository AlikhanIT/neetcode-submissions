func topKFrequent(nums []int, k int) []int {
	dict := make(map[int][]int, k)

	for _, i := range nums {
		if dict[i] == nil {
			dict[i] = []int{i, 0}
		}
		iter := dict[i][1]
		dict[i] = []int{i, iter + 1}
	}

	sort.Slice(nums, func(i, j int) bool {
		return dict[nums[i]][1] > dict[nums[j]][1]
	})

	res := make([]int, k)
	dict2 := make(map[int]bool)
	iter := 0
	for _, i := range nums {
		if iter >= k {
			break
		}
		if dict2[i] {
			continue
		}
		dict2[i] = true
		res[iter] = i
		iter++
	}

	return res
}