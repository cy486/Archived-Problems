package main

import (
	"fmt"
	"os"
)

/*通过添加自然数生成三角数的序列。所以第7 个三角形数字是1 + 2 + 3 + 4 + 5 + 6 + 7 = 28.前十个术语是：

1,3,6,10,15,21,28,36,45,55 ......

让我们列出前七个三角形数字的因子：

1：1
3：1,3
6：1,2,3,6
10：1,2,5,10
15：1,3,5,15
21：1,3,7,21
28：1,2， 4,7,14,28
我们可以看到28是第一个超过五个除数的三角形数。

拥有超过500个除数的第一个三角形数的值是多少？*/

//分解因式的程序（返回所有短除法剩下的值）
/*
1.从2开始，寻找输入数字的最小的整除数，加入到切片arr中
2.如果除到最后的水不为一，最后的在数字也加到切片中
3.返回切片
*/
func find2(a int) []int {
	var arr []int
	for i := 2; i*i <= a; i++ {
		for a%i == 0 {
			a = a / i
			arr = append(arr, i)
		}
	}
	if a != 1 {
		arr = append(arr, a)
		arr = append(arr, a)
	}
	/*fmt.Println(arr)*/
	return arr
}

//求一个数组不同的数字的个数（返回个数切片）
/*
1./从1开始遍历，arr[0]赋值给m，如果arr[i]的值和m不相等，把用切片记录开始变化的下标
2.a切片为数组值开始变化的下标
3.把a[0]的值赋值给b[0]
4.b切片储存a切片的两个相邻数字的差值
5.根据约数定理，返回约数个数
*/
func WordCount(arr []int) int {
	var a []int
	var b []int
	m := arr[0]
	count := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != m {
			/*fmt.Println("i", i)*/
			a = append(a, i)
			m = arr[i]
		}
	}
	b = append(b, a[0])
	for i := 1; i < len(a); i++ {
		b = append(b, a[i]-a[i-1])
	}
	if arr[len(arr)-1] != a[len(a)-1] {
		b = append(b, 1)
	}
	for _, i := range b {
		count *= i + 1
	}
	return count
}

/*
分三步进行，根据约数定理
1.进行所有的三角数寻找
2.找到所有的短除数
3.计算因数个数
*/
/*
约数定理：
n可以分解质因数：n=p1^a1×p2^a2×p3^a3*…*pk^ak,
由约数定义可知p1^a1的约数有:p1^0, p1^1, p1^2......p1^a1 ，共（a1+1）个;同理p2^a2的约数有（a2+1）个......pk^ak的约数有（ak+1）个。
故根据乘法原理：n的约数的个数就是(a1+1)(a2+1)(a3+1)…(ak+1)。
*/
func main() {
	b := 28
	a := 7
	c := 0
	/*var arr []int*/
	for {
		/*fmt.Println(b)*/
		a++
		c = WordCount(find2(b))
		/*fmt.Println(c)*/
		if c > 500 {
			fmt.Println(b)
			os.Exit(0)
		}
		b += a
	}
}
