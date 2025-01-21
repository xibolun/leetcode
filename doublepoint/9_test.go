package doublepoint

import (
	"reflect"
	"strconv"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var source = []struct {
		a   int
		rse bool
	}{
		{121, true},
	}

	for _, v := range source {
		r := isPalindrome(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 字符串,双指针,回文字符串
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	l, r := 0, len(s)-1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
