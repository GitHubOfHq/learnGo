package main

import (
	"fmt"
)

func main() {
	score := 90
	/*
		>90  A
		80-89 B
		70-79 C
		60-69 D
		<60 E
	*/
	grade := "A"

	//1.go 默认有break 若想穿透这一层用fallthrough
	/*switch score {
	case 90:
		grade = "A"
		fallthrough
	case 80:
		grade = "B"
	default:
		grade = "D"

	}
	fmt.Println(grade)*/

	//2.go switch可以做区间
	switch {
	case score >= 90:
		grade = "A"
		fmt.Println("优")
	case score >= 80 && score <= 89:
		grade = "B"
	case score >= 70 && score <= 79:
		grade = "C"
	case score >= 70 && score <= 79:
		grade = "D"
	default:
		grade = "E"
	}

	fmt.Println(grade)
	//3.一个case有多个条件即"或",中间用逗号分隔。一分支多条件值
	var a = "daddy"
	switch a {
	case "mum", "daddy":
		fmt.Println("family")
	}
	//做接口类型判断
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
