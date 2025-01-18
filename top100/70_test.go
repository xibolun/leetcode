package top100

import (
	"reflect"
	"testing"
)

func TestClimbStairs(t *testing.T) {

	var source = []struct {
		a   int
		rse int
	}{
		{2, 2},
		{3, 3},
	}

	for _, v := range source {
		climbStairs(v.a)
		if !reflect.DeepEqual(v.a, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", v.a, v.rse)
		}
	}
}

// 递归
// jump 0楼为0
// jump 1楼为1
// jump 2楼为2 1,1  2
// jump 3楼：若先走1步，即为jump2，若先走2步，即为jump1，所以为jump2+jump1
// jump 4楼为，若先走1步，即为jump3，若先走2步，即为求jump2，所以为jump3+jump2
// jump 5楼：若先走1步，即为jump4，若先走2步，即为求jump3，所以为jump4+jump3

func dp(i int) int {
	if i <= 1 {
		return 1
	}
	return dp(i-1) + dp(i-2)
}

func climbStairs(n int) int {
	return dp(n)
}
