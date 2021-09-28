package string

import "testing"

func TestConvert(t *testing.T) {
	var source = []struct {
		number int    // 数字 
		numCh  string // 中文
	}{
		{
			number: 126,
			numCh:  "壹佰贰拾陆",
		},
		{
			number: 106,
			numCh:  "壹佰零陆",
		},
		{
			number: 1006,
			numCh:  "壹仟零陆",
		},
		{
			number: 1060,
			numCh:  "壹仟零陆拾",
		},
		{
			number: 10600,
			numCh:  "壹万零陆佰",
		},
		{
			number: 1000600,
			numCh:  "壹佰万零陆佰",
		},
	}

	for _, item := range source {
		if result := Convert(item.number); result != item.numCh {
			t.Errorf("test fail, result: %s, should be :%s", result, item.numCh)
		}
	}
}
