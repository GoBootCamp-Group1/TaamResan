package handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"encoding/json"
	"net"
)

// HomeHandler handles requests to the home route.
func HomeHandler(conn net.Conn, request *tcp.Request) {
	responseData := map[string]any{
		"message": "Hello, Home!",
		"request": request,
	}
	tcp.RespondJson(conn, responseData, 200)
}

type RequestBody struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// TodoHandler handles requests to the //todo route.
func TodoHandler(conn net.Conn, request *tcp.Request) {

	var reqParams RequestBody
	// Unmarshal the JSON body into the struct
	err := json.Unmarshal([]byte(request.Body), &reqParams)
	if err != nil {
		tcp.RespondJsonError(conn, err.Error(), 500)
		return
	}

	// Validate the email and phone
	emailValidator := validator.Validate(reqParams.Email).Email()
	phoneValidator := validator.Validate(reqParams.Phone).Phone()

	// Collect all errors
	errors := append(emailValidator.Errors(), phoneValidator.Errors()...)

	// Respond with validation errors if any
	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, 422)
		return
	}

	//if emailValidator.HasErrors() || phoneValidator.HasErrors() {
	//	tcp.RespondJsonValidateError(conn, v.Errors(), 400)
	//	return
	//}

	responseData := map[string]any{
		"message": "Hello, TODO!",
		"request": request,
	}
	tcp.RespondJsonSuccess(conn, responseData)
	return
}
