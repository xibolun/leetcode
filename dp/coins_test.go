package dp

import "testing"

func TestCoins(t *testing.T) {
	var source = []struct {
		a     int
		b     int
		c     int
		x     int
		total int
	}{
		{3, 7, 10, 70, 16},
	}

	for _, v := range source {
		res := calc(v.a, v.b, v.c, v.x)
		if res != v.total {
			t.Errorf("test fail,expect %d,but %d", v.total, res)
		}

		val := dp(v.a, v.b, v.c, v.x)
		if val != v.total {
			t.Errorf("test fail,expect %d,but %d", v.total, val)
		}
	}
}
