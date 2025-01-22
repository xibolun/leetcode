package array

import (
	"reflect"
	"sort"
	"testing"
)

func TestMaxCoins(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		{[]int{2, 4, 1, 2, 7, 8}, 9},
		{[]int{2, 4, 5}, 4},
		{[]int{9, 8, 7, 6, 5, 1, 2, 3, 4}, 18},
	}

	for _, v := range source {
		r := maxCoins(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: array
// 思路：先排序，只用取最大值左则的数即可，即3*n-1的数值进行累加
func maxCoins(piles []int) int {
	sort.Ints(piles)
	n := len(piles)
	times := n / 3
	ans := 0
	for i := n - 2; i >= times; i = i - 2 {
		ans += piles[i]
	}
	return ans
}
