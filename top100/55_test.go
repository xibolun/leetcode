package top100

import (
	"reflect"
	"testing"
)

func TestCanJump(t *testing.T) {
	var source = []struct {
		a   []int
		rse bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
		{[]int{2, 0}, true},
		{[]int{2, 0, 0}, true},
		{[]int{3, 0, 8, 2, 0, 0, 1}, true},
		{[]int{0}, true},
		{[]int{2, 5, 0, 0}, true},
		{[]int{2, 1, 0, 0}, false},
	}

	for _, v := range source {
		r := canJump(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag：贪心
// https://www.bilibili.com/video/BV1VG4y1X7kB/
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	cover := 0
	// 这里需要是cover，这样就会剔除元素为0的情况
	// for i, s := range nums {
	for i := 0; i <= cover; i++ {
		cover = max(cover, nums[i]+i)
		// 这里不能让covert走到底，因为cover有可能>len(nums)
		if cover >= len(nums)-1 {
			return true
		}
	}
	return false
}

func canJump2(nums []int) bool {
	return dfs(0, nums)
}

// 暴力分析
// nums = [2,3,1,1,4]
// 2 跳 2步 得到 1，长度小于length
// 1 跳 1步 得到 1，长度小于length
// 1 跳 1步 得到 4，长度等于length
// 需要对0值进行处理，比如末位都是0值{[]int{2, 5, 0, 0}, true},

// 也可以2跳1步 得到3， 长度小于length
// 3直接到4 长度==length
// 这个算法理解错误，每一次的跳跃可以是 小于nums[i]的任何数，最终可以到达终点即可
func dfs(index int, nums []int) bool {
	if index+1 >= len(nums) {
		return true
	}
	if nums[index] == 0 && nums[index+1] != 0 {
		return false
	}
	step := nums[index]
	// 默认向前走一位
	if step == 0 {
		step = 1
	}
	stepIndex := index + step
	return dfs(stepIndex, nums)
}
