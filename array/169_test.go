package array

import "testing"

func TestMajorityElement(t *testing.T) {
	var source = []struct {
		a   []int
		rse int
	}{
		// {[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		//  因为要求出现次数大于n/2，所以这种情况是不存在的
		// {[]int{2, 3, 2, 4, 5, 6}, 2},
	}

	for _, v := range source {
		r := majorityElement(v.a)
		if r != v.rse {
			t.Errorf("test fail, result: %d, should be :%d", r, v.rse)
		}
	}
}

func majorityElement(nums []int) int {
	n := len(nums)
	ans := nums[0]
	count := 1
	for i := 1; i < n; i++ {
		if count == 0 {
			count = 1
			ans = nums[i]
			continue
		}
		if nums[i] == ans {
			count++
		} else {
			//  count-- 当count为0的时候，说明此时的出现的数值非ans
			//  当count为0的时候，说明ans的数值应该变掉了
			//  不要担心这个count值为负的情况，因为要求是出现次数大于n/2，所以这个count一定是>0的
			count--
		}
	}
	return ans
}
