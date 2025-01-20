package doublepoint

import (
	"reflect"
	"testing"
)

//  287. 寻找重复数
// tag: 数组,哈希表,快慢指针
// https://leetcode.cn/problems/find-the-duplicate-number/description/

func TestFindDuplicate(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{3, 3, 3, 3, 3}, 3},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := findDuplicate(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func findDuplicate(nums []int) int {
	windows := make(map[int]int)
	for _, c := range nums {
		if v := windows[c]; v > 0 {
			return c
		}
		windows[c]++
	}
	return 0
}
