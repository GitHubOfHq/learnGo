package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//动态分配内存:make和new函数
	//1.new函数用法
	//var p *int //申明了一个变量p p为nil
	//*p = 10//nil地址的值设为10错误
	//默认值 int byte rune float bool string 这些类型声明都有默认值
	//指针， 切片 map， 接口这些默认值是nil 理解为none
	var p *int = new(int) //go的编译器就知道先申请一个内存空间，这里的内存中的值全部设置为0
	*p = 10

	//2.更加常用的是make函数，new函数返回的是这个值的地址 指针 make函数返回的是指定类型的实例
	//make函数一般用于切片 map
	var info map[string]string = make(map[string]string)
	info["c"] = "bobby"

	//3.nil的一些细节
	var info2 map[string]string
	if info2 == nil {
		fmt.Println("map的默认值是 nil")
	}

	var slice []string
	if slice == nil {
		fmt.Println("slice的默认值是 nil")
	}

	var err error
	if err == nil {
		fmt.Println("error的默认值是 nil")
	}

	//python中的None和go语言中的nil类型不一样，None是全局唯一的
	//go语言中 nil是唯一可以用来表示部分类型的零值的标识符， nil因类型不同而大小不同
	fmt.Println(unsafe.Sizeof(slice), unsafe.Sizeof(info2)) //24 8

}
