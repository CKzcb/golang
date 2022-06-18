// package main

// import "fmt"

// func s(arr []int, l_s int, r_s int, t int) {
// 	if l_s > r_s {
// 		fmt.Println("没查到")
// 		return
// 	}
// 	m_s := (l_s + r_s) / 2

// 	if arr[m_s] > t {
// 		s(arr, l_s, m_s, t)
// 	} else if arr[m_s] < t {
// 		s(arr, m_s, r_s, t)
// 	} else {
// 		fmt.Println("找到了，下标是：", m_s)
// 		return
// 	}
// }

// func bS(arr []int) {
// 	// 冒泡排序
// 	for i := 0; i < len(arr)-1; i++ {
// 		for j := 0; j < len(arr)-i-1; j++ {
// 			if arr[j+1] < arr[j] {
// 				temp := arr[j]
// 				arr[j] = arr[j+1]
// 				arr[j+1] = temp
// 			}
// 		}
// 	}
// }

// func main() {
// 	// 死循环
// 	// a := 20
// 	// for {
// 	// 	time.Sleep(1)
// 	// 	fmt.Printf("a: %v\n", a)
// 	// 	a++
// 	// }
// 	// for a := 2; a < 10; a++ {
// 	// 	fmt.Printf("a=%v\n", a)
// 	// }

// 	// a := 1
// 	// for a < 10 {
// 	// 	fmt.Printf("a: %v\n", a)
// 	// 	a++
// 	// }
// 	// for i := 1; i < 10; i++ {
// 	// 	for j := 1; j <= i; j++ {
// 	// 		fmt.Printf("%d+%d=%d", j, i, i*j)
// 	// 		if j <= i-1 {
// 	// 			fmt.Printf(" ")
// 	// 		}
// 	// 	}
// 	// 	fmt.Printf("\n")

// 	// }

// 	// for range是值拷贝
// 	// var a [3]*int
// 	// c := 20
// 	// d := 3
// 	// e := 30
// 	// a[1] = &c
// 	// fmt.Printf("c: %v\n", a[1])
// 	// a[2] = &d
// 	// a[0] = &e
// 	// f := 1
// 	// for _, v := range a {
// 	// 	v = &f
// 	// 	fmt.Printf("v: %v\n", v)
// 	// }
// 	// fmt.Printf("a[1]: %v\n", a[1])

// 	// switch 支持多个条件，支持条件表达式, 使用fallthrough可以执行下面的语句
// 	// a := 2
// 	// switch {
// 	// case 1 == a:
// 	// 	fmt.Printf("1")
// 	// case 2 == a:
// 	// 	fmt.Printf("2")
// 	// 	fallthrough
// 	// case a < 30:
// 	// 	fmt.Printf("<30")
// 	// default:
// 	// 	fmt.Printf("default")
// 	// }

// 	// goto 语句跳转

// 	// 	a := 20
// 	// 	b := 30
// 	// 	c := a + b
// 	// 	fmt.Println("c = ", c)
// 	// 	if c > 10 {
// 	// 		goto l1
// 	// 	}
// 	// 	fmt.Println("no c>10")

// 	// l1:
// 	// 	fmt.Println("in c 10")

// 	// continue也可以加标签
// 	// l1:
// 	// 	for i := 1; i < 5; i++ {
// 	// 		for j := 0; j < 3; j++ {
// 	// 			switch j {
// 	// 			case 1:
// 	// 				fmt.Printf("%v %v\n", i, j)
// 	// 				continue l1
// 	// 			}
// 	// 			fmt.Printf("j: %v\n", j)
// 	// 		}
// 	// 	}

// 	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	// s(arr, 0, len(arr), 5)
// 	arr := []int{5, 2, 6, 18, 4, 9, 1}
// 	for _, v := range arr {
// 		fmt.Print(v, ",")
// 	}
// 	fmt.Println()
// 	bS(arr)
// 	for _, v := range arr {
// 		fmt.Print(v, ",")
// 	}
// 	fmt.Println()

// }
