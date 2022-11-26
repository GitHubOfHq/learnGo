package main

import (
	"fmt"
	"strconv"
)

func main() {
	//printf 格式化
	name := "bobby"
	age := 18
	fmt.Println("name:"+name+", age:"+strconv.Itoa(age))
	fmt.Printf("name:%T, age: %T\n", name, age)
	fmt.Printf("name:%s, age:%x,\n ", name, age)
	desc := fmt.Sprintf("name:%s, age:%x,\n ", name, age)
	fmt.Println(desc)

	data := 65
	fmt.Printf("%q\n", data)
	fmt.Printf("%e", 65.1)
	
	//输入
	var n string
	var a int
	//fmt.Println("请输入你的姓名和年龄:")
	//fmt.Scanln(&n, &a)
	//fmt.Println(n, a)

	//通过scanf输入
	fmt.Println("请输入你的姓名和年龄:")
	//输入的内容必须符合指定的格式
	fmt.Scanf("请输入你的姓名和年龄:%s %d", &n, &a)
	fmt.Println(n, a)

}
