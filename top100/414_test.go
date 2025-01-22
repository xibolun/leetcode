package top100

import (
	"math"
	"reflect"
	"testing"
)

func TestThirdMax(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		{[]int{3, 2, 1}, 1},
		{[]int{2, 2, 3, 1}, 1},
		// {[]int{1, 2}, 2},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := thirdMax(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// 描述：给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。要求O(n)
func thirdMax(nums []int) int {
	maxv, secv, thirdv := math.MinInt, math.MinInt, math.MinInt
	// arr := make([]int,0)
	for i := 1; i < len(nums); i++ {
		if nums[i] >= maxv {
			maxv,
				thirdv = secv
			secv = maxv
			maxv = nums[i]
		} else if nums[i] >= secv {
			thirdv = secv
			secv = nums[i]
		} else if nums[i] > thirdv {
			thirdv = nums[i]
		}
	}
	if thirdv == 0 {
		return maxv
	}
	return thirdv

}
