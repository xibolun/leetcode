package backtrack

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	var source = []struct {
		a   int
		rse []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, v := range source {
		r := generateParenthesis(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 回溯
func generateParenthesis(n int) []string {
	var ans []string
	max := n * 2 // 需要循环2n次
	path := make([]byte, max)
	var dfs func(i, open int)
	dfs = func(i, open int) {
		// 到达叶子节点
		if i == max {
			ans = append(ans, string(path))
		}

		if open < n {
			path[i] = '('
			dfs(i+1, open+1)
		}
		// if i-open<n{
		// 这里为何不是小于n，因为需要与左括号open相匹配
		if i-open < open {
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return ans
}
