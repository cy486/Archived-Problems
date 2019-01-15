package main

import (
	"fmt"
	"math"
)

func main() {
	sum := 0
	for i := 1; i < 2000000; i++ {
		if Prime1(i) {
			sum += i
		}
	}
	fmt.Println(sum - 1)
}
func Prime1(num int) bool {
	if num == 2 || num == 3 {
		return true
	}
	if num%6 != 1 && num%6 != 5 {
		return false
	}
	tmp := int(math.Sqrt(float64(num)))
	for i := 5; i <= tmp; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
