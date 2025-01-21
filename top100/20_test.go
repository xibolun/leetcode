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
	stack := []rune{}
	for _, item := range s {
		// 没有数据信息，说明是左括号:([{ 三种当中的任意一个
		if mp[item] == 0 {
			stack = append(stack, item)
		} else {
			// 条件是判断有没有匹配的括号
			// 如果顶端的数据值与item不匹配
			if len(stack) == 0 || stack[len(stack)-1] != mp[item] {
				return false
			}
			// 右括号的处理 )]} 三种当中的任意一种如果出现，把顶端移出，因为顶端数值的就是与其闭合的括号对
			stack = stack[:len(stack)-1] //出栈
		}
	}
	return len(stack) == 0
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
