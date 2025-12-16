package main

import (
	"fmt"
	"net"
)

func main() {

	bound, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("Listening on port %s\n", bound.Addr().String())
	var buf []byte

	conn, err := bound.Accept()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	read, err := conn.Read(buf)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("Bytes read: %d\n", read)

	str := string(buf)

	fmt.Println(str)
}
