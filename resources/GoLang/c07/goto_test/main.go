//package main
//
//import "fmt"
//
//func main() {
//	/* 定义局部变量 */
//	var a int = 10
//
//	/* 循环 */
//	LOOP: for a < 20 {
//		if a == 15 {
//			/* 跳过迭代 */
//			a = a + 1
//			goto LOOP
//		}
//		fmt.Printf("a的值为 : %d\n", a)
//		a++
//	}
//}

//package main
//import "fmt"
//func main() {
//	var breakAgain bool
//	// 外循环
//	for x := 0; x < 10; x++ {
//		// 内循环
//		for y := 0; y < 10; y++ {
//			// 满足某个条件时, 退出循环
//			if y == 2 {
//				// 设置退出标记
//				breakAgain = true
//				// 退出本次循环
//				break
//			}
//		}
//		// 根据标记, 还需要退出一次循环
//		if breakAgain {
//			break
//		}
//	}
//	fmt.Println("done")
//}

package main
import "fmt"
func main() {
	//goto能不能则不用，goto过于， label过多，整个程序到后期维护就会麻烦
	//最容易理解的代码逐行的执行，哪怕多一个函数的调用对于我们都是理解上的负担
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}
	// 手动返回, 避免执行进入标签
	return
	// 标签
breakHere:
	fmt.Println("done")
}