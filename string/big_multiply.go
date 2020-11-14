package string

import "fmt"

// 大数相乘
// 给两个小于1w长度的string
//
//```
//a=789
//b=987
//```
//
//|      | a0-9  | a1-8  | a2-7  |
//| ---- | ----- | ----- | ----- |
//| b0-7 | c0=63 | c1=56 | c2=49 |
//| b1-8 | c1=72 | c2=64 | c3=56 |
//| b2-9 | c2=81 | c3=72 | c4=63 |
//
//
func BigMultiply1(str1, str2 []int) []int {
	// 结果不会超过len(str1) +len(str2)
	res := make([]int, len(str1)+len(str2))

	for i := range str1 {
		for j := range str2 {
			// 数组当中重复的元素需要累加，c1、c2、c3
			res[i+j+1] += str1[i] * str2[j]
		}
	}
	fmt.Println(res)
	// 需要倒序，从高位开始，除数*10留给低位，即从右向左开始计算
	for i := len(res) - 1; i > 0; i-- {
		if res[i] > 10 {
			res[i-1] += res[i] / 10
			res[i] %= 10
		}
	}
	return res
}
