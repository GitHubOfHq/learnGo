package main

import "fmt"

type EmptyInterface interface {
}

type Course struct {
	name  string
	price int
	url   string
}

// 2.打印任何类型
func print(i interface{}) {
	fmt.Printf("%v\n", i)
}

// 4.1
func (c Course) printInfo() string {
	return "课程信息"
}

type Printer interface {
	printInfo() string
}

// 4.2
// 需求1:根据类型做不同打印
func printByType(x interface{}) {
	//判断x是不是int类型 相当于java的instance of 如果x是int 则ok是true反之false
	//牵扯到go的另一个默认问题 返回值不是err 用ok代替
	//if v, ok := x.(int);ok {
	//	fmt.Printf("%d,整数\n", v)
	//}

	//这样会有大量的if语句 用switch优化
	switch v := x.(type) {
	case string:
		fmt.Printf("%s(字符串)\n", v)
	case int:
		fmt.Printf("%s(整数)\n", v)
	}
}

// 需求2:往本地或阿里云上传
type AliOss struct {
}
type LocalFile struct {
}

func store(x interface{}) {
	switch y := x.(type) {
	case AliOss:
		//此处要做一些特殊处理,比如设置阿里云权限问题
		fmt.Println(y)
	case LocalFile:
		//检查路径的权限
	}
}

func main() {
	//1.空接口 不需要名字只把上面EmptyInterface 后面的interface{}拿下来就好了
	var i interface{}
	//空姐口可以类似于我们java的object最顶层的多台 即任何类型都可以放到空接口里面
	i = Course{}
	fmt.Println(i)
	i = "bobby" //也行
	fmt.Println(i)
	i = []string{"django", "scrapy"} //也行
	fmt.Println(i)
	//空接口的作用一: 接受任何类型变量

	//2.空接口作为参数传递 见print函数

	//3.空接口作为map值 接受任何类型
	var teacherInfo = make(map[string]interface{})
	teacherInfo["name"] = "bobby"
	teacherInfo["age"] = "18"
	teacherInfo["weight"] = 75.2
	print(teacherInfo)

	//4 类型断言
	//4.1 接口一个坑:接口不能接受地址传递 当实现函数是地址时 不能实例化对象 下面错误
	//c Printer := Course{} //错
	//如果想把对象赋值给接口类型 必须是值传递
	//即func (c *Course) printInfo() string { 把*去掉然后按下述赋值
	var c Printer = Course{}
	c.printInfo()
	// 接口有一个默认规范 接口的名称一般以 er结尾

	//4.2 类型断言 即java中instance of 见printByType函数
}
