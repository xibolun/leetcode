package dp

//评测题目
/*
 a，b，c 三种币值无限供应，要凑齐X元，有几种办法。输出多少种即可，不需要给出具体的拼凑方式
 比如a=3, b=7, c=10, X=70。可以有如下的拼凑方法：

7个10元
或者 10个7元
或者 1个10元 & 20个3元
或者 ……
*/

var coins []int

// 第二种求法，动态规划： http://note.youdao.com/s/15Ay0xyR
func dp(a, b, c, x int) int {
	coins = []int{a, b, c}
	// 第一步，先构建数组
	dp := make([]int, x+1)
	dp[0] = 1

	// 第二步，循环i，当值相等时+1，当值不相等时，取左边和上边的元素的最大值
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= x; j++ {
			dp[j] = (dp[j] + dp[j-coins[i]]) % x
		}
	}
	return dp[x]
}

// 第一种求法，暴力求解
func calc(a, b, c, x int) int {
	coins = []int{a, b, c}
	total := 0
	for j := 0; j < len(coins); j++ {
		if x%coins[j] == 0 {
			total++
		}
		for i := 0; i < x/getMin(); i++ {
			if (x-x%coins[i])%coins[getNext(i)] == 0 || (x-x%coins[i])%coins[getPre(i)] == 0 {
				total++
			}

			if (x-(x-x%coins[i])%coins[getNext(i)])%coins[getNext(getNext(i))] == 0 || (x-(x-x%coins[i])%coins[getPre(i)])%coins[getPre(getPre(i))] == 0 {
				total++
			}
		}
	}
	return total
}

func getNext(i int) int {
	if i <= len(coins)-1 {
		return i + 1
	}
	return 0
}

func getPre(i int) int {
	if i <= 0 {
		return len(coins) - 1
	}
	return i - 1
}

func getMin() int {
	for i := range coins {
		return Min(coins[i], coins[getNext(i)])
	}
	return coins[0]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
