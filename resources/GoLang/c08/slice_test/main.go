package main

import "fmt"

func replace(mySlice []string) {
	mySlice[0] = "bobby"
}

func main() {
	//什么是切片
	//数组有一个很大的问题：大小确定，不能修改 -> 切片 - 动态数组
	// 1切片定义
	//1.1法一: 无长度数组法 var courses []string //定义了一个切片
	//var courses = []string{"django", "scrapy", "tornado"}
	//fmt.Printf("%T", courses)

	//1.2法二: make法
	//courses := make([]string, 5) //参数:类型,长度
	//fmt.Println(len(courses))//res:5
	//切片不是没有长度限制，为什么使用make初始化的需要我们传递一个长度， 那我传递了长度之后是不是意味着就像数组一样长度不能变了呢？
	//后面小结讲原理
	//slice对标的就是python中list

	//1.3法三: 从数组截取法,通过数组变成一个切片
	var courses = [5]string{"a", "b", "c", "e", "f"} //定义一个数组
	subCourse := courses[1:4]                        //等式右边返回的是切片类型;将数组类型转为切片类型
	fmt.Printf("%T", subCourse)                      //res:[]string
	fmt.Println(subCourse)                           //res: 取左闭右开1-4个元素
	//数组的传递是值传递 切片是引用传递
	subCourse[0] = "bobby" //切片是地址引用 改切片 数组内容也改变
	fmt.Println(courses)   //courses第一个数据变为 bobby
	replace(subCourse)
	fmt.Println(subCourse) //bobby c e

	//1.4法四: new
	//subCourse2 := new([]int)
	//fmt.Println(subCourse2) //带* res:[]  不带*res:&[]

	//2slice操作
	//slice是动态数组，所以说我们需要动态添加值
	//2.1修改值
	//subCourse[1] = "bobby"
	//fmt.Println(subCourse[1])

	//2.2截取slice获得扔是slice
	subCourse2 := subCourse[1:3]
	fmt.Printf("%T, %v\n", subCourse2, subCourse2)

	//2.3append 可以向切片追加元素
	subCourse2 = append(subCourse2, "imooc", "imooc2", "imooc3")
	fmt.Println(subCourse2)
	//也可以追加切片
	appendedCourse := []string{"追加切片", "追加切片", "追加切片"}
	subCourse2 = append(subCourse2, appendedCourse...)
	fmt.Println(subCourse2)

	//2.4 copy
	//拷贝的时候 目标对象长度需要设置好
	//subCourse3 := []string{}//无长度数组法后面拷贝不了
	subCourse3 := make([]string, 2) //长度是几 就拷贝几个
	copy(subCourse3, subCourse2)
	fmt.Println(subCourse3) //打印subCourse2 前两个
	//既然是切片，那为什么这个时候有来提出长度的要求(后面讲原理)

	//2.5 删除元素
	deleteCourse := [5]string{"a", "b", "c", "e", "f"}
	courseSlice := deleteCourse[:]
	courseSlice = append(courseSlice[:1], courseSlice[2:]...) //删除b元素
	fmt.Println(courseSlice)
	//如何判断某个元素是否在切片中 for循环查看

	//3 python和go语言的slice区别
	//1.go的slice更像是python的list， go语言的底层是基于数组实现的 python的list底层也是基于数组实现的
	//2.slice进行的操作都会影响原来的数组， slice更像是一个指针 本身不存值

	//slice的原理，因为很多底层的知识相对来说很多时候并不难而是需要花费比较多的时间去慢慢理解
	//3.1 第一个现象 不会去扩展a的空间
	a := make([]int, 0)
	b := []int{1, 2, 3}
	fmt.Println(copy(a, b))
	//解答 从b往a拷值 不会扩展a的空间 a的空间实际上为0
	//2 第二个现象 修改切片影响到原数组
	c := b[:]
	//c[0] = 8
	//fmt.Println(b)
	//fmt.Println(c)
	//解答 bd底层指向同一块内存
	//3.3 第三个现象 append函数没有影响到原来数组
	c = append(c, 9)
	fmt.Println(b) //[8 2 3]     注释2.代码后 [1 2 3]
	fmt.Println(c) //[8 2 3 9]	 注释2.代码后 [1 2 3 9]
	c[0] = 8
	fmt.Println(b) //注释2.代码后res:[1 2 3]
	fmt.Println(c) //注释2.代码后res:[8 2 3 9]
	//为什么append函数之后再调用c[0]=8不会影响原来数组 append后切片指向新的地址
	//解答 这就是因为产生了扩容机制，扩容机制一旦产生 这个时候切片就会指向新的内存地址

	//3.4 第四个现象
	fmt.Println(len(c)) //res4
	fmt.Println(cap(c)) //res6 cap 指的是容量 长度和容量

	//4 切片底层原理
	//切片底层是使用数组实现的，既要使用数组 又要满足动态的功能 怎么实现？
	//假设有一个值 实际上申请数组的时候可能是两个，如果后续要增加数据那么就直接添加到数据的结尾，这个时候我不要额外重新申请
	//切片有不同的初始化方式
	//4.1. 使用make方法初始化 len和cap是多少. 不会有多余的预留空间
	d := make([]int, 5, 6) //第三个参数是cap容量 不写就和len一样为5
	fmt.Printf("len=%d, cap=%d\n", len(d), cap(d))
	//4.2. 通过数组取切片
	data := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Slice := data[2:4]                                                             //指向data位置2
	newSlice := data[3:6]                                                          //指向data位置3
	fmt.Printf("Slice len=%d, Slice cap=%d\n", len(Slice), cap(Slice))             // 2 8
	fmt.Printf("newSlice len=%d, newSlice cap=%d\n", len(newSlice), cap(newSlice)) //3 7

	//4.3.直接定义
	data2 := []int{1, 2, 3}
	fmt.Printf("data2 len=%d, data2 cap=%d\n", len(data2), cap(data2)) //3 3

	//5 扩容机制
	oldSlice := make([]int, 0)
	oldSlice = append(oldSlice, 1)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))
	oldSlice = append(oldSlice, 2)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))
	oldSlice = append(oldSlice, 3)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))
	oldSlice = append(oldSlice, 4)
	oldSlice = append(oldSlice, 5)
	fmt.Printf("len=%d, cap=%d\n", len(oldSlice), cap(oldSlice))
	//res
	//len 1 2 3 5
	//cap 1 2 4 8

	/*
		Go 中切片扩容的策略是这样的：
		new>=2old res=new
		new<2old  且old<1024 res=2old
		new<2old  且old>1024 res=1.25old
		(文字解释:加上添加数据后容量和旧容量比较)
		首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）
		否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap）
		否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的 1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
		//如果小于1024 扩容的速度是2倍 如果大于了1024 扩容的速度就是1.25
		如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）

	*/

	//切片来总结
	//1. 底层是数组，如果切片是基于数组产生的 会有一个问题就是会影响原来的数组。
	//2. 切片的扩容机制
	//3. 切片的传递是引用传递
	//当make遇到了append容易出现的坑
	s1 := make([]int, 5)
	s1 = append(s1, 6)
	fmt.Println(s1) //[0 0 0 0 0 6]
	//对于很多初学者来说我们期望的是只有一个数字就是6
	//为什么不是只有一个6呢? 因为make创建slice时 指明了len为5即有效数据为5个 虽然是int初始值也没事
}
