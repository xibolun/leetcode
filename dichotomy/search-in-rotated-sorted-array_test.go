package dichotomy

import "testing"

// https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
func TestSearch(t *testing.T) {
	var source = []struct {
		nums   []int
		target int
		rse    int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
	}
	for _, v := range source {
		// r := search(v.nums, v.target)
		// r := search2(v.nums, v.target)
		// r := search3(v.nums, v.target)
		r := search4(v.nums, v.target)
		if r != v.rse {
			t.Errorf("test fail, result: %d, should be :%d", r, v.rse)
		}
	}
}

func findMin(nums []int) int {
	left, right := -1, len(nums)-1 // 开区间 (-1, n-1)
	for left+1 < right {           // 开区间不为空
		mid := left + (right-left)/2
		if nums[mid] < nums[len(nums)-1] {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

// 有序数组中找 target 的下标
func lowerBound(nums []int, left, right, target int) int {
	for left+1 < right { // 开区间不为空
		// 循环不变量：
		// nums[left] < target
		// nums[right] >= target
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid // 范围缩小到 (mid, right)
		} else {
			right = mid // 范围缩小到 (left, mid)
		}
	}
	if nums[right] != target {
		return -1
	}
	return right
}

func search(nums []int, target int) int {
	i := findMin(nums)
	if target > nums[len(nums)-1] { // target 在第一段
		return lowerBound(nums, -1, i, target) // 开区间 (-1, i)
	}
	// target 在第二段
	return lowerBound(nums, i-1, len(nums), target) // 开区间 (i-1, n)
}

func search2(nums []int, target int) int {
	// 递归后的数组里面没有此元素
	if len(nums) == 0 {
		return -1
	}
	mid := getMiddleIndex(nums)
	if mid == 0 {
		return -1
	}
	if nums[mid] == target {
		return mid
	}

	//  右侧为旋转段
	if nums[mid] > nums[len(nums)-1] {
		if nums[mid] < target && target >= nums[0] {
			return search2(nums[:mid-1], target)
		} else {
			return search2(nums[mid+1:], target)
		}
	} else {
		if nums[mid] > target && target <= nums[len(nums)-1] {
			return search2(nums[mid+1:], target)
		} else {
			return search2(nums[:mid-1], target)
		}
	}
}

func getMiddleIndex(nums []int) int {
	l, r := 0, len(nums)-1
	if l <= r {
		return (l + r) / 2
	}
	return 0
}

func search3(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		// 右段为旋转段
		if nums[mid] > nums[r] {
			// 使用l->mid来递归
			if nums[mid] > target && target >= nums[l] {
				r = mid - 1
			} else {
				//  使用mid->r来递归
				l = mid + 1
			}
		} else {
			// 左段为旋转段
			// 使用mid->r来递归
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				// 使用l->mid来递归
				r = mid - 1
			}
		}
	}
	return -1
}
