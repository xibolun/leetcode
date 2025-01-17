package top100

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	var source = []struct {
		a      []int
		target int
		rse    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, v := range source {
		r := twoSum(v.a, v.target)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func twoSum(nums []int, target int) []int {
	l, r := 0, 1
	n := len(nums)
	for {
		if l >= n {
			break
		}
		if r == n {
			l++
			r = l + 1
		}
		if nums[l]+nums[r] == target {
			return []int{l, r}
		}
		r++
	}
	return []int{0, 0}
}
