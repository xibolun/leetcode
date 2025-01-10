package array

import (
	"reflect"
	"testing"
)

//  https://leetcode.cn/problems/rotate-image/description/

func TestRotate(t *testing.T) {
	var source = []struct {
		a   [][]int
		rse [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
	}

	for _, v := range source {
		rorate(v.a)
		if !reflect.DeepEqual(v.a, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", v.a, v.rse)
		}
	}
}

func rorate(matrix [][]int) {
	n := len(matrix)
	// 第一行和最后一行交换
	// 第二行和倒数第二行交换
	// .....
	// 得到如下数据
	// [ [7,8,9],
	//   [4,5,6],
	//   [1,2,3]]
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}

	// 将上面的交换结果进行对角线交换
	// [[7,4,1],
	// [8,5,2],
	// [9,6,3]]
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func TestRotate2(t *testing.T) {
	var source = []struct {
		a   [][]int
		rse [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{3, 6, 9}, {2, 5, 8}, {1, 4, 7}}},
	}

	for _, v := range source {
		rorate2(v.a)
		if !reflect.DeepEqual(v.a, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", v.a, v.rse)
		}
	}
}

// 题目是顺时针，如果是逆时针，只需要先换列，再对角换即可
func rorate2(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
