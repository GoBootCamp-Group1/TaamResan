package authentication_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"encoding/json"
	"errors"
	"net"
)

type RequestBody struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	ErrPhoneOrEmailIsEmpty = errors.New("phone or email is empty")
)

func Login(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams RequestBody
		err := json.Unmarshal([]byte(request.Body), &reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INVALID_INPUT)
			return
		}

		if reqParams.Phone == "" && reqParams.Email == "" {
			tcp.RespondJsonError(conn, ErrPhoneOrEmailIsEmpty.Error(), tcp.INVALID_INPUT)
			return
		}

		var errs []string
		if reqParams.Email != "" {
			emailValidator := validator.Validate(reqParams.Email).Email()
			errs = append(errs, emailValidator.Errors()...)
		} else if reqParams.Phone != "" {
			phoneValidator := validator.Validate(reqParams.Phone).Phone()
			errs = append(errs, phoneValidator.Errors()...)
		}

		passwordValidator := validator.Validate(reqParams.Password).Password()
		errs = append(errs, passwordValidator.Errors()...)

		if len(errs) > 0 {
			tcp.RespondJsonValidateError(conn, errs, tcp.INVALID_INPUT)
			return
		}

		var token *service.UserToken
		if reqParams.Phone != "" {
			token, err = app.AuthService().LoginWithMobile(request.Context(), reqParams.Phone, reqParams.Password)
			if err != nil {
				tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
				return
			}
		} else if reqParams.Email != "" {
			token, err = app.AuthService().LoginWithEmail(request.Context(), reqParams.Email, reqParams.Password)
			if err != nil {
				tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
				return
			}
		}

		responseBody := map[string]any{
			"message": "you are logged in successfully",
			"data":    map[string]any{"token": token},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
