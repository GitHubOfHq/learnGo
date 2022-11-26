package main

import (
	"fmt"
	"reflect"
)

// 结构体能基本上达到类的一个效果 多态
// 1.标签示例:将属性按需序列化
//type Info struct { //能表述的信息是有限的
//	Name   string `json:"name"`          //name是映射成mysql中char类型还是varchar类型还是text类型， 即使能够说明 但是额外的信息 max_length
//	Age    int    `json:"age,omitempty"` //age字段在序列化的时候 序列化为age omitempty表示忽略0值
//	Gender string `json:"-"`             //不序列化
//}

// 2.自定义标签
type Info struct {
	Name   string `orm:"name,max_length=17,min_length=5"` //自定义标签如orm,条件以逗号隔开,key-value形式用标签限制属性
	Age    int    `orm:"age,min=18,max=70"`
	Gender string `orm:"gender,required"`
}

//type Info struct { //能表述的信息是有限的
//	Name string `orm:"name, max_length=17, min_length=5"`
//	Age int `orm:"age, min=18, max=70"`
//	Gender string `orm:"gender, required"`
//}

//反射包

func main() {
	//结构体标签
	/*
		结构体的字段除了名字和类型外，还可以有一个可选的标签（tag）：
		它是一个附属于字段的字符串，可以是文档或其他的重要标记。
		比如在我们解析json或生成json文件时，常用到encoding/json包，
		它提供一些默认标签，例如：omitempty标签可以在序列化的时候忽略0值或者空值。
		而-标签的作用是不进行序列化，其效果和和直接将结构体中的字段写成小写的效果一样。
	*/
	info := Info{
		Name:   "bobby",
		Gender: "男",
	}
	//1.
	//re, _ := json.Marshal(info) //序列化为字符串
	//fmt.Println(string(re))     //{"name":"bobby"}

	//2.
	//通过反射包去识别这些tag 简单的体验了一下反射的威力 spring 底层都是反射
	t := reflect.TypeOf(info)
	fmt.Println("Type:", t.Name()) //Type: Info
	fmt.Println("Kind:", t.Kind()) //Kind: struct

	for i := 0; i < t.NumField(); i++ { //numfield 能获取有多少个最字段
		field := t.Field(i)         //获取结构体第i个字段
		tag := field.Tag.Get("orm") //获取tagname为orm的标签
		fmt.Printf("%d. %v (%v), tag: `%v`\n", i+1, field.Name, field.Type.Name(), tag)
	}
	//1. Name (string), tag:
	// `name,max_length=17,m
	//in_length=5`
	//2. Age (int), tag: `ag
	//e,min=18,max=70`
	//3. Gender (string), ta
	//g: `gender,required`

	//具体的应用绝大部分情况之下我们是不需要使用到反射的 实际开发的项目中会用到的
	//接口 java 实际上在go语言中接口这个概念的地位和java中接口的地位是不一样， go语言的接口实际上就是python中的协议 - 鸭子类型
}
