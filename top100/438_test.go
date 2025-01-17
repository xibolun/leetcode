package top100

import (
	"reflect"
	"slices"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	var source = []struct {
		a   string
		p   string
		rse []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	}

	for _, v := range source {
		r := findAnagrams(v.a, v.p)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func findAnagrams(s string, p string) []int {
	sortP := []byte(p)
	slices.Sort(sortP)
	n := len(p)
	l := 0
	mapp := map[string]struct{}{string(sortP): {}}
	var ans []int

	for {
		if l+n > len(s) {
			break
		}
		m := []byte(s[l : l+n])
		slices.Sort(m)
		if _, ok := mapp[string(m)]; ok {
			ans = append(ans, l)
		}
		l++
	}
	return ans

}
