// goecho is a simple tcp echo example. An https request is also sent before
// the echo server runs in order to ensure that a docker container built from
// the provided dockerfile has working access to the net lib and ca-certs.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/user-agent")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	_ = resp.Body.Close()

	fmt.Println(string(body))

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

		go handle(conn)
	}
}

func handle(conn io.ReadWriteCloser) {
	buf := make([]byte, 102400)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		_, _ = conn.Write(buf[:n])
	}

	_ = conn.Close()
}
