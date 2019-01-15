package main

import "fmt"

/*如果我们列出10以下的所有自然数是3或5的倍数，我们得到3,5,6和9.这些倍数的总和是23。

求出1000以下3或5的所有倍数的总和。*/
func main() {
	result := 0
	fmt.Println("1000以内3和5的倍数之和")
	for i := 0; i < 1000; i++ {
		if (i%3 == 0 && i%5 != 0) || (i%5 == 0 && i%3 != 0) || (i%3 == 0 || i%5 == 0) {
			result += i
		}
	}
	fmt.Println(result)
}
