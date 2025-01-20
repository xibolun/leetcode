package top100

import (
	"reflect"
	"sort"
	"testing"
)

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
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func findKthLargest(nums []int, k int) int {
	if k > len(nums) {
		return 0
	}
	sort.Ints(nums)
	return nums[len(nums)-1-k]
}
