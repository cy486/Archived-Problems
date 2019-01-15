package main

import (
	"fmt"
	"math"
)

func main() {
	i := 2
	num := 1
	for num <= 10001 {
		if Prime(i) {
			num++
		}
		i++
	}
	fmt.Println(i - 1)
}
func Prime(num int) bool {
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
