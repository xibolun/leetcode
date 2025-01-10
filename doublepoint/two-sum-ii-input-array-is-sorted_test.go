package doublepoint

import (
	"reflect"
	"testing"
)

// 167. 两数之和 II - 输入有序数组
//  相向双指针
//  https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/

func TestTwoSum(t *testing.T) {
	var source = []struct {
		a      []int
		target int
		rse    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := twoSum(v.a, v.target)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for {
		s := numbers[l] + numbers[r]
		if s == target {
			// return []int{l, r} // 题目要求返回下标+1
			return []int{l + 1, r + 1}
		}
		//  若是大于，说明右边的数字比较大，需要左移
		if s > target {
			r--
		} else {
			l++
		}
	}
}
