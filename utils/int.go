package utils

import "strconv"

func IntJoin(arr []int, sep string) string {
	if len(arr) == 0 {
		return ""
	}
	var ret string
	for i, v := range arr {
		ret += strconv.Itoa(v)
		if i != len(arr)-1 {
			ret += sep
		}
	}
	return ret
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FloatMax(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
