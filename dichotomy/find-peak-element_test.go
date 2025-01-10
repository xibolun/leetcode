package dichotomy

import (
	"reflect"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		// {[]int{1, 2, 3, 1}, 2},
		// {[]int{1, 2, 1, 3, 5, 6, 4}, 5},
		{[]int{2, 1}, 0},
	}

	for _, v := range source {
		r := findPeakElement(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

//  162. 寻找峰值
// https://leetcode.cn/problems/find-peak-element/description/

func findPeakElement(nums []int) int {
	// l, r := 0, len(nums)-1
	// 这里需要开区间，若初始值l为0，则[2,1]这种情形就失效
	l, r := -1, len(nums)-1
	// 这里需要是l+1<r，说明l已经在r的右边，循环结束
	// 若为l<r，则l一直在r的左边，就是死循环
	for l+1 < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			// 峰顶在左侧，所以指针向左移
			r = mid
		} else {
			// 峰顶在右侧，指针右移
			l = mid
		}
	}
	return r
}
