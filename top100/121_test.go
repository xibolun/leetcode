package top100

import (
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		// {[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		// {[]int{-1, 2}, 2},
	}

	for _, v := range source {
		r := maxProfit(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 前缀和
func maxProfit(prices []int) int {
	maxPreSum := prices[0]
	minPreSum := prices[0]
	ans := 0
	for _, s := range prices {
		// 注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
		if maxPreSum > s {
			maxPreSum = s
		}
		maxPreSum = max(maxPreSum, s)
		minPreSum = min(minPreSum, s)
		ans = max(ans, maxPreSum-minPreSum)
	}
	return ans
}
