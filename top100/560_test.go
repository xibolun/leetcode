package top100

import (
	"reflect"
	"testing"
)

// tag: 前缀和
// similar: 1031/523/304

func TestSubarraySum(t *testing.T) {
	var source = []struct {
		a      []int
		target int
		rse    int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{-1, -1, 1}, 0, 1},
	}

	for _, v := range source {
		r := subarraySum(v.a, v.target)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// https://www.bilibili.com/video/BV1gN411E7Zx/?spm_id_from=333.337.search-card.all.click&vd_source=b5cc04f324fc9d6ee48a5febd77392fc

func subarraySum(nums []int, k int) int {
	// 构建前缀和
	pre := make([]int, len(nums)+1)
	// 这里要构建一个虚拟的，因为元素有可能相同，pre-k有可能会出现多次
	pre[0] = 0
	ans := 0

	for i := 0; i < len(nums); i++ {
		pre[i+1] = pre[i] + nums[i]
	}

	cnt := map[int]int{}
	// 获取前缀和与k的差值出现次数
	for _, s := range pre {
		ans += cnt[s-k]
		cnt[s]++
	}
	return ans
}

func subarraySum2(nums []int, k int) int {
	if len(nums) == 1 {
		if nums[0] == k {
			return 1
		} else {
			return 0
		}
	}
	n := len(nums)
	l, start := 0, 0
	sum := 0
	ans := 0

	for {
		if l > n-1 && start < n-1 {
			start++
			l = start + 1
		}
		if start > n-1 {
			break
		}
		sum = nums[l] + sum
		if sum == k || sum > k {
			ans++
			start = start + 1
			l = start + 1
			sum = nums[start]

			continue
		}

		if sum < k {
			l++
		}
	}
	return ans
}
