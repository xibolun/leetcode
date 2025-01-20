package top100

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	var source = []struct {
		a   []int
		rse []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}

	for _, v := range source {
		sortColors(v.a)
		if !reflect.DeepEqual(v.a, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", v.a, v.rse)
		}
	}
}

func sortColors(nums []int) {
	count0 := 0
	count2 := len(nums) - 1
	for i, s := range nums {
		// 将0都移动到前面
		if s == 0 {
			nums[i], nums[count0] = nums[count0], nums[i]
			count0++
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		// 将2都移动到后面
		if nums[i] == 2 {
			nums[i], nums[count2] = nums[count2], nums[i]
			count2--
		}
	}

	// 中间部分都是1
}
