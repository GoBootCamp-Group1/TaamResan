package handlers

import (
	"net"

	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
)

// HomeHandler handles requests to the home route.
func HomeHandler(conn net.Conn, request *tcp.Request) {
	responseData := map[string]any{
		"message": "Hello, Home!",
		"request": request,
	}
	tcp.RespondJson(conn, responseData, 200)
}

// TodoHandler handles requests to the //todo route.
func TodoHandler(conn net.Conn, request *tcp.Request) {

	email := "samankarbasi@live.com"

	v := validator.Validate(email).
		Email().
		MinLength(5).
		MaxLength(50).
		NonEmpty()

	if v.HasErrors() {
		tcp.RespondJsonValidateError(conn, v.Errors(), 400)
		return
	}

	responseData := map[string]any{
		"message": "Hello, TODO!",
		"request": request,
	}
	tcp.RespondJsonSuccess(conn, responseData)
	return
}
