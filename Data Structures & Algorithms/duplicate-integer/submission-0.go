
func hasDuplicate(nums []int) bool {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dict[nums[i]]++
		if dict[nums[i]] > 1 {
			return true
		}
	}
	return false
}