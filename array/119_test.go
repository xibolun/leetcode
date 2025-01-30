package array

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	var source = []struct {
		a   int
		rse []int
	}{

		{3, []int{1, 3, 3, 1}},
		// {0, []int{1}},
		// {1, []int{1, 1}},
	}

	for _, v := range source {
		r := getRow(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// func getRow(rowIndex int) []int {
// 	res := make([][]int, rowIndex+1)

// 	for i := range res {
// 		res[i] = make([]int, i+1)
// 		res[i][0], res[i][i] = 1, 1
// 		for j := 1; j < i; j++ {
// 			res[i][j] = res[i-1][j-1] + res[i-1][j]
// 		}
// 		if rowIndex == i {
// 			return res[i]
// 		}
// 	}
// 	return []int{}
// }

// tag: 迭代
// similar: 118
// 每行的第一个和最后一个都是1，其他的是上一行的[i-1]+[j-1]
// 1 1
// 1 2 1
// 1 3 3 1
func getRow(rowIndex int) []int {
	ans := make([][]int, rowIndex+1)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0], ans[i][i] = 1, 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
		if rowIndex == i {
			return ans[i]
		}
	}
	return []int{}
}
