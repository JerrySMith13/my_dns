package main

import (
	"fmt"
	"net"
	"os"
)

var REQ_MAXSIZE = 1024

func handle_conn(conn: net.TCPConn*) error {

}

func main() {

	bound, err := net.Listen("tcp", "127.0.0.1:8080")
	net.ParseIP(os.Args[1])
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("Listening on port %s\n", bound.Addr().String())

	buf := make([]byte, 4096)

}
