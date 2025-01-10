package window

import (
	"reflect"
	"testing"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	var source = []struct {
		a      []int
		target int
		rse    int
	}{
		{[]int{10, 5, 2, 6}, 100, 8},
		// 下面这两种情况需要单独考虑和处理
		{[]int{1, 2, 3}, 0, 0},
		{[]int{1, 1, 1}, 1, 0},
	}

	for _, v := range source {
		r := numSubarrayProductLessThanK2(v.a, v.target)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	windows := 1 // 窗口当中的乘积
	left := 0
	value := make([][]int, 0)

	//  数据当中不可能会存在乘积为0的数据
	if k <= 1 {
		return 0
	}

	for right, c := range nums {
		for windows*c >= k {
			windows = windows / nums[left]
			left++
		}
		windows *= c
		child := make([]int, 0)
		for i := left; i <= right; i++ {
			child = append(child, nums[i])
			value = append(value, child)
		}

	}
	return len(value)
}

func numSubarrayProductLessThanK2(nums []int, k int) int {
	windows := 1 // 窗口当中的乘积
	left := 0
	ans := 0

	//  数据当中不可能会存在乘积为0的数据
	if k <= 1 {
		return ans
	}

	for right, c := range nums {
		windows *= c
		for windows >= k {
			windows = windows / nums[left]
			left++
		}
		if windows < k {
			ans += (right - left + 1)
		}
	}
	return ans
}
