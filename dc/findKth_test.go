package dc

import "testing"

func TestFindKth(t *testing.T) {
	var source = []struct {
		a   []int
		n   int
		k   int
		rse int // 结果
	}{
		{[]int{1, 3, 5, 2, 2}, 5, 3, 2},
	}

	for _, v := range source {
		r := findKth(v.a, v.n, v.k)
		if r != v.rse {
			t.Errorf("test fail, result: %d, should be :%d", r, v.rse)
		}
	}
}

func findKth(a []int, n int, K int) int {
	// write code here
	// write code here
	// write code here
	mid, i := a[0], 1
	first, tail := 0, n-1
	for first < tail {
		if a[i] < mid {
			a[i], a[tail] = a[tail], a[i]
			tail--
		} else {
			a[i], a[first] = a[first], a[i]
			first++
			i++
		}
	}
	a[first] = mid

	if K-1 == first {
		return mid
	}
	if K-1 < first {
		return findKth(a[:first], first, K)
	}
	return findKth(a[first+1:], n-first-1, K-first-1)
}
