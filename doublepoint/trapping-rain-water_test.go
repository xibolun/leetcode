package doublepoint

import (
	"reflect"
	"testing"

	"github.com/kedadiannao220/leetcode/utils"
)

// 42. 接雨水
// https://leetcode.cn/problems/trapping-rain-water/description/
func TestTrap(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}

	for _, v := range source {
		r := trap(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func trap(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	preMax, subMax := 0, 0

	for l <= r {
		preMax = utils.Max(preMax, height[l])
		subMax = utils.Max(subMax, height[r])

		if preMax < subMax {
			ans += preMax - height[l]
			l++
		} else {
			ans += subMax - height[r]
			r--
		}

	}
	return ans

}
