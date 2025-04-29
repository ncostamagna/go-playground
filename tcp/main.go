package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	// go run main.go 9090
	// echo Hello, World! | nc localhost 9090
	if len(os.Args) < 2 {
		fmt.Println("error: no arguments provided")
		os.Exit(1)
	}

	port := ":" + os.Args[1]

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("error: failed to start server", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("server started on port", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error: failed to accept connection", err)
			continue
		}
		defer conn.Close()
		fmt.Println("client connected")

		go handleConnection(conn)
		
	}
	
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	//writer := bufio.NewWriter(conn)

	for {
		bytes, err := reader.ReadBytes(byte('\n'))
		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read bytes", err)
			}
			return
		}

		fmt.Println("received message:", string(bytes))

		line := fmt.Sprintf("Echo: %s", string(bytes))
		fmt.Println("sending response:", line)

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("failed to send response", err)
			return
		}
	}

}
