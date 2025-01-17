package backtrack

import (
	"reflect"
	"slices"
	"testing"
)

func TestPartition(t *testing.T) {
	var source = []struct {
		a   string
		rse [][]string
	}{
		{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
		{"a", [][]string{{"a"}}},
	}

	for _, v := range source {
		r := partition(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// FIXME  这个算法不仅仅需要回溯，还需要其他来结合，先放一下
func partition(s string) [][]string {
	n := len(s)
	var ans [][]string
	var path []string
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
		}
		for _, c := range s {
			path = append(path, string(c))
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
