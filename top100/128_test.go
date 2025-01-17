package top100

import (
	"reflect"
	"sort"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {

	var source = []struct {
		a   []int
		rse int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0}, 1},
		{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7},
	}

	for _, v := range source {
		r := longestConsecutive(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func longestConsecutive(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	sort.Ints(nums)

	ans := 1
	maxLen := 1

	for i := 1; i < len(nums); i++ {
		step := nums[i] - nums[i-1]
		if step > 1 {
			ans = 1
			continue
		}
		if step == 0 {
			continue
		}
		if step == 1 {
			ans++
		}
		maxLen = max(maxLen, ans)
	}
	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
