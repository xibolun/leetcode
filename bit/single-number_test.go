package bit

import "testing"

// https://leetcode.cn/problems/single-number/description/

func TestSingleNumber(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{2, 3, 4, 5, 5, 4, 3, 2, 1}, 1},
	}

	for _, v := range source {
		r := singleNumber(v.a)
		if r != v.rse {
			t.Errorf("test fail, result: %d, should be :%d", r, v.rse)
		}
	}
}

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		//  2进制，相同为0，不同为1
		res ^= v
	}
	return res
}
