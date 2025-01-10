package array

import (
	"testing"

	"github.com/kedadiannao220/leetcode/utils"
)

// https://leetcode.cn/problems/kth-largest-element-in-an-array/description/

func TestFindKthLargest(t *testing.T) {
	var source = []struct {
		a   []int
		k   int
		rse int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}

	for _, v := range source {
		r := findKthLargest(v.a, v.k)
		if r != v.rse {
			t.Errorf("test fail, result: %d, should be :%d", r, v.rse)
		}
	}
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k)
}

func quickSelect(nums []int, left, right, k int) int {
	if left >= right {
		return nums[left]
	}
	//  先排个序
	utils.QuickSort(nums, left, right)
	//  取第k下标数值
	return nums[len(nums)-k]
}
