package string

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	var source = []struct {
		a   string
		rse int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := romanToInt(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

var mapper = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}

func romanToInt(s string) int {
	ans := 0

	// 为何从1开始？因为需要计算s[i-1]
	// 为何是len(s)，因为最后一位肯定是不会出现特殊情况的，累加即可
	for i := 1; i < len(s); i++ {
		l, r := mapper[s[i-1]], mapper[s[i]]
		// 说明为特殊的情况
		if l < r {
			// 相当于ans = ans + (r-l) 由于ans+r是i+1的操作，所以直接-l即可
			ans = ans - l
		} else {
			ans = ans + l
		}
	}
	return ans + mapper[s[len(s)-1]]
}
