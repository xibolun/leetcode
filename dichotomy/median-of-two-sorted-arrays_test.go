package dichotomy

import "testing"

// https://leetcode.cn/problems/median-of-two-sorted-arrays/description/

func TestFindMedianSortedArrays(t *testing.T) {
	var source = []struct {
		a   []int
		b   []int
		rse float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, v := range source {
		r := findMedianSortedArrays(v.a, v.b)
		if r != v.rse {
			t.Errorf("test fail, result: %f, should be :%f", r, v.rse)
		}
	}
}

// 解法思想在这里
// https://github.com/javasmall/bigsai-algorithm/blob/master/leetcode/problems/LeetCode%2004%E5%AF%BB%E6%89%BE%E4%B8%A4%E4%B8%AA%E6%AD%A3%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E4%B8%AD%E4%BD%8D%E6%95%B0(%E5%9B%B0%E9%9A%BE)%E4%BA%8C%E5%88%86%E6%B3%95.md
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total%2 == 1 {
		return float64(getKth(nums1, nums2, total/2+1))
	} else {
		return float64(getKth(nums1, nums2, total/2)+getKth(nums1, nums2, total/2-1)) / 2
	}
}

func getKth(a, b []int, k int) int {
	if len(a) > len(b) {
		return getKth(b, a, k)
	}
	if len(a) == 0 {
		return b[k-1]
	}
	if k == 1 {
		return min(a[0], b[0])
	}
	pa := min(len(a), k/2)
	pb := k - pa
	if a[pa-1] < b[pb-1] {
		return getKth(a[pa:], b, pb)
	}
	return getKth(a, b[pb:], pa)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
