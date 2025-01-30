package array

// tag: 迭代
// similar: 119

// [1]
// [1,1]
// [1,2,1]
// [1,3,3,1]
// [1,4,6,4,1]
func generate(numRows int) [][]int {
	ans := make([][]int, numRows)

	for i := range ans {
		ans[i] = make([]int, i+1)
		// 每一行的第一个数字和最后一个数字都是0
		ans[i][0], ans[i][i] = 1, 1
		// 其他的都满足c[i][j]=c[i−1][j−1]+c[i−1][j]
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}

	return ans
}
