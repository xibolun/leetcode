package main

import (
	"fmt"
	"math"
)

func myPow(x float64, n int) float64 {

	y := math.Abs(float64(n))

	v := math.Pow(x, y)

	if n < 0 {
		return 1 / v
	}

	return v
}

func main() {
	fmt.Println(myPow(2, -3))
}
