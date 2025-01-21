package string

import "strconv"

// tag: 字符串，二进制求和
func addBinary(a string, b string) string {
	carry := 0
	// 二进制加法，进制为2
	la, lb := len(a), len(b)
	n := max(la, lb)
	ans := ""
	for i := 0; i < n; i++ {
		// carry承担着每一次循环的sum
		if i < len(a) {
			carry = carry + int(a[la-1-i]) - '0'
		}
		if i < len(b) {
			carry = carry + int(b[lb-1-i]) - '0'
		}

		// 两者求和后算结果和进制
		ans = strconv.Itoa(carry%2) + ans
		carry = carry / 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
