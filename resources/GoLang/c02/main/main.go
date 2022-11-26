package main

import "fmt"

func main() {
	//变量的定义
	//静态的语言的变量和动态语言的变量定义差异很大
	//1. 最基础的变量定义
	//var i int // int i
	//i = 20
	//fmt.Println(i)
	//定义并初始化
	//var i int = 10
	//fmt.Println(i)
	//2. 根据值自行判断变量类型（类型推断）
	//var i = 100 //并没有设置类型
	//fmt.Println(i)

	// 3. go语言既然一门新语言 省略var
	//i := 100
	//fmt.Println(i)

	//var a int = 10
	//var b = 10
	//c := 10
	//fmt.Println(a, b , c)

	//多变量定义
	//var a, b, c int
	//a, b, c = 10, 20, 30
	//fmt.Println(a, b, c)

	//var a, b, c int = 10, 20 ,30
	//fmt.Println(a, b, c)

	//// 集合类型
	//var (
	//	a int
	//	name string
	//)

	//var i int =10
	//i = 20
	//fmt.Println(i)

	//匿名变量， 变量一旦被定义 不使用
	//常量 -
	//const PI = 3.1415926
	//r := 2.0
	//运行过程中， 代码写的不好 一不小心在某个地方将PI给改掉了

	//枚举
	//const (
	//	Unknown = 0
	//	Female = 1
	//	Male = 2
	//)
	//
	////常量组如不指定类型和初始化值，该类型和值和上一行的类型一致
	//const (
	//	x int = 16
	//	y
	//	s = "abc"
	//	z
	//)
	//fmt.Println(x,y,s)
	//1. 常量的数据类型值可以是布尔，数字 和字符串
	//2. 不曾使用的常量，在编译的时候是不会报错
	//const常量的iota， 常量计数器 枚举
	const (
		Book = iota //计算器
		Cloth //行
		Phone
		DeskTop
	)
	//0,1,2 本身不重要， 这三个值不一样
	//iota 该常量的值等于上一个常量的表达式
	fmt.Println(Book, Cloth, Phone, DeskTop)
	//iota你真的懂了吗？
	// 1. iota只能在常量组中是使用
	// 2. 不同的const定义块互相不干扰
	const (
		Unknown = iota
		Female
		Male
	)
	fmt.Println(Unknown, Female, Male)

	//3. 没有表达式的常量定义复用上一行的表达式
	//4. 从第一行开始，iota从0逐行加一
	const (
		a = iota //iota  = 0
		b = 10 //iota = 1 每一行iota加一
		c // c=10 , iota = 2
		d,e = iota, iota //iota=3
		f = iota //iota=4
		//iota代表的是这里的行数
	)
	
	fmt.Println(a, b, c, d, e, f)
	//变量的作用域
}
