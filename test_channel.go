// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	/*
// 		Go语言中的通道（channel）是一种特殊的类型。在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据。goroutine 间通过通道就可以通信。

// 		通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
// 		声明通道类型
// 		通道本身需要一个类型进行修饰，就像切片类型需要标识元素类型。通道的元素类型就是在其内部传输的数据类型，声明如下：
// 		var 通道变量 chan 通道类型

// 		通道类型：通道内的数据类型。
// 		通道变量：保存通道的变量。

// 		chan 类型的空值是 nil，声明后需要配合 make 后才能使用。

// 		创建通道
// 		通道是引用类型，需要使用 make 进行创建，格式如下：
// 		通道实例 := make(chan 数据类型)

// 		数据类型：通道内传输的元素类型。
// 		通道实例：通过make创建的通道句柄。

// 	*/

// 	// var p_c chan int
// 	// p_c = make(chan int)
// 	// fmt.Printf("p_c: %v\n", p_c)
// 	// ch1 := make(chan int

// 	/*
// 	   使用通道发送数据
// 	   通道创建后，就可以使用通道进行发送和接收操作。
// 	   1) 通道发送数据的格式
// 	   通道的发送使用特殊的操作符<-，将数据通过通道发送的格式为：
// 	   通道变量 <- 值

// 	   通道变量：通过make创建好的通道实例。
// 	   值：可以是变量、常量、表达式或者函数返回值等。值的类型必须与ch通道的元素类型一致。

// 	   把数据往通道中发送时，如果接收方一直都没有接收，那么发送操作将持续阻塞。Go 程序运行时能智能地发现一些永远无法发送成功的语句并做出提示
// 	*/
// 	// ch := make(chan interface{})
// 	// ch <- 0
// 	// ch <- "hellop"
// 	// ch1 := make(chan int)
// 	// ch1 <- 20
// 	// ch1 <- "fafds"
// 	/*
// 		通道接收同样使用<-操作符，通道接收有如下特性：
// 		① 通道的收发操作在不同的两个 goroutine 间进行。

// 		由于通道的数据在没有接收方处理时，数据发送方会持续阻塞，因此通道的接收必定在另外一个 goroutine 中进行。

// 		② 接收将持续阻塞直到发送方发送数据。

// 		如果接收方接收时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止。

// 		③ 每次接收一个元素。
// 		通道一次只能接收一个数据元素。
// 	*/
// 	ch := make(chan string)
// 	var rCh <-chan string = ch
// 	var wCh chan<- string = ch
// 	go f(rCh)
// 	var input string
// 	for {
// 		fmt.Scanln(&input)
// 		if input == "quite" {
// 			close(ch)
// 		} else {
// 			wCh <- input
// 		}

// 	}

// }

// func f(ch <-chan string) {
// 	i := 1
// 	for {
// 		// s, ok := <-ch
// 		// if ok {
// 		// 	fmt.Printf("in f s: %v\n", s)
// 		// } else {
// 		// 	fmt.Printf("...")
// 		// }
// 		select {
// 		case s := <-ch:
// 			fmt.Printf("s: %v\n", s)
// 			i = 1
// 		default:
// 			fmt.Println("..", i)
// 			i++
// 			time.Sleep(time.Second * 2)
// 		}

// 	}
// }
