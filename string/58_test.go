package string

import (
	"reflect"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	var source = []struct {
		a   string
		rse int
	}{
		{"Hello World", 5},
		{"   fly me   to   the moon  ", 4},
		{"luffy is still joyboy", 6},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := lengthOfLastWord(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func lengthOfLastWord(s string) int {
	r := len(s) - 1

	ans := 0
	for r >= 0 {
		if s[r] == ' ' {
			// 已经找到了对应的字符串，现在的位置是在字符串的左边空格
			if ans > 0 {
				break
			}
		} else {
			ans++
		}
		r--
	}
	return ans
}
