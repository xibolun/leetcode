package doublepoint

import (
	"reflect"
	"testing"

	"github.com/kedadiannao220/leetcode/utils"
)

func TestMaxArea(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		// {[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		// {[]int{1, 1}, 1},
		// {[]int{0, 2}, 0},
		// {[]int{8, 7, 2, 1}, 7},
		{[]int{1, 8, 100, 2, 100, 4, 8, 3, 7}, 200},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := maxArea(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// 11. 盛最多水的容器
// https://leetcode.cn/problems/container-with-most-water/description/
func maxArea(height []int) int {
	area := 0 // 窗口面积
	l, r := 0, len(height)-1

	for l < r {
		s := (r - l) * utils.Min(height[l], height[r])
		area = utils.Max(area, s)
		if height[l] < height[r] {
			l++
			// 由于还有一个依赖指标为r，所以无法使用此场景来做优化
			// for l > 1 && height[l] < height[l-1] {
			// 	l++
			// }
		} else {
			r--
			// 由于还有一个依赖指标为l，所以无法使用此场景来做优化
			// for r > 0 && height[r] > height[r-1] {
			// r--
			// }
		}
	}
	return area
}
