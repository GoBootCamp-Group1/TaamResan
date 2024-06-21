package signup_handlers

import (
	"TaamResan/internal/user"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"github.com/google/uuid"
	"net"
	"time"
)

type RequestBody struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
}

func SignUp(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams RequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		nameValidator := validator.Validate(reqParams.Name).MinLength(3)
		phoneValidator := validator.Validate(reqParams.Phone).Phone()
		emailValidator := validator.Validate(reqParams.Email).Email()
		passwordValidator := validator.Validate(reqParams.Password).Password()

		var errors []string
		errors = append(errors, nameValidator.Errors()...)
		errors = append(errors, phoneValidator.Errors()...)
		errors = append(errors, emailValidator.Errors()...)
		errors = append(errors, passwordValidator.Errors()...)

		if len(errors) > 0 {
			tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
			return
		}

		newUser := user.User{
			Uuid:      uuid.NewString(),
			Name:      reqParams.Name,
			Email:     reqParams.Email,
			Mobile:    reqParams.Phone,
			Password:  reqParams.Password,
			BirthDate: reqParams.BirthDate,
		}

		err = app.UserService().CreateUser(request.Context(), &newUser)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseData := map[string]any{
			"message": "you are signed up successfully",
		}
		tcp.RespondJsonSuccess(conn, responseData)
		return
	}
}
