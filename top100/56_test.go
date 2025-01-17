package top100

func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	for i, s := range intervals {
		smin := s[0]
		smax := s[1]
		for j := i + 1; j < len(intervals); j++ {
			ssmin := intervals[j][0]
			ssmax := intervals[j][1]

			if smax > ssmin {
				ans = append(ans, []int{smin, ssmax})
			}
			i++
			break
		}
	}
	return ans
}
