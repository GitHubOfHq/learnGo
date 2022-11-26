package main

import "fmt"

//go语言的继承
// 例子:Course继承Teacher内容
// 一句话总结: 可以用"对象.属性/方法"方式访问父类属性和方法 若有重名再用类名标注

type Teacher struct {
	Name  string
	Age   int
	Title string
}

func (t Teacher) teacherInfo() {
	fmt.Printf("姓名:%s, 年龄:%d, 职称:%s", t.Name, t.Age, t.Title)
}

type Course struct {
	Teacher //如果讲师的信息比较多怎么办 将另一个结构体的变量放进来
	Name    string
	Price   int
	Url     string
}

// 匿名嵌套:在子类中只声明父类而不带对象(即只有Teacher)
func (c Course) courseInfo() {
	fmt.Printf("课程名:%s, 价格:%d, 讲师信息：%s %d %s\n", c.Name, c.Price, c.Teacher.Name, c.Age, c.Title)
}

func main() {
	//go语言的继承 组合
	t := Teacher{
		Name:  "bobby",
		Age:   18,
		Title: "程序员",
	}
	c := Course{
		Teacher: t,
		Price:   100,
		Url:     "",
		Name:    "django",
	}
	c.courseInfo()
	c.teacherInfo()

}
