package handlers

import (
	"TaamResan/pkg/tcp_http_server"
	"net"
)

// HomeHandler handles requests to the home route.
func HomeHandler(conn net.Conn, request *tcp_http_server.Request) {
	responseData := map[string]any{
		"message": "Hello, Home!",
		"request": request,
	}
	tcp_http_server.RespondJson(conn, responseData, 200)
}

// TodoHandler handles requests to the //todo route.
func TodoHandler(conn net.Conn, request *tcp_http_server.Request) {
	responseData := map[string]any{
		"message": "Hello, TODO!",
		"request": request,
	}
	tcp_http_server.RespondJson(conn, responseData, 200)
}
