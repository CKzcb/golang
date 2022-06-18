// package main

// import "fmt"

// func main() {
// 	/*
// 		接口的方法与实现接口的类型方法格式一致
// 		在类型中添加与接口签名一致的方法就可以实现该方法

// 	*/

// 	p := new(People)
// 	p.name = "zz"
// 	var i Say
// 	i = p
// 	i.SayHello()
// }

// type Say interface {
// 	SayHello()
// 	// WriteName()
// }

// type Eat interface {
// 	EatFood(interface{})
// }

// type People struct {
// 	name string
// }

// func (p *People) SayHello() {
// 	fmt.Printf("p.name: %v\n", p.name+" ~ hello ~")
// }
