package main

import (
	"fmt"
)

// 接口是一个协议- 现在自我创建一个程序员协议 - 只要你能够 1. 写代码 2. 解决bug 其实就是一组方法的集合

type Programmer interface {
	Coding() string //函数名 返回值(可没有)
	Debug() string
}

//对比结构体
//type name struct{
//属性集合
//}

type Designer interface {
	Design() string
}
type Manger interface {
	Programmer      //既能写代码
	Designer        //又能设计
	Manage() string //又能管理
}

// java的话 java里面一种类型只要继承一个接口 才行 如果你继承了这个接口的话 那么这个接口里面的所有方法你必须要全部实现
type UIDesigner struct {
}

func (d UIDesigner) Design() string {
	fmt.Println("我会ui设计")
	return "我会ui设计"
}

type Pythoner struct {
	UIDesigner //pythoner虽然没有Designer的功能 但可以组合designer的功能 从而实现Manager接口 可以付给Manager
	lib        []string
	kj         []string
	years      int
}

type G struct {
}

// Pythoner 和 G 都实现Coding和Debug方法 就实现了Programma接口
func (p G) Coding() string {
	fmt.Println("go开发者")
	return "go开发者"
}

func (p G) Debug() string {
	fmt.Println("我会go的debug")
	return "我会go的debug"
}

func (p Pythoner) Coding() string {
	fmt.Println("python开发者")
	return "python开发者"
}

func (p Pythoner) Debug() string {
	fmt.Println("我会python的debug")
	return "我会python的debug"
}

func (p Pythoner) Manage() string {
	fmt.Println("不好意思，管理我也懂")
	return "不好意思，管理我也懂"
}

//
//func (p Pythoner) Design() string{
//	fmt.Println("我是一个python开发者，但是我会ui设计")
//	return "我是一个python开发者，但是我会ui设计"
//}
//对于Pythoner这个结构体来说 你实现任何方法都可以 ，但是只有全部实现Coding Debug方法 才是一个Programmer
//1. Pythoner本身自己就是一个类型 那我何必在意我是不是Programmer
//2. 多态的概念对于很多pythoner来说会有点陌生
//3. 在讲解多态之前 我们来对interface做一个说明：在go语言中接口是一种类型(类似int)，是一种抽象类型

//开发中经常会遇到的问题
//开发一个电商网站， 支付环节 使用 微信、支付宝、银行卡 你的系统支持各种类型的支付 每一种支付类型都有统一的接口
// 定一个协议 1. 创建订单 2. 支付 3. 查询支付状态 4. 退款
//支付发起了
//type AliPay struct {
//
//}
//type WeChat struct {
//
//}
//
//type Bank struct {
//
//}
//
//var b Bank
//var a AliPay
//var w WeChat

//多态 什么类型的时候你申明的类型是一种兼容类型， 但是实际赋值的时候是另一种类型
//接口的强制性
//你现在有一个缓存 - 这个地方你一开始使用的缓存是redis 但是后期你考虑到可能使用其他的缓存技术 - 本地 memcache

//这种多态特性 其实在python中不需要多态 python是动态语言

//go语言中并不支持继承

//如果后期接入一种新的支付 或者取消已有的支付

// 接口作为函数参数
func HandlePy(p Programmer) {

}

//3.

type MyError struct {
}

func (e MyError) Error() string {
	return "错误"
}

func main() {

	//新的语言出来了, 接口帮我们完成了go语言的多态
	//var pro Programmer = Pythoner{}

	var pros []Programmer
	pros = append(pros, Pythoner{})
	pros = append(pros, G{})

	//接口虽然是一种类型 但是和其他类型不太一样 接口是一种抽象类型只有声明没有实现 struct是具象
	p := Pythoner{}
	fmt.Printf("%T\n", p)
	var pro Programmer = Pythoner{}
	fmt.Printf("%T\n", pro) //Pythoner类型
	var pro2 Programmer = G{}
	fmt.Printf("%T", pro2) //G类型
	//如果大家对象面向对象理解的话 java 里面的抽象类型

	//2.下一部分
	//1. go struct组合 组合一起实现了所有的接口的方法也是可以的
	//2. 接口本身也支持组合

	//pythoner必须全部实现Manager方法 自己方法不能实现Manager部分方法 可以通过组合其他结构体实现Manager剩余方法
	var m Manger = Pythoner{} //这里Pythoner不实现Manager全部方法 会报错
	m.Design()
	//python语言本身设计上是采用了完全的基于鸭子类型 - 协议 影响了python语法的 for len()
	//struct组合完成了接口 1. 接口支持组合 - 继承 2. 结构体组合实现了所有的接口方法也没有问题

	//3.下一部分
	//go语言本身也推荐鸭子类型  error
	//法一 errors包实现了error接口 errors.New实现了接口的函数 通过下面函数可以生成一个字符串error
	//var err error = errors.New(fmt.Sprintf(""))
	//法二 通过fmt.Errorf也可以实现 更常用的方法
	s := "文件不存在"
	var err error = fmt.Errorf("错误:%s", s)
	fmt.Println(err)
}
