package top100

// tag: 贪心

// #55求的是能否跳到终点
// #45求的是跳到终点的最小次数
// 看一下灵神的解答
func jump(nums []int) int {
	curRight := 0
	nextRight := 0
	ans := 0
	for i, s := range nums[:len(nums)-1] {

		nextRight = max(nextRight, s+i)
		if i == curRight {
			curRight = nextRight
			ans++
		}
	}
	return ans
}
