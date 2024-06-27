package role_handlers

import (
	"TaamResan/internal/role"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
)

type CreateRequestBody struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateRole(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams CreateRequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateCreateInputs(conn, reqParams)

		newRole := role.Role{
			ID:   reqParams.ID,
			Name: reqParams.Name,
		}

		err = app.RoleService().Create(request.Context(), &newRole)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "role created successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateCreateInputs(conn net.Conn, reqParams CreateRequestBody) {
	var errors []string

	if reqParams.ID <= 0 {
		errors = append(errors, "id must be positive")
	}

	nameValidator := validator.Validate(reqParams.Name).MinLength(3)
	errors = append(errors, nameValidator.Errors()...)

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
