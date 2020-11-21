package dp

import "testing"

func TestClimbStairs(t *testing.T) {
	var source = []struct {
		numbers int
		count   int
	}{
		{1, 1}, {2, 2}, {3, 3},
	}

	for _, v := range source {
		res := climbStairs(v.numbers)
		if res != v.count {
			t.Errorf("test fail, input: %d, result: %d, should be :%d", v.numbers, res, v.count)
		}
	}
}

/**
你在爬楼梯，需要n步才能爬到楼梯顶部
每次你只能向上爬1步或者2步。有多少种方法可以爬到楼梯顶部？
e.g.
	1 --> 1
	3 --> 3
此题类似一个斐波那契数列
*/
func climbStairs(n int) int {
	dp := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
