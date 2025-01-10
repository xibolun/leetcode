package doublepoint

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	var source = []struct {
		a   []int
		rse [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}

	for _, v := range source {
		r := threeSum(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// 15. 三数之和
// https://leetcode.cn/problems/3sum/description/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	n := len(nums)
	if n < 3 {
		return ans
	}
	// 优化：若最小的三个数和大于0，则不存在这样的数组
	if nums[0]+nums[1]+nums[2] > 0 {
		return ans
	}
	// 优化：若最大的三个数和小于0，则不存在这样的数组
	if nums[n-3]+nums[n-2]+nums[n-1] < 0 {
		return ans
	}

	for i, c := range nums {
		// 这里说明到头了，留两个给左右指针变量
		if i >= len(nums)-2 {
			break
		}
		// 和上一个元素一样，则忽略，因为要求是不重复的数组
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 优化：若最小的三数都小于0，则这个递归就不存在满足条件的数组
		if c+nums[i+1]+nums[i+2] > 0 { // 优化一
			break
		}
		// 优化：若最大的三数都小于0，则这个递归就不存在满足条件的数组
		if c+nums[n-2]+nums[n-1] < 0 { // 优化二
			continue
		}
		//  创建左、右指针，让其相向
		l, r := i+1, len(nums)-1
		// 当左、右指针相向后，才需要对i进行递归
		for l < r {
			s := nums[l] + nums[r] + c
			if s > 0 {
				r--
				continue
			}
			if s < 0 {
				l++
				continue
			}
			ans = append(ans, []int{nums[l], nums[r], c})
			//  右指针左移时，若有重复的跳过
			for r--; r > l && nums[r] == nums[r+1]; r-- {
			}
			// 左指针右移时，若有重复的跳过
			for l++; l < r && nums[l] == nums[l-1]; l++ {
			}
		}

	}
	return ans
}
