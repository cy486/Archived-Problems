package main

import "fmt"

/*Fibonacci序列中的每个新术语都是通过添加前两个术语生成的。从1和2开始，前10个术语将是：

1,2,3,5,8,13,21,34,55,89 ......

通过考虑Fibonacci序列中的值不超过四百万的项，找到偶数项的总和。*/

/*
1.不断的进行斐波那契数的累加
2.验证是不是偶数，是偶数便加到和上面
3.加到4000000项，程序结束

*/
func main() {
	i := 1
	j := 2
	sum := 0
	for {
		//1.不断的进行斐波那契数的累加
		temp := i + j
		i = j
		j = temp
		//2.验证是不是偶数，是偶数便加到和上面
		if j%2 == 0 {
			sum += j
		}
		//3.加到4000000项，程序结束
		if j > 4000000 {
			break
		}
	}
	//因为2也是，但是我的程序没有把2加进去
	fmt.Println(sum + 2)
}

//斐波那契数列
//斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
