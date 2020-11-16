package simple

/**
	求平方根
 */
func sqrt(x int) int {
	if x == 1 {
		return x
	}
	// 二分法
	start, end := 1, x
	for start <= end {
		mid := start + (end-start)/2
		if mid*mid <= x {
			start = mid
		} else {
			end = mid - 1
		}
	}
	return start
}
