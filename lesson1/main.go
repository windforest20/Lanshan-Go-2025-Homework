package main

import "fmt"

const pi = 3.14

func factorial(n int) int {
	j := 1
	var i int
	for i = 1; i <= n; i++ {
		j = j * i
	}
	return j
}
func main() {
	//LV0
	fmt.Println("Hello fenglin!")
	//LV1
	var r float64
	r = 5
	area := pi * r * r
	fmt.Println(area)
	//LV2
	var sum int
	for i := 1; i <= 1000; i++ {
		sum += i
	}
	fmt.Println(sum)
	//LV3
	a := factorial(10)
	fmt.Println(a)
	//LVX

	var j, num, cout int
	cout = 0
	for {
		fmt.Print("请输入一个整数（输入0结束）：") //1.不断让用户输入整数，输入0表示结束
		fmt.Scanln(&j)
		num += j
		if j == 0 {
			break
		}
		cout++
	} //2.使用for循环统计输入数字的总和与个数
	average := func(x, y int) int {
		return x / y
	} //3.定义一个函数用于计算平均值
	result := average(num, cout)
	fmt.Println(result) //4.根据平均值输出结果
}
