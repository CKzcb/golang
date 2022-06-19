package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "陌生人", "姓名")
	conn, err := net.Dial("tcp", "47.233.27.128:5454")
	if err != nil {
		log.Fatal("err conn ..")
	}
	log.Println("start connecting ... ")
	log.Println("你好啊~ ", name)
	go clientRead(conn)
	var input string
	for {
		fmt.Scanln(&input)
		if input == "q" {
			break
		}
		// conn.Write([]byte(input))
		fmt.Fprintf(conn, input)
	}
}

func clientRead(conn net.Conn) {
	read := bufio.NewReader(conn)
	for {
		var p [512]byte
		n, err := read.Read(p[:])
		if err != nil {
			log.Println("read err ...")
		} else {
			if n > 0 {
				fmt.Println(time.Now().Format("%T-%m-%d %H:%M:%S"), ":", string(p[:n]))
			}
		}

	}
}
