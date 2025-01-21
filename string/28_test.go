package string

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	var source = []struct {
		a   string
		b   string
		rse int
	}{
		{"sadbutsad", "sad", 0},
		// {"aabaabaaf", "aabaaf", 3},
		{"leetcode", "leeto", -1},
	}

	for _, v := range source {
		r := strStr(v.a, v.b)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 子串,KMP
// 算法思想：https://www.bilibili.com/video/BV1AY4y157yL/?spm_id_from=333.337.search-card.all.click
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	next := buildNext(needle)

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && needle[j] != haystack[i] {
			j = next[j-1]
		}

		if needle[j] == haystack[i] {
			j++
		}

		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func buildNext(needle string) []int {
	m := len(needle)
	ans := make([]int, m)

	// 这里的技巧i是从1开始，因为ans[i]一定是为0的
	for i, j := 1, 0; i < m; i++ {
		// 每次遍历一下
		for j > 0 && needle[j] != needle[i] {
			j = ans[j-1]
		}

		// 相同前后缀的长度
		if needle[i] == needle[j] {
			j++
		}

		ans[i] = j
	}
	return ans
}
