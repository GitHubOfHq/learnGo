package main

import "fmt"

func main()  {
	//for init; condition; post { } 计算1-10的和
	sum := 0
	for i := 1; i<=10; i++ {
		sum += i
	}

	fmt.Println(sum)

	//1. i := 1
	//2. 是否 i<=10
	//3. sum = 1
	//4. i ++  -> i = 2
	//5. 是否 i <= 10
	//6. sum = 3
	//7. i++ ... i=11
	//i := 0
	//for ;i < 10; i ++ {
	//	fmt.Println("bobby")
	//}
	
	//循环字符串
	name := "bobby:慕课网"
	for _, value := range name {
		//为什么是数字 ， ascii
		fmt.Printf("%c\n",  value)
		//fmt.Println(index, value)
	}

	//1. name是一个字符串， 2. 字符串是字符串的数组
	name_arr := []rune(name)
	for i := 0; i<len(name_arr); i++ {
		fmt.Printf("%c\n", name_arr[i])
	}

	//2. 在做字符串遍历的时候要尽量使用range
}
