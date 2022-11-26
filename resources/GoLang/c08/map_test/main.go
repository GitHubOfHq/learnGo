package main

import "fmt"

func main() {
	//go中的map ->python中的dict
	//go语言中的map的key和value类型申明的就要指明
	//1map定义
	//1. 直接定义
	m1 := map[string]string{
		"m1": "v1",
		"m2": "v2",
	}
	fmt.Printf("%v\n", m1)
	//2. make函数创建
	m2 := make(map[string]string)
	m2["m2"] = "v2"
	fmt.Printf("%v\n", m2)
	//3. 定义一个空的map
	m3 := map[string]string{}
	fmt.Printf("%v\n", m3)

	//map中的key 不是所有的类型都支持，该类型需要支持 == 或者 != 操作
	//切片类型不支持放入key
	//数组类型支持放入key.

	//2 map基本操作
	m := map[string]string{
		"a": "va",
		"b": "vb",
	}
	//2.1增加 修改 :没有则添加 有则替换
	m["c"] = "vc"
	m["b"] = "vb1"
	fmt.Printf("%v\n", m) //map[a:va b:vb1 c:vc]

	//2.2查询 可以有第二个返回值 代表是否取值成功
	v, ok := m["d"] //key不存在打印缺省value
	if ok {
		fmt.Println("找到了", v)
	} else {
		fmt.Println("没找到")
	}
	fmt.Println(v, ok) //"" false

	//2.3删除
	delete(m, "a") //删除成功
	delete(m, "e") //删除失败 但不抛异常

	//2.4便利
	for k, v := range m {
		fmt.Println(k, v)
		//c vc
		//b vb1
	}
	//go语言中也有一个list 就是数据结构中提到的链表
}
