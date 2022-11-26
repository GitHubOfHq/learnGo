package main

import (
	"fmt"
	"strings"
)

func main() {
	//1.字符串长度
	//返回的是字节长度
	//unicode 字符集 go
	//go采用utf8编码 3个字节表示一个中文 一个字节表示英文 \n表示一个字符
	//var name string = "bobby:魔酷网"
	//fmt.Println(len(name))
	////如何求真正长度?类型转换
	////如将每个字符扩大程int32类型的rune 统计rune个数
	//name_arr := []rune(name)
	//fmt.Println(len(name_arr))

	//转义符
	var name string = "bobby:\"魔酷网\""
	//go中`可以将字符串内容原样输出
	date := `2020\06\18`
	fmt.Println(date)
	fmt.Println(len(name))

	//2. 是否包含某个子串
	fmt.Println(strings.Contains(name, "魔酷网"))
	//位置
	fmt.Println(strings.Index(name, "魔"))
	//次数
	fmt.Println(strings.Count(name, "b"))
	//前缀和后缀
	fmt.Println(strings.HasPrefix(name, "bo"))
	fmt.Println(strings.HasSuffix(name, "b"))
	//大小写
	fmt.Println(strings.ToUpper("bobby"))
	fmt.Println(strings.ToLower("Bobby"))

	//3.比较 去空格或指定字符串 分割 替换 连接
	fmt.Println(strings.Compare("a", "b"))       //-1
	fmt.Println(strings.TrimSpace("  bob by  ")) //bob by
	fmt.Println(strings.TrimLeft("bobbyb", "b")) //obbyb
	fmt.Println(strings.Trim("bobbyb", "b"))     //obby
	// 分割
	arrs := strings.Split("bobby,imooc", ",")
	fmt.Println(arrs) //[bobby imooc]
	// 合并 将字符串数组连接起来
	fmt.Println(strings.Join(arrs, "-")) //bobby-imooc
	// 替换
	//参数: 待替换字符串 旧子串 新子串 替换个数n
	//从左往右替换n个
	fmt.Println(strings.Replace("bobby: 18 电话:188888888", "18", "19", 1))

}
