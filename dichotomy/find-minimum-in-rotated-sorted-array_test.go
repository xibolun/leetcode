package dichotomy

import (
	"reflect"
	"testing"
)

func TestFindMinInRotatedSortedArray(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{1}, 1},
	}

	for _, v := range source {
		r := findMinInRotatedSortedArray(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// 153. 寻找旋转排序数组中的最小值
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/
// 数组特点
//  1. nums[0]>nums[n-1]
//  2. 局部有序
//  3. nums[mid] <nums[n-1]:说明
func findMinInRotatedSortedArray(nums []int) int {
	// l, r := 0, len(nums)-1
	// 这里需要开区间，若初始值l为0，则[2,1]这种情形就失效
	l, r := -1, len(nums)-1
	// 这里需要是l+1<r，说明l已经在r的右边，循环结束
	// 若为l<r，则l一直在r的左边，就是死循环
	for l+1 < r {
		mid := (l + r) / 2
		if nums[mid] > nums[len(nums)-1] {
			//最小值在左侧，指针右移
			l = mid
		} else {
			// 峰顶在右侧，指针右移
			r = mid
		}
	}
	return nums[r]
}
