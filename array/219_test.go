package array

import (
	"math"
	"reflect"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	var source = []struct {
		a   []int
		k   int
		rse bool
	}{
		// {[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		// {[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}

	for _, v := range source {
		r := containsNearbyDuplicate(v.a, v.k)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int) // 元素出现的index
	for index, s := range nums {
		v, ok := m[s]
		if !ok {
			m[s] = index
			continue
		}
		if math.Abs(float64(v-index)) <= float64(k) {
			return true
		} else {
			m[s] = index
		}
	}
	return false
}
