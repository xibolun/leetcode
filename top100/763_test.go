package top100

// tag: 贪心,hash表
// 如果一个区间要包括a，那么分割的区间里面要包括所有的a
func partitionLabels(s string) []int {
	maxMap := make(map[int32]int)
	ans := make([]int, 0)
	for i, item := range s {
		maxMap[item] = i
	}

	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		r = max(maxMap[int32(s[i])], r)
		if i == r { //走到了边界点
			ans = append(ans, r-l+1)
			l = i + 1
		}
	}
	return ans
}
