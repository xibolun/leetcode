package dichotomy

import (
	"reflect"
	"testing"
)

func TestFindMagicIndex(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		{[]int{0, 2, 3, 4, 5}, 0},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 1},
		{[]int{3, 4, 5, 5, 5, 5}, 5},
	}

	for _, v := range source {
		r := findMagicIndex2(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// 暴力法，复杂度O(n)
func findMagicIndex2(nums []int) int {
	for i, c := range nums {
		if c == i {
			return c
		}
	}
	return -1
}

func findMagicIndex(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}
	l := 0
	r := len(nums) - 1
	mid := (l + r) / 2
	if nums[mid] == mid {
		return mid
	}
	//FIXME 因为数组是会有重复的，所以不能取左或取右，只能剪枝，先左，后右
	if nums[mid] < mid {
		return findMagicIndex(nums[:mid])
	}
	if nums[mid] > mid {
		return findMagicIndex(nums[mid:])
	}
	return -1

}
