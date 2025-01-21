package array

import (
	"reflect"
	"testing"
)

func TestFindClosestNumber(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		{[]int{-4, -2, 1, 4, 8}, 1},
		{[]int{2, -1, 1}, 1},
		{[]int{2, 1, 1, -1, 100000}, 1},
		{[]int{-1, 1, 0, 0}, 0},
	}

	for _, v := range source {
		r := findClosestNumber(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func findClosestNumber(nums []int) int {
	ans := nums[0]
	for _, s := range nums {
		if abs(s) < abs(ans) || abs(s) == abs(ans) && s > 0 {
			ans = s
		}
	}
	return ans
}

func abs(s int) int {
	if s > 0 {
		return s
	}
	return -s
}
