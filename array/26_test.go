package array

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := removeDuplicates(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 数组,双指针
func removeDuplicates(nums []int) int {
	k := 1 //因为len(nums)>=1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] { // nums[i] 不是重复项
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
