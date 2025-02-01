package array

import (
	"reflect"
	"testing"
)

func TestMaxScore(t *testing.T) {
	var source = []struct {
		a   []int
		k   int
		rse int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 1}, 3, 12},
		{[]int{2, 2, 2}, 2, 4},
		{[]int{9, 7, 7, 9, 7, 7, 9}, 7, 55},
		{[]int{1, 1000, 1}, 1, 1},
		{[]int{1, 79, 80, 1, 1, 1, 200, 1}, 3, 202},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := maxScore(v.a, v.k)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 滑动窗口
func maxScore(cardPoints []int, k int) int {
	// 滑动窗口大小
	m := len(cardPoints) - k
	// 求滑动窗口大小内的元素和
	tmp := 0
	mins := 0
	// 所有元素的和
	// ans = sum - mins
	sum := 0

	l, r := 0, m-1

	for i := 0; i < len(cardPoints); i++ {
		sum += cardPoints[i]
		if r > len(cardPoints)-1 {
			break
		}
		if i < r {
			tmp += cardPoints[i]
		} else if i == r {
			tmp += cardPoints[i]
			mins = tmp
		} else {
			tmp = tmp + cardPoints[i] - cardPoints[l]
			mins = min(mins, tmp)
			// 向右滑动
			r++
			l++
		}
	}
	return sum - mins
}
