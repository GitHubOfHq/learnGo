package main

import (
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		panic("被除数不能为0")
	}
	return a / b, nil
}

//错误和异常 错误就是可以预先知道出错的可能, 这个时候我们把这种情况叫做错误,不可预知的问题叫做异常

func main() {
	//错误就是能遇到可能出现的情况，这些情况会导致你的代码出问题， 参数检查， 数据库访问不了
	//strconv.Itoa(data) //go 认为这个Itoa的函数不可能出错， 没有必要返回error 内部代码出错这个时候应该抛出异常 panic
	//_, err := strconv.Atoi("abcd") //Atoi这个函数认为我的函数内部会出现一些预知错误情况
	//if err != nil {
	//	//错误
	//}

	//异常 go语言中如何抛出异常和如何捕捉异常
	//go语言认为错误就要自己处理， 就个人而言，go语言的这种想法是正确的。但是实际的使用中确实人有点烦人
	a := 12
	b := 0
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("异常被捕获")
		}
		fmt.Println("bobby")
	}()
	fmt.Println(div(a, b))

	//panic的坑:再http_server文件

}
