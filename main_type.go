/*

boolean: true或false   如：var isOpen bool = true

整型数值：int、float32、float64，支持复数

字符串：Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。

派生类型:
	包括：
	(a) 指针类型（Pointer）
	(b) 数组类型
	(c) 结构化类型(struct)
	(d) Channel 类型
	(e) 函数类型
	(f) 切片类型
	(g) 接口类型（interface）
	(h) Map 类型

*/

/*

变量定义：变量名由字母、数字、下划线组成，其中首个字符不能为数字。

声明变量的一般形式是使用 var 关键字：
				var identifier type
可以一次声明多个变量，但是属于同一个类型
				var identifier1, identifier2 type

只声明不初始化的话默认值为：
	- 数值类型（包括complex64/128）为 0

	- 布尔类型为 false

	- 字符串为 ""（空字符串）

第二种，根据值自行判定变量类型,自动推到类型

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

:= 只能在函数体内使用

局部变量不能只声明使用，全局变量可以只声明不使用

*/

// func main() {
// 	// int
// 	var a int = 20
// 	fmt.Printf("a: %v\n", a)
// }

/*
	值类型和引用类型
	支持指针，和c语言几乎一样
*/
// package main

// import "fmt"

// var a int

// func main() {

// 	/*
// 		常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。
// 		const identifier [type] = value
// 		可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
// 			- 显式类型定义： const b string = "abc"
// 			- 隐式类型定义： const b = "abc"
// 		常量可以用作枚举
// 			const (
// 				Unknown = 0
// 				Female = 1
// 				Male = 2
// 			)

// 		iota
// 			iota，特殊常量，可以认为是一个可以被编译器修改的常量。

// 		iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
// 	*/
// 	/* const pi = 3.1415
// 	fmt.Printf("pi: %v, %T\n", pi, pi) */
// 	const (
// 		a = iota //0
// 		b        //1
// 		c        //2
// 		d = "ha" //独立值，iota += 1
// 		e        //"ha"   iota += 1
// 		f = 100  //iota +=1
// 		g        //100  iota +=1
// 		h = iota //7,恢复计数
// 		i        //8
// 	)
// 	fmt.Println(a, b, c, d, e, f, g, h, i)
// }
