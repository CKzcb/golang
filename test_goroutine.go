// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func main() {

// 	/*

// 				Go 程序中使用 go 关键字为一个函数创建一个 goroutine。一个函数可以被创建多个 goroutine，一个 goroutine 必定对应一个函数。
// 				1) 格式
// 				为一个普通函数创建 goroutine 的写法如下：
// 				go 函数名( 参数列表 )

// 				函数名：要调用的函数名。
// 				参数列表：调用函数需要传入的参数。

// 				使用 go 关键字创建 goroutine 时，被调用函数的返回值会被忽略。

// 				所有 goroutine 在 main() 函数结束时会一同结束。

// 		goroutine 虽然类似于线程概念，但是从调度性能上没有线程细致，而细致程度取决于 Go 程序的 goroutine 调度器的实现和运行环境。

// 		终止 goroutine 的最好方法就是自然返回 goroutine 对应的函数。虽然可以用 golang.org/x/net/context 包进行 goroutine 生命期深度控制，但这种方法仍然处于内部试验阶段，并不是官方推荐的特性。
// 	*/

// 	// go f()
// 	// var inputName string
// 	// fmt.Scanln(&inputName)
// 	// fmt.Printf("inputName: %v\n", inputName)

// 	// fmt.Scanln(&inputName)

// 	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())

// }

// func f() {
// 	var i = 0
// 	for {
// 		i++
// 		fmt.Printf("i: %v\n", i)
// 		time.Sleep(time.Second * 5)
// 	}
// }
