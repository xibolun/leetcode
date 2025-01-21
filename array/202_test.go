package array

// 最终会得到 1。
// 最终会进入循环。
// 值会越来越大，最后接近无穷大。
// tag: 快乐数

// FIXME
func isHappy(n int) bool {
	arr := make([]int, 0)
	buildArray(n, arr)

	if v := sqrt(arr); v == 1 {
		return true
	} else {
		return isHappy(v)
	}
}

func buildArray(n int, arr []int) {
	arr = append(arr, n%10)
	if n/10 > 0 {
		buildArray(n, arr)
	}
}

func sqrt(arr []int) int {
	ans := 0
	for _, item := range arr {
		ans += item * item
	}
	return ans
}
