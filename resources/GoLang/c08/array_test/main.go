package main

import (
	"fmt"
)

func printArray(toPrint [5]string) { //省略var
	//[]string这是切片类型
	toPrint[0] = "bobby"
	fmt.Println(toPrint)
}

func main() {
	//go语言中的数组和python的list可以对应起来理解，slice和python的list更像
	//1 静态语言中的数组： 1. 大小确定 2. 类型一致
	//1.1数组的申明及创建
	//var courses [10] string
	//var courses = [5]string{"django", "scrapy", "tornado"}
	course := [5]string{"django", "scrapy", "tornado"}

	//数组的另一种创建方式 省略号
	d := [...]int{1, 2, 3, 4, 5, 6} //不指定长度
	fmt.Println(d)
	//指定位置设置值:指定位置3个为100 不设置用默认值[1 0 0 100 0 90]
	e := [...]int{0: 1, 5: 90, 3: 100}
	fmt.Println(e)

	//1.2 修改值， 取值： 删除值, 数组一开始就要指定大小
	//取值
	fmt.Println(course[0])
	//修改值
	course[0] = "django3"
	fmt.Println(course)

	//1.3数组操作
	//数组操作第一种场景： 求长度
	fmt.Println(len(e))
	//数组操作第二种场景： 遍历数组
	for i, value := range course {
		fmt.Println(i, value)
	}
	//使用for语句也可以遍历数组
	sum := 0
	for i := 0; i < len(course); i++ {
		sum += e[i]
	}
	fmt.Println(sum) /*191*/

	//2 数组是值类型(重要区别)
	courseA := [3]string{"django", "scrapy", "tornado"}
	courseB := [...]string{"django1", "scrapy1", "torn" +
		"ado1", "python+go", "asyncio"}
	//courseA和courseB应该是同一种类型， 都是数组类型
	//在go语言中，courseA和courseB都是数组，但是不是同一种类型
	fmt.Printf("%T\n", courseA)
	fmt.Printf("%T\n", courseB)
	/*res:
	[3]string
	[5]string
	*/
	//如果courseA和courseB是一种类型的话 为什么前面要加一个数组， 长度不一样的数组类型是不一样
	//正是基于这些，在go语言中函数传递参数的时候，数组作为参数 实际调用的时候是值传递(即浅拷贝)
	printArray(courseB) //这个地方传递courseA就报错, 因为printArray函数形参个数为5
}
