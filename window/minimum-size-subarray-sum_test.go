package window

import (
	"reflect"
	"testing"

	"github.com/kedadiannao220/leetcode/utils"
)

func TestMinSubArrayLen(t *testing.T) {
	var source = []struct {
		a      []int
		target int
		rse    int
	}{
		{[]int{2, 3, 1, 2, 4, 3}, 7, 2},
		{[]int{1, 4, 4}, 4, 1},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1}, 11, 0},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := minSubArrayLen(v.target, v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// 209. 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum/description/

func minSubArrayLen(target int, nums []int) int {
	windows := 0 // 窗口当中的数组和
	left := 0
	ans := len(nums) + 1 // 设置一个最大值，用于进行比较取小
	for right, c := range nums {
		windows += c                       // 窗口右指针扩张
		for windows-nums[left] >= target { // 窗口左指针缩小，至到<=target，因为和为target才是我们想要的窗口
			windows -= nums[left]
			left++
		}
		if windows >= target { // 当窗口为target的时候，我们就找到了一个满足的区间，要看一下长度是不是最小的
			ans = utils.Min(ans, right-left+1)
		}
	}
	//  若是没有找到，则ans还是原始值
	if ans < len(nums)+1 {
		return ans
	}
	return 0
}
