package main

import (
	"fmt"
	"net"
	"os"

	Routes "github.com/codecrafters-io/http-server-starter-go/app/routes"
	Utils "github.com/codecrafters-io/http-server-starter-go/app/utils"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading buffer: ", err.Error())
		os.Exit(1)
	}
	if !(n > 0) {
		fmt.Println("Empty buffer, exit")
		os.Exit(1)
	}

	httpRawData, err := Utils.ParseHttpRawData(buffer, n)
	if err != nil {
		fmt.Println("Failed to parse buffer: ", err.Error())
		os.Exit(1)
	}

	Routes.HandleRoutes(conn, httpRawData.Request)
}
