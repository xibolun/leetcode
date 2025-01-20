package top100

import (
	"reflect"
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	var source = []struct {
		gas  []int
		cost []int
		rse  int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
	}

	for _, v := range source {
		r := canCompleteCircuit(v.gas, v.cost)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 贪心, #55
func canCompleteCircuit(gas []int, cost []int) int {
	// 加油-消耗 = 剩余油

	// 若存在剩余油的和 >0，则最左侧的数值即为起始点
	curSum := 0
	totalSum := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			// 小于0，说明一定不从i开始
			start = i + 1
			curSum = 0
		}
	}
	// 剩余油的和 <0 ，说明根本跑不了一圈
	if totalSum < 0 {
		return -1
	}
	return start
}
