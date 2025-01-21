package bfs

import (
	"reflect"
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	var source = []struct {
		a   [][]int
		rse int
	}{
		{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		// {[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{[][]int{{0, 2}}, 0},
	}

	for _, v := range source {
		r := orangesRotting(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 图,BFS

type pair struct{ x, y int } // pair坐标

var directory = []pair{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} // 四个方向，每次走一步

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0]) //行、列总数
	fresh := 0                      //新鲜个数
	q := make([]pair, 0)            //腐烂列表
	ans := 0                        // 经过的分钟数
	for i, row := range grid {
		for j, x := range row {
			// 新鲜
			if x == 1 {
				fresh++
			}

			// 腐烂
			if x == 2 {
				q = append(q, pair{i, j}) // 收集腐烂的数据
			}
		}
	}
	for fresh > 0 && len(q) > 0 {
		ans++ //循环一次就是+1分钟
		tmp := q
		q = []pair{} //新被感染的列表
		for _, p := range tmp {
			for _, d := range directory {
				i, j := p.x+d.x, p.y+d.y
				// 若为新鲜橘子，需要被感染
				// 注意这里的边界条件
				if 0 <= i && i < m && 0 <= j && j < n && grid[i][j] == 1 {
					fresh--
					grid[i][j] = 2
					// 新的腐烂橘子列表，用于下次感染源
					q = append(q, pair{i, j})
				}
			}
		}
	}
	if fresh <= 0 {
		return ans
	}
	return -1
}
