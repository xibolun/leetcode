package array

import (
	"reflect"
	"testing"
)

//  https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/description/

func TestFindDisappearedNumbers(t *testing.T) {
	var source = []struct {
		a   []int
		rse []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{5, 6}},
		{[]int{1, 1}, []int{2}},
	}

	for _, v := range source {
		r := FindDisappearedNumbers(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}

}

func FindDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		//  循环，找到正确的位置
		//  nums经过循环数据为： []int len: 8, cap: 8, [1,2,3,4,3,2,7,8]
		for nums[i] != i+1 {
			if nums[nums[i]-1] == nums[i] {
				break
			}
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	var res []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}
