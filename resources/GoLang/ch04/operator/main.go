package main

import "fmt"

func main() {
	//算数运算符
	/*	a := 12
		b := 22
		fmt.Println(a - b)
		a++
		fmt.Println(a)*/

	//逻辑运算符
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Println("第一行 - 条件为 true\n")
	}
	if a || b {
		fmt.Println("第二行 - 条件为 true\n")
	}
	/*修改a和b的值*/
	a = false
	b = true
	if a && b {
		fmt.Println("第三行为true")
	} else {
		fmt.Println("第三行为true")
	}
}
