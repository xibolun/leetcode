package top100

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	var source = []struct {
		a   []int
		rse []int
	}{
		// {[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
	}

	for _, v := range source {
		nextPermutation(v.a)
		if !reflect.DeepEqual(v.a, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", v.a, v.rse)
		}
	}
}

func nextPermutation(nums []int) {
	// 倒序找峰值点，左边的即为最大序应该替换的位置
	// index := 0
	// for i := len(nums) - 1; i > 0; i-- {
	// 	if nums[i] > nums[i-1] {
	// 		index = i
	// 	}
	// }

	index := len(nums) - 2
	for index >= 0 && nums[index] >= nums[index+1] {
		index--
	}

	// index右侧的数据都是降序的
	// 找出最大于等于index的节点，进行更换
	if index >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[index] {
			j--
		}
		nums[index], nums[j] = nums[j], nums[index]
	}

	// index右侧的数据需要反转为增序
	l, r := index+1, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
