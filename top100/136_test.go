package top100

// tag: 异或运算
func singleNumber(nums []int) int {
	var ans int
	for _, s := range nums {
		ans = ans ^ s
	}
	return ans
}

// tag: 哈希表
func singleNumber2(nums []int) int {
	h := make(map[int]struct{})
	for _, s := range nums {
		if _, ok := h[s]; ok {
			delete(h, s)
		} else {
			h[s] = struct{}{}
		}
	}

	for k, _ := range h {
		return k
	}
	return 0
}
