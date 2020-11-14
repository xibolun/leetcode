package dp

import (
	"fmt"
	"testing"
)

func TestLCS(t *testing.T) {
	// 求公有的字符串长度
	fmt.Printf("lcs length %d\n", LCSLength("12345EF", "1AB2345CD"))
	// 求公有的字符串
	fmt.Printf("lcs value %s\n", LCSValue("12345EF", "1AB2345CD"))
	// 求连续的子串
	fmt.Printf("lcs value %s\n", LCSValue2("12345EF", "1AB2345CD"))

}

func LCSLength(str1 string, str2 string) int {
	dp := GetDpArray(str1, str2)
	return dp[len(str1)][len(str2)]
}

/**
题目描述
给定两个字符串str1和str2,输出两个字符串的公共子串
示例1
输入

"1AB2345CD","12345EF"
返回值 "12345"
*/
func LCSValue(str1 string, str2 string) string {
	dp := GetDpArray(str1, str2)
	var s string
	for i := range dp {
		if i == 0 {
			continue
		}
		for j := range dp[i] {
			if j == 0 {
				continue
			}
			if dp[i][j] > dp[i-1][j] && dp[i][j] > dp[i][j-1] {
				// 这里的i对应的string里面乃是+1的，因为前置和左置先放了一层0
				s += string(str1[i-1])
			}
		}
	}
	return s
}

// GetDpArray 构建DP数组: http://note.youdao.com/s/15Ay0xyR
func GetDpArray(str1, str2 string) [10][10]int {
	var dp [10][10]int

	// put 0
	for i := range str1 {
		dp[i][0] = 0
	}
	for j := range str2 {
		dp[0][j] = 0
	}
	// put values
	for i, v := range str1 {
		for j, a := range str2 {
			// 对角线
			if v == a {
				dp[i+1][j+1] = dp[i][j+1] + 1
			}
			// 左邻元素与上邻元素取最大值
			if v != a {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	//outputDP(dp)
	return dp
}

// LCSValue2 构建DP数组: http://note.youdao.com/s/15Ay0xyR
// 这个与LCS不同的是，需要连续的子串，
// 需要设置一个单独的变量来存储
func LCSValue2(str1, str2 string) string {
	var dp [10][10]int

	var s string
	// put 0
	for i := range str1 {
		dp[i][0] = 0
	}
	for j := range str2 {
		dp[0][j] = 0
	}
	// put values
	for i, v := range str1 {
		for j, a := range str2 {
			// 对角线
			if v == a {
				// TODO 若对角线不连续，则需要将s置空
				dp[i+1][j+1] = dp[i][j+1] + 1
				s += string(str1[i])
			}
			// 若要求的是公共的子串，那么这个DP构建的时候，不相等直接为0即可
			if v != a {
				dp[i+1][j+1] = 0
			}
		}
	}

	outputDP(dp)
	return s
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func outputDP(dp [10][10]int) {
	for i := range dp {
		fmt.Println(dp[i])
	}
}
