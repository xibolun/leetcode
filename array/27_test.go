package array

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var source = []struct {
		a   []int
		k   int
		rse int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := removeElement(v.a, v.k)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func removeElement(nums []int, val int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[ans] = nums[i]
			ans++
		}
	}
	return ans
}
