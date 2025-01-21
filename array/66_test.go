package array

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	var source = []struct {
		a   []int
		rse []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := plusOne(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 数组,求和进位
func plusOne(digits []int) []int {
	carry := 1
	ans := make([]int, 0)
	for i := len(digits) - 1; i >= 0; i-- {
		if carry == 0 {
			continue
		}
		v := digits[i] + carry
		digits[i] = v % 10
		carry = v / 10
	}
	// 若最后还需要再+1
	if carry == 1 {
		ans = append([]int{1}, digits...)
		return ans
	}
	return digits
}
