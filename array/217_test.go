package array

// tag: 数组,哈希
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	for _, s := range nums {
		if _, ok := m[s]; ok {
			return true
		} else {
			m[s] = struct{}{}
		}
	}
	return false
}
