// // go 循环
// /*
// Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号。

// 和 C 语言的 for 一样：
// for init; condition; post { }

// 	- init： 一般为赋值表达式，给控制变量赋初值；
// 	- condition： 关系表达式或逻辑表达式，循环控制条件；
// 	- post： 一般为赋值表达式，给控制变量增量或减量。

// 和 C 的 while 一样：
// for condition { }

// 和 C 的 for(;;) 一样：
// for { }
// */

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// var m = 3
// 	// if m == 1 {
// 	// 	println("in 1")
// 	// } else if m == 2 {
// 	// 	println("in 2")
// 	// } else {
// 	// 	println("in 3")
// 	// }

// 	// for i := 0; i < 10; {
// 	// 	i++
// 	// 	fmt.Printf("i: %v\n", i)
// 	// }

// 	// for a, v := range "nihao" {
// 	// 	fmt.Printf("a: %v, v: %v\n", a, v)
// 	// }

// 	// go 语言break可以跳出当前循环，带标签可以跳出指定循环

// 	// for i := 0; i < 10; i++ {
// 	// 	if i < 4 {
// 	// 		continue
// 	// 	}
// 	// 	fmt.Printf("i: %v\n", i)
// 	// 	if i > 5 {
// 	// 		break
// 	// 	}
// 	// }

// 	for i := 0; i < 3; i++ {
// 		fmt.Printf("i: %v\n", i)
// 	label1:
// 		for j := 0; j < 3; j++ {
// 			fmt.Printf("j: %v\n", j)
// 			for z := 0; z < 4; z++ {

// 				if z > 0 {
// 					continue label1
// 				}
// 				fmt.Printf("z: %v\n", z)
// 			}
// 			fmt.Printf("in j\n")
// 		}
// 		fmt.Printf("in i\n")
// 	}

// }
