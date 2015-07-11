package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func sendMessage(msg string) {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	if err != nil {
		panic(err)
	}

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buf := []byte(msg)
	_, err = conn.Write(buf)
	if err != nil {
		fmt.Println(msg, err)
	}
}

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		line, err := r.ReadString('\n')

		if err != nil && err == io.EOF {
			break
		}
		fmt.Printf("%s", line)
		sendMessage(line)
	}
}
