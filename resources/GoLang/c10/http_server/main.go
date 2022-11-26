package main

import (
	"fmt"
	"net/http"
)

func f1() {
	////使用组合一
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Println("捕获到了")
	//	}
	//}()
	//panic("出错了1")
	//time.Sleep(10 * time.Second)
	//使用组合二
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("捕获到了2")
			}
		}()

		panic("出错了")
	}()
	//总结:即panic和recover必须再同一协程
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		//go func() {
		//	panic("error")
		//}()
		//上述三行将panic放到协程里 抛出错误会导致主线程挂掉 同时导致其他协程也挂掉
		//panic会引起主线程的挂掉， 同时会导致其他的协程都挂掉

		//panic("error") //不放协程里就没事
		w.Write([]byte("hello world!"))
	})

	http.ListenAndServe("127.0.0.1:8080", nil)

	//上述func(w http.Res....){}函数已经注册一个recover,这时候直接写panic,func函数可以捕获到
	//而如果将panic放到go func(){}内部 go func(){}的recover函数捕获到panic异常,从而导致外面的func(w http...0{}
	//函数无法捕获异常
	//为什么直接在, 在父协程中无法捕获子协程中出现的异常
	//操作redis的，有人觉得这段代码可以放到协程中去运行。有一个非常大的隐患了

	//调用上面f1函数
	f1()

}
