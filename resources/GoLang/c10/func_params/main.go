package main

import "fmt"

// 省略号 ...
// 作用一
func add(params ...int) (sum int) {
	//有不定个int值传递进来
	for _, v := range params {
		sum += v
	}
	params[0] = 9
	return
}

// 2.2小节使用
type sub func(a, b int) int //定义一种函数风格:参数是int 返回值是int sub就是后面一坨

// 参数一 切片 参数二 函数
func filter(score []int, f func(int) bool) []int {
	reSlice := make([]int, 0)
	for _, v := range score {
		if f(v) { //如果v符合函数f 则将v加入slice
			reSlice = append(reSlice, v)
		}
	}
	return reSlice
}

func main() {
	//1 省略号作用
	//通过省略号去动态设置多个参数值
	slice := []int{1, 2, 3, 4, 5}
	//传值和传slice区别，slice是一种类型， 还是引用传递,可以改slice的值
	fmt.Println(add(slice...)) //...第二个作用将slice打散成过个int传进函数
	fmt.Println(slice)

	//省略号的用途 1. 函数参数不定长 2. 将slice打散 3.省略定义数组时长度设置
	arr := [...]int{1, 2, 3}
	fmt.Printf("%T-%v\n", arr, arr) //[3]int-[1 2 3]

	//2 函数一等公民特性
	//2.1函数可赋值给变量
	myFunc := add
	fmt.Printf("%T\n", myFunc) //func(...int) int

	//将匿名函数赋值给变量
	res := func(a, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(res) //3

	//2.2函数做参数
	//定义一种函数风格 参数是两个int 返回值也是int
	var mySub sub = func(a, b int) int {
		return a - b
	}
	fmt.Println(mySub(1, 2))

	//现在才开始讲 将函数作为另一个函数的参数
	//写一个函数用于过滤一部分数据
	score := []int{10, 50, 80, 90, 85}
	//写一个函数过滤掉不合格的成绩
	fmt.Println(filter(score, func(a int) bool {
		if a >= 90 {
			return true
		} else {
			return false
		}
	}))

}
