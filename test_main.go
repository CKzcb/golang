package main

import (
	"fmt"

	t "./tpackage"
)

func main() {

	s := t.NewStudent("aa", 20)
	fmt.Printf("s.GetName(): %v\n", s.GetName())
}
