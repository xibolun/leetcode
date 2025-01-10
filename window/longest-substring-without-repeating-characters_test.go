package window

import (
	"reflect"
	"testing"

	"github.com/kedadiannao220/leetcode/utils"
)

func TestLegthOfLongestSubstriing(t *testing.T) {
	var source = []struct {
		a   string
		rse int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		// r := lengthOfLongestSubstring2(v.a)
		r := lengthOfLongestSubstring3(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// 3. 无重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/

func lengthOfLongestSubstring(s string) int {
	windows := [128]bool{}
	left := 0
	ans := 0
	for right, c := range s {
		for windows[c] {
			windows[s[left]] = false
			left++
		}
		windows[c] = true
		ans = utils.Max(ans, right-left+1)
	}
	return ans
}

func lengthOfLongestSubstring2(s string) int {
	windows := make(map[string]int)
	left := 0
	ans := 0
	for right, c := range s {
		for {
			if v := windows[string(c)]; v >= 1 {
				windows[string(s[left])] -= 1
				left++
			} else {
				break
			}
		}
		windows[string(c)]++
		ans = utils.Max(ans, right-left+1)
	}
	return ans
}

func lengthOfLongestSubstring3(s string) int {
	windows := make(map[string]bool) // 窗口当中每个字符串出现的次数
	left := 0
	ans := 0 // 要求的是最长的，所以初始值设置为0
	for right, c := range s {
		k := string(c)
		for {
			if v := windows[k]; v {
				delete(windows, string(s[left])) // windows当中保留未出现的数据
				left++
			} else {
				break
			}
		}
		windows[k] = true
		ans = utils.Max(ans, right-left+1)
	}
	return ans
}
