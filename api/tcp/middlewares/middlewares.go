package middlewares

import (
	"TaamResan/pkg/jwt"
	tcp "TaamResan/pkg/tcp_http_server"
	"context"
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
func AuthMiddleware(secret string) tcp.MiddlewareFunc {
	return func(next tcp.HandlerFunc) tcp.HandlerFunc {
		return func(conn net.Conn, request *tcp.Request) {
			authHeader := request.Headers["Authorization"]
			if len(authHeader) == 0 {
				tcp.RespondJsonError(conn, "Unauthorized", tcp.UNAUTHORIZED)
				return
			}

			claims, err := jwt.ParseToken(authHeader, []byte(secret))
			if err != nil {
				tcp.RespondJsonError(conn, "Unauthorized", tcp.UNAUTHORIZED)
				return
			}
			ctx := context.WithValue(request.Context(), jwt.UserClaimKey, claims)
			request = request.WithContext(ctx)

			next(conn, request)
		}
	}
}
