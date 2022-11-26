package main

//
//import "fmt"
//
//// finally是在return之前调用
////go语言并没有提供try except finally
////解决1. 大量的嵌套try finally 2. 打开和关系逻辑离的太远
//
//// //1. defer 第一个特性: 压栈先入后出
//func main() { //下述代码结果是 test1->test2->defer test3->2->1
//	fmt.Println("test1")
//	//defer之后只能是函数调用 不能是表达式 比如 a++
//	defer fmt.Println("defer test1")
//	defer fmt.Println("defer test2")
//	defer fmt.Println("defer test3")
//	/*
//		defer 语句是go的一种用于注册延迟调用的机制，将函数调用压栈延迟处理 它可以让当前函数执行完毕之后执行
//		对于python的with语句来说，
//	*/
//	//此处有大量的逻辑需要读取文件
//	fmt.Println("test2")
//	//1. 如果有多个defer会出现什么情况 多个defer是按照先入后出的顺序执行
//}
//
//// 2.defer拷贝机制-浅拷贝 将函数浅拷贝一份压栈 后面改不影响栈中内容
//func main() { //下面结果test3->test1
//	//defer语句执行时的拷贝机制
//	test := func() {
//		fmt.Println("test1")
//	}
//	defer test()
//	test = func() { //在调用test前尝试改掉
//		fmt.Println("test2")
//	}
//	fmt.Println("test3")
//}
//
//func main() { //结果是10 因为defer的逻辑时候带上参数 后面修改不影响压栈后的内容
//	//defer语句执行时的拷贝机制
//	x := 10
//	defer func(a *int) {
//		fmt.Println(*a)
//	}(x)
//	x++
//}
//
//func main() { //结果11  压栈时内部没有提供x 只能指向外部的x
//	//defer语句执行时的拷贝机制
//	x := 10
//	//此处的defer函数并没有参数，函数内部使用的值是全局的值
//	defer func(a int) {
//		fmt.Println(x)
//	}(x) //这个小括号代表使用匿名函数 传递的参数是x 没有()则只是函数声明
//	x++
//}
//
//// 带return 先执行defer出栈再return
//func f1() int { //结果10
//	x := 10
//	defer func() {
//		x++
//	}()
//	tmp := x //x是int类型 值传递
//	return tmp
//}
//
//func f2() *int { //结果11
//	a := 10
//	b := &a
//	defer func() {
//		*b++
//	}()
//	temp_data := b
//	return temp_data
//}
//func main() {
//	fmt.Println(f1()) //是不是就意味着 defer中影响不到外部的值呢
//	fmt.Println(*f2())
//	//defer本质上是注册了一个延迟函数，defer函数的执行顺序已经确定
//	//defer 没有嵌套 defer的机制是要取代try except finally
//	//https://www.cnblogs.com/zhangboyu/p/7911190.html
//	//https://studygolang.com/articles/24044?fr=sidebar
//}
