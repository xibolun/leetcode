package top100

import (
	"maps"
	"slices"
)

// 用hash表来解决
func groupAnagrams(strs []string) [][]string {
	ans := make(map[string][]string)

	for _, str := range strs {
		// 字符排序
		b := []byte(str)
		slices.Sort(b)
		if v, ok := ans[string(b)]; ok {
			v = append(v, str)
			ans[string(b)] = v
		} else {
			ans[string(b)] = []string{str}
		}
	}
	return slices.Collect(maps.Values(ans))
}
