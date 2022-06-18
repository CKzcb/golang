// package main

// import "fmt"

// func f1(a, b int) (int, int) {
// 	c := a - b
// 	d := a + b
// 	return c, d
// }

// func s1(a, b *int) {
// 	temp := *a
// 	*a = *b
// 	*b = temp
// }

// func s2(a, b int) (c, d int) {
// 	c = 1
// 	d = 2
// 	return
// }

// func main() {
// 	/*
// 		三种方式
// 		1.普通函数
// 		fun name(参数)(返回值){}

// 		2.函数也是一种类型
// 		函数变量

// 		3.匿名函数
// 			func(参数列表)(返回参数列表){
// 				函数体
// 			}

// 		4.闭包=函数+引用环境
// 		闭包中的逻辑可以修改闭包捕获的变量，变量会跟随闭包生命期一直存在，闭包本身就如同变量一样拥有了记忆效应。

// 	*/
// 	// _, f := f1(2, 3)
// 	// fmt.Println(f)
// 	// a := 1
// 	// b := 2

// 	// fmt.Println(a, b)
// 	// s1(&a, &b)
// 	// fmt.Println(a, b)
// 	// fmt.Println(s2(3, 4))
// 	// a := make(map[string]int, 10)
// 	// for k, v := range a {
// 	// 	fmt.Println(k, v)
// 	// }
// 	// tm(a)
// 	// fmt.Println("m ++++++++++++++")
// 	// for k, v := range a {
// 	// 	fmt.Println(k, v)
// 	// }

// 	// var f func(int, int) (int, int)
// 	// f = s2
// 	// fmt.Println(f(1, 2))

// 	// s := "hi~j"
// 	// n := 1
// 	// f := func() {
// 	// 	s = "hi~b"
// 	// 	n++
// 	// }
// 	// fmt.Println(s, n)

// 	// n = n + 20
// 	// f()
// 	// fmt.Println(s, n)
// 	// switch n.(type) {
// 	// case bool:
// 	// 	fmt.Println(1)
// 	// }

// 	// f2(1, 2, 3, 4)
// 	/* defer
// 		   有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出），下面的代码是将一系列的数值打印语句按顺序延迟处理，

// 		   处理业务或逻辑中涉及成对的操作是一件比较烦琐的事情，比如打开和关闭文件、接收请求和回复请求、加锁和解锁等。在这些操作中，最容易忽略的就是在每个函数退出处正确地释放和关闭资源。

// 	defer 语句正好是在函数退出时执行的语句，所以使用 defer 能非常方便地处理资源释放问题。
// 	*/

// 	// defer fmt.Println(1)
// 	// defer fmt.Println(2)
// 	// defer fmt.Println(3)
// 	st := new(Student)
// 	st.say("hi~")
// }

// type People interface {
// 	say(s string)
// 	eat(s string)
// }

// type Student struct {
// }

// func (a Student) say(s string) {
// 	fmt.Println("say: ", s)
// }

// func f2(arg ...int) {
// 	fmt.Println(arg)
// }

// func tm(m map[string]int) {
// 	m["a"] = 20
// }
