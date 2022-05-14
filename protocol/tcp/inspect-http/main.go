package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:9999")
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	defer listener.Close()
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("Masuk sini")
		handleClient(conn)
		fmt.Println("Keluar sini")

		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
	}

	os.Exit(0)
}

func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		ht := checkHTTP(conn, string(buf[0:]))
		if ht {
			return
		}
		fmt.Println(string(buf[0:]))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkHTTP(conn net.Conn, data string) bool {
	fmt.Println("ini isi datanya ", data)
	if strings.Contains(data, "HTTP") {
		_, err := conn.Write([]byte("HTTP/1.1 200 OK\r\nServer:Mac\r\n\r\nnah ini bodynya ngga tau bener ngga"))
		conn.Close()
		checkError(err)
		return true
	}
	return false
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
