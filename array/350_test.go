package array

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	var source = []struct {
		a   []int
		b   []int
		rse []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := intersect(v.a, v.b)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	ans := make([]int, 0)

	for _, s := range nums1 {
		// v := m[s]
		// v++
		// m[s] = v
		m[s]++
	}

	for _, s := range nums2 {
		// v, ok := m[s]
		// if !ok {
		// 	continue
		// }
		// v--
		// if v == 0 {
		// 	ans = append(ans, s)
		// }
		// m[s] = v
		if m[s] > 0 {
			ans = append(ans, s)
			m[s]--
		}
	}
	return ans
}
