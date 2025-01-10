package dichotomy

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	var source = []struct {
		nums   []int
		target int
		rse    []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
		{[]int{2, 2}, 2, []int{0, 1}},
	}
	for _, v := range source {
		r := searchRange(v.nums, v.target)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}
func searchRange(nums []int, target int) []int {
	start := findMid(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}

	end := findMid(nums, target+1) - 1
	return []int{start, end}

}

func findMid(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	// 循环结束后 left = right+1
	// 此时 nums[left-1] < target 而 nums[left] = nums[right+1] >= target
	// 所以 left 就是第一个 >= target 的元素下标
	return l
}
