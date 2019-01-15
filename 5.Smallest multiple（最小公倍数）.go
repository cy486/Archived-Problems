package main

import "fmt"

func main() {
	a := 1
	for i := 2; i <= 19; i++ {
		a = Gb(a, i+1)
	}
	fmt.Println(a)
}
func Gb(a, b int) (result int) {
	m := a
	n := b
	for b != 0 {
		c := a % b
		fmt.Println(c)
		a, b = b, a%b
	}
	return m * n / a
}
