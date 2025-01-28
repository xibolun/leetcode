package string

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var source = []struct {
		a   string
		rse bool
	}{

		// {"A man, a plan, a canal: Panama", true},
		// {"race a car", false},
		{" ", true},
		{"0P", false},
		{"a", true},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := isPalindrome(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func isPalindrome(s string) bool {
	ms := []byte{}
	for _, r := range s {
		if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') {
			ms = append(ms, byte(r))
		}
		if r >= 'A' && r <= 'Z' {
			ms = append(ms, byte(r+32))
		}
	}

	l, r := 0, len(ms)-1
	for l <= r {
		if ms[l] == ms[r] {
			l++
			r--
			continue
		}
		return false
	}
	return true
}
