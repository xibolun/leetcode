package nc

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolve(t *testing.T) {
	var source = []struct {
		a   string
		b   string
		rse string
	}{
		{"1", "99", "100"},
		{"114514", "", "114514"},
	}

	for _, v := range source {
		r := solve(v.a, v.b)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}

}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算两个数之和
 * @param s string字符串 表示第一个整数
 * @param t string字符串 表示第二个整数
 * @return string字符串
 */
func solve(s string, t string) string {
	v := ""
	carry := 0
	for i, j := len(s)-1, len(t)-1; i >= 0 || j >= 0; {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(s[i] - '0')
		}

		if j >= 0 {
			n2 = int(t[j] - '0')
		}

		tmp := n1 + n2 + carry
		carry = tmp / 10 //进位

		v = strconv.Itoa(tmp%10) + v
		i--
		j--
	}
	if carry > 0 {
		v = strconv.Itoa(carry) + v
	}
	return v
}
