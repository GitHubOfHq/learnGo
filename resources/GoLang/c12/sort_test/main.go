package main

import (
	"fmt"
	"sort"
)

type Course struct {
	Name  string
	Price int
	Url   string
}

type Courses []Course

func (c Courses) Len() int {
	return len(c)
}
func (c Courses) Less(i, j int) bool {
	return c[i].Price < c[j].Price
}
func (c Courses) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	//通过sort来排序
	//很多 冒泡 插入 快速 归并 桶排序
	//你的排序算法是否能够应付各种类型的排序
	//查看源码如下:必须要实现下面三个方法 下面算法适用通过比较排序
	//Len() int 			//长度
	//Less(i, j int) bool   //对象i是不是比j小
	//Swap(i, j int)		//
	//想对slice类型排序 而slice类本身没有实现上述三个方法
	//可以通过type重新声明一个类型 对这个类型添加上述方法 见Courses方法
	courses := Courses{
		Course{"张三", 300, ""},
		Course{"李四", 100, ""},
		Course{"王五", 200, ""},
	}
	sort.Sort(courses) //协议 你的目的不是告诉别人具体的类型(不管数据类型或对象),重要的是你的类型必须提供具体的方法即实现三个方法
	for _, v := range courses {
		fmt.Println(v)
	}
}
