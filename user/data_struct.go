package data_struct

import "fmt"

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

*/




func m(){

	var name string
	fmt.Printf(name)

}


