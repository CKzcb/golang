/* package main

func main() {

	// 	switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。

	// switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。

	// switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。

	var m int = 2
	// switch m {
	// case 1:
	// 	println("in 1")
	// case 2:
	// 	println("in 2")
	// 	fallthrough
	// case 3:
	// 	println("in 3")
	// default:
	// 	println("in default")
	// }

	m = 10

	// 两种形式，switch可以跟变量也可以直接把条件写到case上

	switch {
	case m == 1:
		println("in 1")
	case m == 2:
		println("in 2")
	case m >= 3:
		println("in 3")
	case m == 4:
		println("in 4")

	}

}
*/