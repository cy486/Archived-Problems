package main

import (
	"fmt"
	"os"
)

/*13195的主要因素是5,7,13和29。

600851475143的最大主要因素是什么？*/
/*
这个程序是某天晚上写的，由于大脑异常活跃，导致忘记当初思路
想了半天。
程序求最大的因数，所以我从10开始，用所给的数除10以内的数，如果可以整除就记录这个数，并继续遍历100以内的
*/
func main() {

	a1 := 600851475143
	a2 := 10
	saun(a2, a1)
}

func saun(a2, a1 int) {
	temp := 0
	m := 1
	for i := 2; i < a2; i++ {
		for j := 2; j < i; j++ {
			temp = j
			if i%j == 0 {
				break
			}
		}
		if temp == i-1 {
			if a1%i == 0 {
				if a1/i == 1 {
					fmt.Println(a1)
					os.Exit(0)
				}
				m = a1 / i
				a1 = m
				m = 1
			}
		}
	}
	saun(a2*10, a1)
}
