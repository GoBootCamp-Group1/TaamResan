package middlewares

import (
	config2 "TaamResan/cmd/api/config"
	"TaamResan/internal/action_log"
	"TaamResan/pkg/jwt"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// LoggingMiddleware logs the request details.
func LoggingMiddleware(actionLogService *service.ActionLogService) tcp.MiddlewareFunc {
	return func(next tcp.HandlerFunc) tcp.HandlerFunc {
		return func(conn net.Conn, request *tcp.Request) {

			var payload map[string]any
			err := json.Unmarshal([]byte(request.Body), &payload)
			if err != nil && request.Body != "" {
				fmt.Printf(err.Error() + "\n")
			}

			actionLog := action_log.ActionLog{
				UserID:   getUserId(request),
				Action:   request.Method,
				IP:       request.IP,
				Endpoint: request.Uri,
				Payload:  payload,
				Method:   request.Method,
			}

			log, errLog := actionLogService.Create(request.Context(), &actionLog)
			if errLog != nil {
				fmt.Printf(errLog.Error() + "\n")
			}

			writeLogToFile(request)

			//set log to context
			ctx := context.WithValue(request.Context(), action_log.LogCtxKey, log)
			request = request.WithContext(ctx)

			next(conn, request)
		}
	}
}

func getUserId(request *tcp.Request) *uint {

	//read secret
	config, _ := config2.GetConfig()

	authHeader := request.Headers["Authorization"]
	if len(authHeader) == 0 {
		return nil
	}

	// Check if the Authorization header starts with "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return nil
	}

	// Extract the token part
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims, err := jwt.ParseToken(tokenString, []byte(config.Server.TokenSecret))
	if err != nil {
		return nil
	}

	return &claims.UserID
}

func writeLogToFile(request *tcp.Request) {
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	logger := log.New(multiWriter, "Logger: ", log.LstdFlags)

	logger.Printf("Received %s request for %s", request.Method, request.Uri)

}
