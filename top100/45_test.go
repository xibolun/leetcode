package top100

// tag: 贪心

// #55求的是能否跳到终点
// #45求的是跳到终点的最小次数
// 看一下灵神的解答
func jump(nums []int) int {
	curRight := 0
	nextRight := 0 //下一座桥的右端点，即桥的最大长度
	ans := 0
	for i, s := range nums[:len(nums)-1] {
		nextRight = max(nextRight, s+i) // 计算下一座桥最大长度
		if i == curRight {              // 到达当前桥的最右端，就要开始使用下一座桥了
			curRight = nextRight
			ans++
		}
	}
	return ans
}
