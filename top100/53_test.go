package top100

import (
	"math"
	"reflect"
	"testing"
)

// tag: 前缀和

func TestMaxSubArray(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		// {[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{-1}, -1},
		{[]int{-1, -2}, -1},
		// {[]int{-1, 2}, 2},
	}

	for _, v := range source {
		// r := maxSubArray(v.a)
		r := maxSubArray1(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func maxSubArray(nums []int) int {
	preSum := 0
	ans := math.MinInt
	// 最小前缀和
	minPreSum := 0
	for _, s := range nums {
		// 维护出当前的前缀和
		preSum = preSum + s
		// 计算前缀和的差值，差值越大，说明此时的数组之和是越大的
		ans = max(ans, preSum-minPreSum)
		minPreSum = min(minPreSum, preSum)
	}
	return ans
}

// FIXME 这个算法的问题在于，前缀和默认第一个元素为0，若数组当中都为负数的时候，就会不正确
func maxSubArray1(nums []int) int {
	// 构建前缀和数组
	preSum := make([]int, len(nums)+1)
	for i, s := range nums {
		preSum[i+1] = preSum[i] + s
	}
	// pj-pk表示k到j区间的数组之和，
	// 若想取得最大值，那么就看前缀和数组当中最大的差值是多少即可
	// 相当于找出最大的前缀和、最小的前缀和
	minPreSum, ans := preSum[0], math.MinInt
	for _, s := range preSum {
		minPreSum = min(minPreSum, s)
		ans = max(ans, s-minPreSum)
	}
	return ans
}

// func maxSubArray(nums []int) int {
// 	sum := nums[0]
// 	total := 0
// 	n := len(nums)

// 	l := 1
// 	start := l

// 	for {
// 		if l > n {
// 			break
// 		}
// 		sum = sum + nums[l]
// 		if sum >= total {
// 			l++
// 		} else {
// 			l = start + 1
// 		}
// 	}
// 	return sum
// }
