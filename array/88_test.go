package array

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var source = []struct {
		a   []int
		m   int
		b   []int
		n   int
		rse []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}

	for _, v := range source {
		merge(v.a, v.m, v.b, v.n)
		if !reflect.DeepEqual(v.a, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", v.a, v.rse)
		}
	}
}

// tag: 数组,合并数组,ToTry

// 思想：灵神的倒序，三指针:https://leetcode.cn/problems/merge-sorted-array/solutions/2385610/dao-xu-shuang-zhi-zhen-wei-shi-yao-dao-x-xxkp/
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, p := m-1, n-1, m+n-1
	// 这里的nums1或者nums2有可能是[0],[]这样的特殊情况
	// 所以判断条件需要使用p2，如果p2<0，说明已经合并完了
	// for p >= 0 && p1 >= 0 && p2 >= 0 {
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] >= nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			// p1<0或者 np1 < np2都让p2来赋值
			nums1[p] = nums2[p2]
			p2--
		}

		// if p1 >= 0 && nums1[p1] <= nums2[p2] {
		// nums1[p] = nums2[p2]
		// p2--
		// }

		// v1, v2 := nums1[p1], nums2[p2]
		// if v1 > v2 {
		// 	nums1[p] = v1
		// 	p1--
		// } else {
		// 	nums1[p] = v2
		// 	p2--
		// }
		p--
	}
}

// FIXME 这个算法没有写完
// 思想：比对后合并至nums1
func merge2(nums1 []int, m int, nums2 []int, n int) {
	//暴力解法
	l := 0
	for _, s := range nums2 {
		for l < len(nums1) {
			if nums1[l] > s {
				tmp := nums1[l:]
				nums1 = append(nums1[0:l], s)
				nums1 = append(nums1, tmp...)
			} else {
				l++
			}
		}
	}
}
