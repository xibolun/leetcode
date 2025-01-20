package top100

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	var source = []struct {
		a   []int
		rse []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}

	for _, v := range source {
		r := dailyTemperatures(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 数组
// 暴力解法，超出时间限制
func dailyTemperatures(temperatures []int) []int {
	l := 0
	r := l + 1
	ans := make([]int, 0)
	for {
		if l > len(temperatures)-1 {
			break
		}
		if r > len(temperatures)-1 {
			ans = append(ans, 0)
			l++
			r = l + 1
			continue
		}
		if temperatures[r] > temperatures[l] {
			ans = append(ans, r-l)
			l++
			r = l + 1
		} else {
			r++
		}
	}
	return ans
}
