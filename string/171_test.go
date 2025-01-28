package string

import (
	"math"
	"reflect"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	var source = []struct {
		a   string
		rse int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
	}

	for _, v := range source {
		r := titleToNumber(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func titleToNumber(columnTitle string) int {
	ans := 0
	n := len(columnTitle)
	for i := len(columnTitle) - 1; i >= 0; i-- {
		m := columnTitle[i]
		ans = ans + int(m-64)*int(math.Pow(26, float64(n-1-i)))
	}
	return ans
}
