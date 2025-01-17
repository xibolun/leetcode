package top100

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {

	var source = []struct {
		a   []int
		rse []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}

	for _, v := range source {
		moveZeroes(v.a)
		if !reflect.DeepEqual(v.a, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", v.a, v.rse)
		}
	}
}

func moveZeroes(nums []int) {
	n := len(nums)
	l, r := 0, n

	for {
		if l >= r {
			break
		}
		if nums[l] != 0 {
			l++
			continue
		}
		r--
		nums1 := nums[:l]
		nums2 := nums[l+1:]
		nums2 = append(nums2, nums[l])
		nums = append(nums1, nums2...)

	}
}
