package main

import "fmt"

func main() {
	sum := 0
	sum2 := 0
	for i := 1; i <= 100; i++ {
		sum += i
		sum2 += i * i
	}
	fmt.Println(sum*sum - sum2)
}
