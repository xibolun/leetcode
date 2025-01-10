package bit

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/counting-bits/description/

func TestCountBits(t *testing.T) {
	var source = []struct {
		a   int
		rse []int
	}{
		{5, []int{0, 1, 1, 2, 1, 2}},
	}

	for _, v := range source {
		// r := countBits(v.a)
		r := countBits2(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 当 i 的最低位是0，则 i 中1的个数和i >> 1中1的个数相同；当i的最低位是1，i 中1的个数是 i >> 1中1的个数再加1
		res[i] = res[i>>1] + (i & 1)
	}
	return res
}

/*
*
  - 位运算
  - 偶数的二进制1个数超级简单，因为偶数是相当于被某个更小的数乘2，乘2怎么来的？在二进制运算中，就是左移一位，也就是在低位多加1个0，那样就说明dp[i] = dp[i / 2]
  - 奇数稍微难想到一点，奇数由不大于该数的偶数+1得到，偶数+1在二进制位上会发生什么？会在低位多加1个1，那样就说明dp[i] = dp[i-1] + 1，当然也可以写成dp[i] = dp[i / 2] + 1
*/
func countBits2(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1
		}
	}

	return res
}
