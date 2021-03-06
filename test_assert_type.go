// package main

// import "fmt"

// func main() {
// 	/*
// 	   在Go语言中类型断言的语法格式如下：
// 	   value, ok := x.(T)

// 	   其中，x 表示一个接口的类型，T 表示一个具体的类型（也可为接口类型）。

// 	   该断言表达式会返回 x 的值（也就是 value）和一个布尔值（也就是 ok），可根据该布尔值判断 x 是否为 T 类型：
// 	   如果 T 是具体某个类型，类型断言会检查 x 的动态类型是否等于具体类型 T。如果检查成功，类型断言返回的结果是 x 的动态值，其类型是 T。
// 	   如果 T 是接口类型，类型断言会检查 x 的动态类型是否满足 T。如果检查成功，x 的动态值不会被提取，返回值是一个类型为 T 的接口值。
// 	   无论 T 是什么类型，如果 x 是 nil 接口值，类型断言都会失败。
// 	*/
// 	// var x interface{}
// 	// x = 20
// 	// v, ok := x.(int)
// 	// fmt.Println(v, ok)
// 	// f(20)
// 	// f("sdaf")

// 	/*
// 	   空接口类型类似于 C# 或 Java 语言中的 Object、C语言中的 void*、C++ 中的 std::any。在泛型和模板出现前，空接口是一种非常灵活的数据抽象保存和使用的方法。

// 	   空接口的内部实现保存了对象的类型和指针。使用空接口保存一个数据的过程会比直接用数据对应类型的变量保存稍慢。因此在开发中，应在需要的地方使用空接口，而不是在所有地方使用空接口。
// 	*/

// 	var a interface{} = 20
// 	v, k := a.(int)
// 	fmt.Printf("a: %v %v\n", v, k)

// }

// func f(a interface{}) {
// 	fmt.Printf("a: %T, %v\n", a, a)
// }
