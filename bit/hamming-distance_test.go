package bit

import "testing"

// https://leetcode.cn/problems/hamming-distance/

func TestHammingDistance(t *testing.T) {
	var source = []struct {
		a   int
		b   int
		rse int
	}{
		{1, 4, 2},
	}

	for _, v := range source {
		r := hammingDistance(v.a, v.b)
		if r != v.rse {
			t.Errorf("test fail, result: %d, should be :%d", r, v.rse)
		}
	}

}

func hammingDistance(x int, y int) int {
	res := 0
	for i := 0; i < 32; i++ {
		if (x>>i)&1 != (y>>i)&1 {
			res++
		}
	}
	return res
}
