package backtrack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	a := "23"
	fmt.Println(a[0])
	fmt.Println(a[0] - '0')

	var source = []struct {
		a   string
		rse []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	}

	for _, v := range source {
		r := letterCombinations(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// [17. 电话号码的字母组合](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/)
var mapping = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) (ans []string) {
	n := len(digits)
	if n == 0 {
		return
	}
	path := make([]byte, n) // 注意 path 长度一开始就是 n，不是空列表
	var dfs func(int)
	dfs = func(i int) {
		// 叶子节点
		if i == n {
			ans = append(ans, string(path))
			// return后即开始回溯
			return
		}
		// digits[i]-'0' 将ascii转为string
		// 否则digits[i]为ascii，比如2的ascii为50
		for _, c := range mapping[digits[i]-'0'] {
			// 这个for循环就是需要选择的次序列表
			path[i] = byte(c) // 直接覆盖
			// 下一层开始递归的位置
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}
