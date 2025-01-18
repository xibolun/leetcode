package array

import (
	"reflect"
	"testing"
)

//  https://leetcode.cn/problems/product-of-array-except-self/description/
//  https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof/description/

func TestProductExceptSelf(t *testing.T) {
	var source = []struct {
		a   []int
		rse []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}

	for _, v := range source {
		// r := productExceptSelf(v.a)
		r := productExceptSelf2(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}
func productExceptSelf(nums []int) []int {
	ansi := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		v := 1
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			v *= nums[j]
		}
		ansi[i] = v
	}
	return ansi
}

// https://leetcode.cn/problems/product-of-array-except-self/solutions/1359641/by-tonngw-l4xc/
func productExceptSelf2(nums []int) []int {
	ansi := make([]int, len(nums))

	// 对于每一个数的结果都可以分成两部分的乘积，左边所有数的乘积 * 右边所有数的乘积

	// 左边的乘积
	// 正序
	for i, p := 0, 1; i <= len(nums)-1; i++ {
		ansi[i] = p
		p *= nums[i]
	}

	// 右边的乘积
	// 倒序
	for i, p := len(nums)-1, 1; i >= 0; i-- {
		ansi[i] *= p
		p *= nums[i]
	}
	return ansi
}
