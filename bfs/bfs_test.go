package bfs

import (
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {
	var grid = [][]byte{
		{1, 1, 0, 0, 0},
		{0, 1, 0, 1, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1},
	}

	fmt.Println(dfs1(grid))
}

/**
	岛屿问题

题目描述
给一个01矩阵，1代表是陆地，0代表海洋， 如果两个1相邻，那么这两个1属于同一个岛。我们只考虑上下左右为相邻。
岛屿: 相邻陆地可以组成一个岛屿（相邻:上下左右） 判断岛屿个数。

[[1,1,0,0,0],[0,1,0,1,1],[0,0,0,1,1],[0,0,0,0,0],[0,0,1,1,1]]  ---> 3
*/

func solve(grid [][]byte) int {
	return dfs1(grid)
}

/**
思路1，若有一个1，说明肯定是一个岛；那就把周边的1全都清0即可
*/
func dfs1(grid [][]byte) int {
	x := len(grid)
	y := len(grid[0])

	var lands int
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if grid[i][j] == 1 {
				fmt.Printf("第 %d 行 第 %d 列, value is %d\n", i, j, grid[i][j])
				lands++
				set0(grid, i, j)
			}
		}
	}
	return lands
}

func set0(grid [][]byte, i, j int) {

	// 自己清0
	grid[i][j] = 0
	// 周边清0
	// 左边清0
	if i-1 > 0 && grid[i-1][j] == 1 {
		set0(grid, i-1, j)
	}
	// 右边清0
	if i+1 < len(grid) && grid[i+1][j] == 1 {
		set0(grid, i+1, j)
	}
	// 上边清0
	if j-1 > 0 && grid[i][j-1] == 1 {
		set0(grid, i, j-1)
	}
	// 下边清0
	if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
		set0(grid, i, j+1)
	}

	//// or write is like this;
	//if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
	//	return
	//}
	//// 若不加这个判断，会一直进行递归
	//if grid[i][j] == 1 {
	//	grid[i][j] = 0
	//	set0(grid, i-1, j)
	//	set0(grid, i+1, j)
	//	set0(grid, i, j-1)
	//	set0(grid, i, j+1)
	//	return
	//}

}

/**
这种做法是将上、下、左、右进行了坐标化
*/
func solve2(grid [][]byte) (ans int) {
	dirs := [][]int{
		[]int{-1, 0},
		[]int{1, 0},
		[]int{0, -1},
		[]int{0, 1},
	}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			dfs(x, y)
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return
}
