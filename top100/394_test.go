package top100

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestDecodeString(t *testing.T) {
	var source = []struct {
		a   string
		rse string
	}{
		// {"3[a]2[bc]", "aaabcbc"},
		// {"100[leetcode]", "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode"},
		{"3[a2[c]]", "accaccacc"},
	}

	for _, v := range source {
		r := decodeString(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 栈,双栈
func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	num := 0
	result := ""
	for _, char := range s {
		// 若为数字
		if char >= '0' && char <= '9' {
			n, _ := strconv.Atoi(string(char))
			num = num*10 + n
		} else if char == '[' {
			// 说明num已经计算结束，可以入栈了
			numStack = append(numStack, num)
			// 重置num和result，继续下一次
			num = 0
			strStack = append(strStack, result)
			result = ""
		} else if char == ']' {
			// 说明encode_str都已经写入完成，可以开始计算
			count := numStack[len(numStack)-1]    // 需要重复的次数
			numStack = numStack[:len(numStack)-1] // 移除count

			// 移除strNum当中的最后一个数组
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			// 3[a2[c]];
			// str=a
			// result = 3[acc]
			result = string(str) + strings.Repeat(result, count)
		} else {
			// 什么也不是，即为encode_string
			result += string(char)
		}

	}

	return result
}

// FIXME
func decodeString1(s string) string {
	stack := make([]byte, 0)
	encode_str := make([]byte, 0)
	num := make([]byte, 0)
	ans := ""

	for j := len(s) - 1; j >= 0; j-- {
		// 将]先都移除掉
		if s[j] != ']' {
			// stack = append([]byte{s[j]}, stack...)
			stack = append(stack, s[j])
			continue
		}
	}

	for _, item := range stack {
		// 现在遍历，只看[,  [左边为数字，右边为encode_string
		if item > '0' && item < '9' {
			num = append([]byte{item}, num...)
		} else {
			if item == '[' {
				continue
			} else {
				if len(num) <= 0 {
					encode_str = append([]byte{item}, encode_str...) // 因为是倒序，所以值需要放在左边
				} else {
					n, _ := strconv.Atoi(string(num))
					// 若不为0，需要重置此encode_str
					for i := 0; i < n; i++ {
						ans = string(encode_str) + ans
					}
					num = []byte{}
					encode_str = []byte{item}
				}
			}
		}

	}
	if len(num) > 0 && len(encode_str) > 0 {
		n, _ := strconv.Atoi(string(num))
		// 若不为0，需要重置此encode_str
		for i := 0; i < n; i++ {
			ans = string(encode_str) + ans
		}
	}

	return ans
}

// 这种做法不满足{"3[a2[c]]", "accaccacc"},
// 因为不是从最左端开始的
// FIXME
func decodeString2(s string) string {
	l := 0
	r := l + 2
	encode_str := []byte{}
	ans := ""
	for {
		if r > len(s)-1 {
			break
		}
		k, _ := strconv.Atoi(string(s[l]))

		for s[r] != ']' {
			encode_str = append(encode_str, s[r])
			r++
		}

		for i := 0; i < k; i++ {
			ans = ans + string(encode_str)
		}

		l = r + 1
		r = l + 2
		encode_str = []byte{}
	}
	return ans
}
