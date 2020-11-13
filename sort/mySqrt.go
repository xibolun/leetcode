package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 2, x/2

	for {
		if left > right {
			return right
		}
		v := left + (right-left)/2
		num := v * v
		if num < x {
			left = v + 1
			continue
		}
		if num > x {
			right = v - 1
			continue
		}
		return v
	}
}

func main() {
	fmt.Println(mySqrt(3))
}
