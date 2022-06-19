package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Data struct {
	conn net.Conn
	s    string
}

var ClientConn []net.Conn

func main() {
	s, err := net.Listen("tcp", "127.0.0.1:5454")
	if err != nil {
		log.Fatal("listen err")
	}
	log.Println("start listening ... ")
	var ch chan Data
	ch = make(chan Data)
	go handleConn(ch)
	for {
		conn, err := s.Accept()
		if err != nil {
			log.Print("conn err ...")
		} else {
			log.Println(conn, "connection ... ")
		}
		// go handleConn(conn)
		ClientConn = append(ClientConn, conn)
		go readConn(conn, ch)
	}
}

func handleConn(ch chan Data) {
	fmt.Println("go handleConn ... ")
	for {
		s, ok := <-ch
		log.Println("recv", s, "do send ...")
		if ok {
			doSend(s)
		} else {
			log.Println("recive err ... ")
		}
	}
}

func doSend(d Data) {
	for _, conn := range ClientConn {
		// conn.Write([]byte(s))
		if d.conn == conn {
			continue
		}
		fmt.Fprintf(conn, d.s)
	}
}

func readConn(conn net.Conn, ch chan Data) {
	read := bufio.NewReader(conn)
	for {
		var p [512]byte
		n, err := read.Read(p[:])
		if err != nil {
			log.Println("read err ...")
		}
		if n > 0 {
			ch <- Data{
				conn: conn,
				s:    string(p[:n]),
			}
		} else {
			break
		}
	}
}
