package main

import (
	"fmt"
)

type HttpResponse struct {
	HttpVersion   string
	StatusCode    int
	ContentType   string
	ContentLength int
	Body          *string
}

func (r HttpResponse) ToString() *string {
	var response string
	var status string

	switch r.StatusCode {
	case 200:
		status = "OK"
	default:
		status = "Not Found"
	}

	if r.StatusCode >= 400 && r.StatusCode <= 599 || r.Body == nil {
		response = fmt.Sprintf("%s %d %s\r\n\r\n", r.HttpVersion, r.StatusCode, status)
		return &response
	}

	response = fmt.Sprintf("%s %d %s\r\nContent-Type: %s\r\nContent-Length: %d\r\n\r\n%v", r.HttpVersion, r.StatusCode, status, r.ContentType, r.ContentLength, *r.Body)
	return &response
}
