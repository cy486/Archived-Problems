package main

import (
	"fmt"
	"strconv"
)

/*
Digit factorialsProblem 34
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
Find the sum of all numbers which are equal to the sum of the factorial of their digits.
Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/
/**
1.把数字转化成字符串，然后拆开
2.判断是否相等
*/
//求一个数的阶乘的函数
func Factorial(n int) int {
	var result int
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
func Breakjc(i int) int {
	result := 0
	str := strconv.Itoa(i)
	for _, n := range str {
		result += Factorial(int(n - 48))
	}
	return result
}
func main() {
	var i, j int
	i = 3
	//求所有组成数字的阶乘和的函数
	for {
		j = Breakjc(i)
		if i == j {
			fmt.Println(i)
		}
		i++
	}

}
