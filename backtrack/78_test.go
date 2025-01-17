package backtrack

import (
	"reflect"
	"slices"
	"testing"
)

func TestSubsets(t *testing.T) {
	var source = []struct {
		a   []int
		rse [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{0}, [][]int{{}, {0}}},
	}

	for _, v := range source {
		r := subsets(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func subsets(nums []int) [][]int {
	var ans [][]int
	var path []int // 叶子节点
	n := len(nums) // 树的高度，递归次数
	var dfs = func(i int) {}
	dfs = func(i int) {
		// 复制一下path，因为path的数据被下面的操作影响
		// 这里不只包括叶子节点，所以不需要使用i==n的边界条件
		ans = append(ans, slices.Clone(path))
		// 这里的for循环是从i开始的
		for j := i; j < n; j++ {
			path = append(path, nums[j])
			// dfs(i+1) 和dfs(j+1有何区别)
			dfs(j + 1)
			// 思考？这里为何要取n-1个数值
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
