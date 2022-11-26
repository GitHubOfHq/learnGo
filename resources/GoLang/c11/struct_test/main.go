package main

import (
	"fmt"
	"unsafe"
)

// 结构体名和属性名首字母尽量大写,能让其他包访问
type Course struct {
	Name  string
	Price int
	Url   string
}

// 结构体封装方法: 在函数名前面标注所属结构体 通过该函数访问结构体变量 达到面向对象封装特性
// 调用方式:c10.printCourseInfo()  在下面
// 不带*值传递
func (c Course) printCourseInfo() {
	fmt.Printf("课程名:%s, 课程价格: %d, 课程的地址:%s", c.Name, c.Price, c.Url)
}

// 带*引用传递 这里的c相当于this
func (c *Course) setPrice(price int) {
	c.Price = price
}

//1. 结构体的方法只能和结构体在同一个包中
//2. 内置的int类型不能加方法

func main() {
	//go语言不支持面向对象
	//面向对象的三个基本特征： 1. 封装 2. 继承 3. 多态 4. 方法重载 4. 抽象基类
	//定义struct go语言没有class这个概念 所以说对于很多人来说会少理解很多面向对象抽象的概念

	//1. 实例化- kv形式
	var c Course = Course{
		Name:  "django",
		Price: 100,
		Url:   "https://www.imooc.com",
	}
	//访问
	fmt.Println(c.Name, c.Price, c.Url)
	//大小写在go语言中的重要性 可见性
	//(重点)一个包中的结构体或者结构体中的变量如果首字母是小写 那么对于另一个包不可见
	//结构体定义的 名称 以及属性首字母大写很重要(大写能被其他包访问,如果只有自己包访问才用小写)

	//2. 第二种实例化方式 - 顺序形式
	c2 := Course{"scrapy", 110, "https://www.imooc.com"}
	fmt.Println(c2.Name, c2.Price, c2.Url)

	//3. 如果一个指向结构体的指针, 通过结构体指针获取对象的值， 让很多人莫名其妙
	c3 := &Course{"tornado", 100, "https://www.imooc.com"}
	//fmt.Printf("%T", c3) //*main.Course c3是一个指向结构体的指针
	//另一个根本的原因 - go语言的指针是受限的 c3是一个地址 再go中不好修改
	fmt.Println(c3.Name, c3.Price, c3.Url) //(重要)这里其实是go语言的一个语法糖 go语言内部会将c3.Name转换成 (*c3).Name

	//4. 零值 如果不给结构体赋值， go语言会默认给每个字段采用默认值
	c4 := Course{}
	fmt.Println(c4.Price) //默认值

	//5. 多种方式零值初始结构体
	var c5 Course = Course{}
	var c6 Course
	var c7 *Course = &Course{}
	//var c8 *Course  //打印c8.Price错误
	//为什么c6和c8表现出来的结果不一样 指针如果只申明不赋值 默认值是nil c6不是指针 是结构体的类型
	//对于指针 slice map用new方法
	fmt.Println("零值初始化")
	fmt.Println(c5.Price) //0
	fmt.Println(c6.Price) //0
	fmt.Println(c7.Price) //0

	//6. 结构体是值类型
	c8 := Course{"scrapy", 110, "https://www.imooc.com"}
	c9 := c8
	fmt.Println(c8) //
	fmt.Println(c9) //两内容一致
	c8.Price = 200
	fmt.Println(c8) //两内容不一致
	fmt.Println(c9)

	//go语言中struct无处不在 string slice map都是结构体
	//7. 结构体的大小 占用内存的大小 可以使用sizeof来查看对象占用的类型
	fmt.Println(unsafe.Sizeof(1)) //8
	//go语言string的本质 其实string是一个结构体
	//type string struct{
	//	Data uintptr //指针占8个长度
	//	Len int //长度64位系统占8个长度
	//}
	fmt.Println(unsafe.Sizeof("")) //字符串都是16字节
	fmt.Println(unsafe.Sizeof(c8)) //16+8+16=40

	//8. slice的大小始终是指向数组结构体大小即8+8+8
	type slice struct {
		array unsafe.Pointer // 底层数组的地址 8字节
		len   int            // 长度
		cap   int            // 容量
	}

	s1 := []string{"django", "tornado", "scrapy", "celery", "snaic", "flask"}
	fmt.Println("切片占用的内存:", unsafe.Sizeof(s1)) //24

	//map大小:8
	m1 := map[string]string{
		"bobby1": "django",
		"bobby2": "tornado",
		"bobby3": "scrapy",
		"bobby4": "celery",
	}
	fmt.Println(unsafe.Sizeof(m1)) //8

	//结构体方法， 达到了封装数据和封装方法的效果
	c10 := Course{"scrapy", 110, "https://www.imooc.com"}
	//c10.printCourseInfo() //本质是Course.printCourseInfo(c10)

	//Course.setPrice(c10, 200)
	(&c10).setPrice(200) //想要修改c10的price?结构体是值传递不行, 则函数接受者(方法名前面结构体加*)加*用地址接收,
	fmt.Println(c10.Price)

	//结构体的接收者有两种形式 1. 值传递 2. 指针传递 如果你想改结构体的值 如果结构体的数据很大
	//go语言不支持继承 但是有办法能达到同样的效果 组合

}
