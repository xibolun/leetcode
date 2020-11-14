package stack

import (
	"strings"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	var source = []struct {
		input  string
		result bool
	}{
		{"((((", false},
		{"()[]", true},
		{"(}", false},
		{"}{", false},
	}
	for _, v := range source {
		result := ValidParentheses(v.input)
		if v.result != result {
			t.Errorf("test fail, input: %s, result: %t, should be :%t", v.input, result, v.result)
		}
	}
}

/**
匹配括号
（） true
()[]{} true
*/
const (
	left  = "[({"
	right = "])}"
)

func ValidParentheses(input string) bool {
	stack := make([]string, 0)
	for i := range input {
		s := string(input[i])
		if strings.Contains(left, s) {
			stack = append(stack, s)
			continue
		}
		// len(stack)>0是为了处理那些以right符号开头的情况
		if s == "]" && len(stack) > 0 && stack[len(stack)-1] == "[" ||
			s == ")" && len(stack) > 0 && stack[len(stack)-1] == "(" ||
			s == "}" && len(stack) > 0 && stack[len(stack)-1] == "{" {
			stack = stack[0 : len(stack)-1]
		}

	}
	return len(stack) == 0
}
