package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	var source = []struct {
		a   []int
		rse []string
	}{
		// {[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6->6", "8->9"}},
	}

	for _, v := range source {
		// r := lengthOfLongestSubstring(v.a)
		r := summaryRanges(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func summaryRanges(nums []int) []string {
	if len(nums) <= 0 {
		return []string{}
	}
	l := 0
	ans := make([]string, 0)
	tmp := make([]int, 0)
	tmp = append(tmp, nums[l])
	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[l] {
			tmp = append(tmp, nums[i])
		} else {
			ans = buildAns(tmp, ans)
			tmp = []int{nums[i]}
		}
		l = i
	}
	ans = buildAns(tmp, ans)

	return ans
}
func buildAns(tmp []int, ans []string) []string {
	if len(tmp) > 1 {
		ans = append(ans, fmt.Sprintf("%d->%d", tmp[0], tmp[len(tmp)-1]))
	}
	if len(tmp) == 1 {
		ans = append(ans, fmt.Sprintf("%d", tmp[0]))
	}
	return ans
}
