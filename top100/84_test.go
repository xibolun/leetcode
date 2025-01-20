package top100

import (
	"reflect"
	"testing"
)

func TestLargestRectangleArea(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
	}

	for _, v := range source {
		r := largestRectangleArea(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 数组,双指针
// 暴力解法
func largestRectangleArea(heights []int) int {
	// 核心思想
	// max (r-l+1)*heights[i]
	n := len(heights)
	ans := 0

	for i := 0; i < len(heights); i++ {
		// 一个向左，一个向右
		l, r := i, i
		iv := heights[i]

		// 找出左侧最高点
		for l-1 >= 0 && heights[l-1] >= iv {
			l--
		}
		// 找右侧最高的柱子
		for r+1 < n && heights[r+1] >= iv {
			r++
		}

		ans = max(ans, (r-l+1)*iv)
	}
	return ans
}
