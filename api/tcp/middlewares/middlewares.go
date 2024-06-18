package middlewares

import (
	"TaamResan/pkg/tcp_http_server"
	"fmt"
	"net"
)

// LoggingMiddleware logs the request details.
func LoggingMiddleware(next tcp_http_server.HandlerFunc) tcp_http_server.HandlerFunc {
	return func(conn net.Conn, request *tcp_http_server.Request) {
		fmt.Printf("Logger: Received %s request for %s\n", request.Method, request.Uri)
		next(conn, request)
	}
}

// AuthMiddleware checks for a specific header to authenticate the request.
func AuthMiddleware(next tcp_http_server.HandlerFunc) tcp_http_server.HandlerFunc {
	return func(conn net.Conn, request *tcp_http_server.Request) {
		if request.Headers["Authorization"] != "Bearer secret" {
			tcp_http_server.RespondJsonError(conn, "Unauthorized", 401)
			return
		}
		next(conn, request)
	}
}
