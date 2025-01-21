package top100

import (
	"reflect"
	"testing"
)

func TestMySqrt(t *testing.T) {
	var source = []struct {
		a   int
		rse int
	}{
		{4, 2},
		{8, 2},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := mySqrt(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func mySqrt(x int) int {
	l, r := 1, x

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}
