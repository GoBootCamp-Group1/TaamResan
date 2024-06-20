package middlewares

import (
	"TaamResan/pkg/jwt"
	tcp "TaamResan/pkg/tcp_http_server"
	"context"
	"fmt"
	"net"
	"strings"
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

			//cfg, err := config.GetConfig()
			//secretFromS := cfg.Server.TokenSecret
			//fmt.Println(secretFromS)

			authHeader := request.Headers["Authorization"]
			if len(authHeader) == 0 {
				tcp.RespondJsonError(conn, "Unauthorized", tcp.UNAUTHORIZED)
				return
			}

			// Check if the Authorization header starts with "Bearer "
			if !strings.HasPrefix(authHeader, "Bearer ") {
				tcp.RespondJsonError(conn, "Unauthorized", tcp.UNAUTHORIZED)
				return
			}

			// Extract the token part
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := jwt.ParseToken(tokenString, []byte(secret))
			if err != nil {
				tcp.RespondJsonError(conn, "Unauthorized", tcp.UNAUTHORIZED)
				return
			}

			//TODO: check for existing user_id in database

			ctx := context.WithValue(request.Context(), jwt.UserClaimKey, claims)
			request = request.WithContext(ctx)

			next(conn, request)
		}
	}
}
