package routes

import (
	"strings"

	Utils "github.com/codecrafters-io/http-server-starter-go/app/utils"
)

func handleUserAgentRoutes(request Utils.HttpRequest) Utils.HttpResponse {
	var response Utils.HttpResponse

	str := Utils.GetLast(strings.Split(request.Target, "/user-agent"))
	if str != nil {
		response = Utils.HttpResponse{
			HttpVersion:   request.HttpVersion,
			StatusCode:    200,
			ContentType:   "text/plain",
			ContentLength: len(*str),
			Body:          str,
		}
	}

	return response
}
