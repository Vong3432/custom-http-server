package routes

import (
	"fmt"
	"net"
	"os"
	"regexp"

	Utils "github.com/codecrafters-io/http-server-starter-go/app/utils"
)

func HandleRoutes(conn net.Conn, request Utils.HttpRequest) {
	target := request.Target

	// Routes
	echoRoute := "^/echo/\\w+"
	matchedEchoRoute, _ := regexp.MatchString(echoRoute, target)

	if matchedEchoRoute {
		response := handleEchoRoutes(request)
		fmt.Fprint(conn, *response.ToString())
		os.Exit(1)
	}

	if target == "/" {
		response := Utils.HttpResponse{
			HttpVersion: request.HttpVersion,
			StatusCode:  200,
		}
		fmt.Fprint(conn, *response.ToString())
		os.Exit(1)
	}

	response := Utils.HttpResponse{
		HttpVersion: request.HttpVersion,
		StatusCode:  404,
	}
	fmt.Fprint(conn, *response.ToString())
}
