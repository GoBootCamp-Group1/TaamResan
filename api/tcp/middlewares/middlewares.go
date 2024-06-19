package middlewares

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"fmt"
	"net"
)

// LoggingMiddleware logs the request details.
func LoggingMiddleware(next tcp.HandlerFunc) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		fmt.Printf("Logger: Received %s request for %s\n", request.Method, request.Uri)
		next(conn, request)
	}
}

// AuthMiddleware checks for a specific header to authenticate the request.
func AuthMiddleware(next tcp.HandlerFunc) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		if request.Headers["Authorization"] != "Bearer secret" {
			tcp.RespondJsonError(conn, "Unauthorized", tcp.UNAUTHORIZED)
			return
		}
		next(conn, request)
	}
}
