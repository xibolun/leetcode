package string

import (
	"reflect"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	var source = []struct {
		a   int
		rse string
	}{
		// {1, "A"},
		// {28, "AB"},
		{701, "ZY"},
	}

	for _, v := range source {
		r := convertToTitle(v.a)
		if !reflect.DeepEqual(r, v.rse) {
			t.Errorf("test fail, result: %v, should be :%v", r, v.rse)
		}
	}
}

// tag: 171相似,进制转换
func convertToTitle(columnNumber int) string {
	ans := []byte{}

	for columnNumber > 0 {
		m := columnNumber % 26
		ans = append([]byte{byte(m + 64)})
		columnNumber = (columnNumber - m)
	}

	// var convert func(ans []byte, carry, m int) []byte
	// convert = func(ans []byte, carry, m int) []byte {
	// 	v := byte((m%26 + carry + 64))
	// 	ans = append([]byte{v}, ans...)
	// 	m = m / 26
	// 	if m < 1 {
	// 		return ans
	// 	}
	// 	carry += 1
	// 	return convert(ans, carry, m)
	// }
	// ans = convert(ans, 0, columnNumber)

	return string(ans)
}
