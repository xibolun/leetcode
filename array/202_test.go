package array

import (
	"reflect"
	"testing"
)

func TestIsHappy(t *testing.T) {
	var source = []struct {
		a   int
		rse bool
	}{
		{19, true},
		{2, false},
	}

	for _, v := range source {
		r := isHappy(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

func isHappy(n int) bool {
	fast, slow := n, n
	for {
		fast = sqrtV(fast)
		fast = sqrtV(fast)
		slow = sqrtV(slow)
		if fast == slow {
			break
		}
	}
	return slow == 1
}

func sqrtV(n int) int {
	sum := 0
	for n > 0 {
		v := n % 10
		sum += v * v
		n = n / 10
	}
	return sum
}

// 最终会得到 1。
// 最终会进入循环。
// 值会越来越大，最后接近无穷大。
// tag: 快乐数

// FIXME
func isHappy2(n int) bool {
	arr := make([]int, 0)
	buildArray(n, arr)

	if v := sqrt(arr); v == 1 {
		return true
	} else {
		return isHappy2(v)
	}
}

func buildArray(n int, arr []int) {
	arr = append(arr, n%10)
	if n/10 > 0 {
		buildArray(n, arr)
	}
}

func sqrt(arr []int) int {
	ans := 0
	for _, item := range arr {
		ans += item * item
	}
	return ans
}
