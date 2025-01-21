package array

import (
	"reflect"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	var source = []struct {
		a      []int
		target int
		rse    int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := searchInsert(v.a, v.target)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 数组,排序

func searchInsert(nums []int, target int) int {
	return lowerBound(nums, target)
}

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
