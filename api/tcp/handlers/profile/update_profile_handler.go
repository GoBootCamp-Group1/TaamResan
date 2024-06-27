package profile

import (
	"TaamResan/internal/user"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
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

func UpdateProfile(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams RequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateInputs(conn, reqParams)

		userId := request.GetUserID()

		updatedUser := user.User{
			ID:        userId,
			Name:      reqParams.Name,
			Email:     reqParams.Email,
			Mobile:    reqParams.Phone,
			Password:  reqParams.Password,
			BirthDate: reqParams.BirthDate,
		}

		err = app.UserService().UpdateUserProfile(request.Context(), &updatedUser)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "your profile had been updated successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateInputs(conn net.Conn, reqParams RequestBody) {
	var errors []string

	if reqParams.Name != "" {
		nameValidator := validator.Validate(reqParams.Name).MinLength(3)
		errors = append(errors, nameValidator.Errors()...)
	}

	if reqParams.Phone != "" {
		phoneValidator := validator.Validate(reqParams.Phone).Phone()
		errors = append(errors, phoneValidator.Errors()...)
	}

	if reqParams.Email != "" {
		emailValidator := validator.Validate(reqParams.Email).Email()
		errors = append(errors, emailValidator.Errors()...)
	}

	if reqParams.Password != "" {
		passwordValidator := validator.Validate(reqParams.Password).Password()
		errors = append(errors, passwordValidator.Errors()...)
	}

	//if reqParams.BirthDate.String() != "" {
	//	dateValidator := validator.Validate(reqParams.BirthDate.String()).Date()
	//	errors = append(errors, dateValidator.Errors()...)
	//}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
