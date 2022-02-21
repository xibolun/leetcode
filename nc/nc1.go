package nc

import "strconv"

// https://www.nowcoder.com/practice/11ae12e8c6fe48f883cad618c2e81475?tpId=196&tqId=37176&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196%26page%3D1%26difficulty%3D3&difficulty=3&judgeStatus=undefined&tags=&title=

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算两个数之和
 * @param s string字符串 表示第一个整数
 * @param t string字符串 表示第二个整数
 * @return string字符串
 */
func solve(s string, t string) string {
	v := ""
	for i, j, carry := len(s)-1, len(t)-1, 0; i >= 0 || j >= 0; {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(s[i] - '0')
		}

		if j >= 0 {
			n2 = int(t[j] - '0')
		}

		tmp := n1 + n2 + carry
		carry = tmp / 10

		v = strconv.Itoa(tmp%10) + v
		i--
		j--
	}
	return v
}
