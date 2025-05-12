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
	echoRoute := regexp.MustCompile(`^/echo/\w+`)
	userAgentRoute := regexp.MustCompile(`^/user-agent`)

	var response Utils.HttpResponse

	switch {
	case echoRoute.MatchString(target):
		response = handleEchoRoutes(request)
	case userAgentRoute.MatchString(target):
		response = handleUserAgentRoutes(request)
	case target == "/":
		response = Utils.HttpResponse{
			HttpVersion: request.HttpVersion,
			StatusCode:  200,
		}
	default:
		fmt.Printf("fallback")
		response = Utils.HttpResponse{
			HttpVersion: request.HttpVersion,
			StatusCode:  404,
		}
	}

	fmt.Fprint(conn, *response.ToString())
	os.Exit(1)
}
