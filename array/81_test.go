package array

import (
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	var source = []struct {
		a   []int
		k   int
		rse bool
	}{
		{[]int{2, 5, 6, 0, 0, 1, 2}, 0, true},
		{[]int{2, 5, 6, 0, 0, 1, 2}, 3, false},
		{[]int{1}, 1, true},
		{[]int{5, 1, 3}, 1, true},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := search(v.a, v.k)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 双指针,数组旋转
func search(nums []int, target int) bool {
	// l为旋转点
	l, r := 0, len(nums)-1
	// 前值
	pre := 0

	if nums[l] > target && nums[r] < target {
		return false
	}
	if nums[l] == target || nums[r] == target {
		return true
	}

	// target位于 [k:]
	if nums[l] < target {
		pre = nums[l]
		for l < len(nums)-1 && pre <= target {
			if nums[l] == target {
				return true
			}
			pre = nums[l]
			l++

		}
	}

	// target位于 [0,k]，需要倒序
	if nums[r] > target {
		pre = target
		for r >= 0 && pre >= target {
			if nums[r] == target {
				return true
			}
			pre = nums[r]
			r--
		}

	}
	return false
}
