package top100

import (
	"reflect"
	"testing"
)

func TestFirstMissingPositive(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		// {[]int{1, 2, 0}, 3},
		// {[]int{3, 4, -1, 1}, 2},
		// {[]int{7, 8, 9, 11, 12}, 1},
		{[]int{1}, 2},
	}

	for _, v := range source {
		r := firstMissingPositive(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: array, hard
// 构建下标与下标所对应的数值相等的数组，若index i对应的nums[i]==0说明这个i即为丢失的正整数
func firstMissingPositive(nums []int) int {
	// ans := make([]int, len(nums)+1)
	// 这里不是len(nums)+1的原因是为了计算[1]单调递增的数组，所以是需要多加两位，ans[0]=0
	ans := make([]int, len(nums)+2)

	for i := 0; i < len(nums); i++ {
		// nums[i]若为负数，则过滤掉
		if nums[i] < 0 {
			continue
		}
		// 若数字太大，肯定存在丢失的正整数，就不写入了
		if nums[i] >= len(nums)+1 {
			continue
		}
		ans[nums[i]] = nums[i]
	}

	for i := 1; i < len(ans); i++ {
		if ans[i] == 0 {
			return i
		}
	}
	return 0

}
