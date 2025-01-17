package top100

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {

	var source = []struct {
		a      []int
		target int
		rse    []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}

	for _, v := range source {
		rotate(v.a, v.target)
		if !reflect.DeepEqual(v.a, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", v.a, v.rse)
		}
	}
}

// tag: 双指针
func rotate(nums []int, k int) {
	l, r := 0, len(nums)-1
	cnt := 0
	for {
		if cnt >= k {
			break
		}
		n1 := nums[r]
		n2 := nums[l:r]
		nums = []int{n1}
		nums = append(nums, n2...)
		cnt++
	}
	fmt.Printf("%v", nums)
}
