package main

import "fmt"

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {

	max := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
