package main

import "fmt"

func swap(a *int, b *int) {
	//用于交换a和b
	c := *a
	*a = *b
	*b = c
}
func main() {
	a, b := 10, 20
	//1.指针定义于使用
	var ip *int //这个变量里面就只能保存地址类型这种值
	ip = &a
	*ip = 30
	fmt.Printf("ip所指向的内存空间地址是：%p, 内存中的值是: %d\n", ip, *ip)
	swap(&a, &b)
	fmt.Println(a, b)

	//2.数组相关
	//指针还可以指向数组 指向数组的指针 数组是值类型
	//数组指针
	arr := [3]int{1, 2, 3}
	var iparr *[3]int = &arr
	fmt.Println(*iparr) //[1 2 3]
	//指针数组
	//var ptrs [3]*int //创建能够存放三个int指针的数组

	//go语言没有屏蔽指针，但是go语言在指针上做了大量的限制，安全性高很多，相比较 c和c++灵活性就降低了
	//指针变量中涉及到两个符号 & 和 *
}
