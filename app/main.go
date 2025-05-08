package main

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
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

	httpRawData, err := parseHttpRawData(buffer, n)
	if err != nil {
		fmt.Println("Failed to parse buffer: ", err.Error())
		os.Exit(1)
	}

	target := httpRawData.Request.Target

	// Routes
	echoRoute := "^/echo/\\w+"
	matchedEchoRoute, _ := regexp.MatchString(echoRoute, target)

	if matchedEchoRoute {
		str := GetLast(strings.Split(target, "/echo/"))
		if str != nil {
			response := HttpResponse{
				HttpVersion:   httpRawData.Request.HttpVersion,
				StatusCode:    200,
				ContentType:   "text/plain",
				ContentLength: len(*str),
				Body:          str,
			}
			fmt.Fprint(conn, *response.ToString())
			os.Exit(1)
		}
	}

	if target == "/" {
		response := HttpResponse{
			HttpVersion: httpRawData.Request.HttpVersion,
			StatusCode:  200,
		}
		fmt.Fprint(conn, *response.ToString())
		os.Exit(1)
	}

	response := HttpResponse{
		HttpVersion: httpRawData.Request.HttpVersion,
		StatusCode:  404,
	}
	fmt.Fprint(conn, *response.ToString())
}
