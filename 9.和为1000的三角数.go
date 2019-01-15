package main

import (
	"fmt"
	"os"
)

/*毕达哥拉斯三元组是一组三个自然数，a < b < c，其中，

a 2 + b 2 = c 2
例如，3 2 + 4 2 = 9 + 16 = 25 = 5 2。

恰好存在一个毕达哥拉斯三元组，其中a + b + c = 1000.
找到产品abc。*/
/*
1.三角形两边之和大于第三边，c边从500开始向下遍历
2.分别遍历500以内的a,b边，如果三个数是三角形数而且和为1000，便结束程序
*/
func main() {
	var a, b, c int
	for a = 500; a > 1; a-- {
		for b = 1; b < 500; b++ {
			for c = 1; c < 500; c++ {
				if b*b+c*c == a*a && a+b+c == 1000 {
					fmt.Println(a * b * c)
					os.Exit(0)
				}

			}
		}
	}
}
