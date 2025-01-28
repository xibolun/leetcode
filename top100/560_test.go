package top100

import (
	"reflect"
	"testing"
)

// tag: 前缀和, 类似303

func TestSubarraySum(t *testing.T) {
	var source = []struct {
		a      []int
		target int
		rse    int
	}{
		// {[]int{1, 1, 1}, 2, 2},
		// {[]int{1, 2, 3}, 3, 2},
		{[]int{-1, -1, 1}, 0, 2},
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
	ans := 0
	pre := make([]int, len(nums)+1)

	// 先计算出前缀和
	// 默认添加一个0出现1的前缀和，防止出现重复元素
	pre[0] = 0
	for i := 0; i < len(nums); i++ {
		pre[i+1] = pre[i] + nums[i]
	}

	// 再算出和为k的区间列表
	// pre[j] - pre[k] = nums[j-1]+...nums[k]
	// 前缀和当中的 pre[j] - pre[k]的值即为[j,k]中间的元素和

	// 用一个hash表来存储值
	cnt := map[int]int{}
	for i := 0; i < len(pre); i++ {
		ans += cnt[pre[i]-k] //pre[i]-k若在hash表当中存在，则说明存在连续的区间和为k，这个是一个特性
		cnt[pre[i]]++
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
