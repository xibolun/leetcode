package top100

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	var source = []struct {
		a   string
		rse bool
	}{
		// {"()", true},
		// {"()[]{}", true},
		// {"(]", false},
		{"([)]", false},
	}

	for _, v := range source {
		// r := isValid(v.a)
		r := isValid2(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 栈
func isValid(s string) bool {
	mp := map[rune]rune{')': '(', ']': '[', '}': '{'}
	st := []rune{}
	for _, item := range s {
		// 没有数据信息，说明是左括号
		if mp[item] == 0 {
			st = append(st, item)
		} else {
			// 条件是判断有没有匹配的括号
			if len(st) == 0 || st[len(st)-1] != mp[item] {
				return false
			}
			// 右括号的处理
			st = st[:len(st)-1] //出栈
		}
	}
	return len(st) == 0
}

// 下面这种解法不满足：左括号必须以正确的顺序闭合。
var strArray = []byte{'(', ')', '{', '}', '[', ']'}

func isValid2(s string) bool {
	queue := map[rune]int{}
	for _, s := range strArray {
		queue[rune(s)] = 0
	}

	for _, item := range s {
		v := queue[item]
		queue[item] = v + 1
	}
	return queue['('] == queue[')'] && queue['['] == queue[']'] && queue['{'] == queue['}']
}
