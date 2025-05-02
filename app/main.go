package main

import (
	"errors"
	"fmt"
	"net"
	"os"
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

	httpRawData, err := parseBufferToHttpRawData(buffer, n)
	if err != nil {
		fmt.Println("Failed to parse buffer: ", err.Error())
		os.Exit(1)
	}

	fmt.Printf("HttpRawData: %v\n", httpRawData)

	if httpRawData.Request.Target != "/" {
		fmt.Fprintf(conn, "%s 404 Not Found\r\n\r\n", httpRawData.Request.HttpVersion)
		os.Exit(1)
	}

	fmt.Fprintf(conn, "%s 200 OK\r\n\r\n", httpRawData.Request.HttpVersion)
}

func parseBufferToHttpRawData(buffer []byte, n int) (HttpRawData, error) {
	separator := "\r\n"
	urlPath := string(buffer[:n])
	urlPaths := strings.Split(urlPath, separator)

	requestLine := urlPaths[0]
	urlPaths = urlPaths[1:]

	var requestHeaders []string
	for {
		if len(urlPaths) == 0 || len(urlPaths[0]) == 0 {
			// each header is ended with CRLF, and header section will be also ended with CRLF
			urlPaths = urlPaths[1:]

			// we're done searching for headers
			fmt.Printf("No more headers: %v, %d\n", strings.Join(urlPaths, ","), len(urlPaths))
			break
		}
		requestHeaders = append(requestHeaders, urlPaths[0])
		fmt.Printf("Added headers: %v\n", strings.Join(requestHeaders, ","))
		urlPaths = urlPaths[1:]
	}

	fmt.Printf("Request line: %s\n", requestLine)
	fmt.Printf("Request headers: %s\n", requestHeaders)

	var rawHttpData HttpRawData
	httpRequest, err := parseHttpRequest(requestLine)
	if err != nil {
		fmt.Println("Invalid http request: ", err.Error())
		return rawHttpData, err
	}
	httpRequestHeaders := parseHttpRequestHeaders(requestHeaders)
	rawHttpData.Request = httpRequest
	rawHttpData.Headers = httpRequestHeaders
	return rawHttpData, nil
}

func parseHttpRequest(value string) (HttpRequest, error) {
	var httpRequest HttpRequest

	arr := strings.Split(value, " ")
	if len(arr) != 3 {
		return httpRequest, errors.New("invalid http request")
	}

	return HttpRequest{
		Method:      arr[0],
		Target:      arr[1],
		HttpVersion: arr[2],
	}, nil
}

func parseHttpRequestHeaders(arr []string) []HttpRequestHeader {
	var headers []HttpRequestHeader
	for _, v := range arr {
		str := strings.Split(v, " ")
		if len(str) != 3 {
			continue
		}
		headers = append(headers, HttpRequestHeader{
			Host:            str[0],
			UserAgent:       str[1],
			AcceptedContent: str[2],
		})
	}
	return headers
}

type HttpRawData struct {
	Request HttpRequest
	Headers []HttpRequestHeader
}

type HttpRequest struct {
	Method      string
	Target      string
	HttpVersion string
}

type HttpRequestHeader struct {
	Host            string
	UserAgent       string
	AcceptedContent string
}
