package string

import (
	"fmt"
	"strings"
)

// 数字转中文： 123456 -> 一十二万三千四百五十六

var chUnit = []string{"拾", "佰", "仟", "万", "拾万", "佰万", "仟万"}
var chNum = map[int]string{0: "零", 1: "壹", 2: "贰", 3: "叁", 4: "肆", 5: "伍", 6: "陆", 7: "柒", 8: "捌", 9: "玖"}

func Convert(source int) string {
	target := make([]int, 0)
	target = splitV(source, target)
	reverseInt(target)

	result := make([]string, 0)
	for i, v := range target {
		vCh, _ := chNum[v]
		// 个位数
		if i >= len(target)-1 {
			result = append(result, fmt.Sprintf("%s", vCh))
			continue
		}
		if v != 0 {
			result = append(result, fmt.Sprintf("%s%s", vCh, chUnit[len(target)-i-2]))
			continue
		}
		// 零的处理；若上一个也是零，则不用再添加
		if result[len(result)-1] == vCh {
			continue
		}
		result = append(result, fmt.Sprintf("%s", vCh))
	}

	return strings.Join(lastZero(result), "")
}

// 对末尾的0进行处理
func lastZero(result []string) []string {
	if result[len(result)-1] != "零" {
		return result
	}
	if result[len(result)-1] == "零" {
		return lastZero(result[:len(result)-1])
	}
	return result
}

func splitV(source int, target []int) []int {
	if source == 0 || source < 10 {
		return target
	}
	for source != 0 {
		mod := source % 10
		target = append(target, mod)
		source = source / 10
		splitV(source, target)
	}
	return target
}

func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
