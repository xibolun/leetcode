package points

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
https://leetcode.cn/problems/container-with-most-water/description/

给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。
*/
func solve1(r []int) int {
	var result int
	for left := 0; left < len(r)-1; left++ {
		for right := len(r) - 1; right >= 0; right-- {
			height := getHeight(left, right, r)
			if v := (right - left) * height; result < v {
				result = v
			}
		}
	}

	return result
}

func getHeight(left, right int, r []int) int {
	if r[left] < r[right] {
		return r[left]
	}
	return r[right]
}

func solve2(r []int) int {
	var left, right, result = 0, len(r) - 1, 0

	for {
		if left >= right {
			break
		}

		height := getHeight(left, right, r)
		if v := (right - left) * height; v > result {
			result = v
		}

		if r[left] < r[right] {
			left++
		} else {
			right--
		}

	}

	return result

}

func TestContainerWithMostWater(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(0, solve1([]int{0}))
	ast.Equal(0, solve1([]int{1}))
	ast.Equal(1, solve1([]int{1, 1}))
	ast.Equal(49, solve1([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

	ast.Equal(0, solve2([]int{0}))
	ast.Equal(0, solve2([]int{1}))
	ast.Equal(1, solve2([]int{1, 1}))
	ast.Equal(49, solve2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
