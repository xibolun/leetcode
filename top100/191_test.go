package top100

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	var source = []struct {
		a   int
		rse int
	}{
		{11, 3},
		{128, 1},
		{2147483645, 30},
	}

	for _, v := range source {
		r := hammingWeight(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func hammingWeight(n int) int {
	b := fmt.Sprintf("%b", n)
	ans := 0
	for _, s := range b {
		if s == '1' {
			ans++
		}
	}
	return ans
}
