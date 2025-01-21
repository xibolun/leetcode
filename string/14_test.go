package string

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var source = []struct {
		a   []string
		rse string
	}{
		// {[]string{"flower", "flow", "flight"}, "fl"},
		// {[]string{"dog", "racecar", "car"}, ""},
		{[]string{"c", "acc", "ccc"}, ""},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := longestCommonPrefix(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	minl := len(strs[0])
	for _, item := range strs {
		minl = min(minl, len(item))

	}

	ans := []byte{}
	for i := 0; i < minl; i++ {
		isMatch := true
		v := strs[0][i]

		for j := 1; j < len(strs); j++ {
			if !isMatch {
				break
			}
			isMatch = strs[j][i] == v
		}
		if !isMatch {
			break
		}
		ans = append(ans, v)
	}
	return string(ans)
}
