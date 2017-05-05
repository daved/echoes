package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":25000")
	if err != nil {
		panic(err)
	}
	for i := 0; true; i++ {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("cannot accept: %s", err)
			continue
		}
		fmt.Printf("%d: %v <-> %v\n", i, conn.LocalAddr(), conn.RemoteAddr())
		go func(conn io.ReadWriteCloser) {
			buf := make([]byte, 102400)
			for {
				n, err := conn.Read(buf)
				if err != nil {
					break
				}
				conn.Write(buf[:n])
			}
			conn.Close()
		}(conn)
	}
}
