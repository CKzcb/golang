package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5454")
	if err != nil {
		log.Fatal("err conn ..")
	}
	log.Println("start connecting ... ")
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
				fmt.Println("other say:", string(p[:n]))
			}
		}

	}
}
