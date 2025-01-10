package array

import (
	"reflect"
	"sort"
	"testing"
)

//  https://leetcode.cn/problems/next-permutation/

func TestNextPermutation(t *testing.T) {
	var source = []struct {
		a   []int
		rse []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1, 3, 5, 4, 3, 2, 1}, []int{1, 4, 1, 2, 3, 3, 5}},
	}

	for _, v := range source {
		nextPermutation(v.a)
		if !reflect.DeepEqual(v.a, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", v.a, v.rse)
		}
	}
}

func nextPermutation(nums []int) {
	k := 0
	//  找出第一个逆序
	//  k为逆序开始的下标
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			k = i
			break
		}
	}
	//  说明原来的序列是从最大到小排列，找不到逆序点
	if k == 0 {
		sort.Ints(nums)
		return
	}

	//  将逆序的最小元素和逆序开始的元素交换
	for i := len(nums) - 1; i >= k; i-- {
		if nums[i] > nums[k-1] {
			nums[i], nums[k-1] = nums[k-1], nums[i]
			break
		}
	}
	//  将逆序的元素从大到小排序
	sort.Ints(nums[k:])
}
