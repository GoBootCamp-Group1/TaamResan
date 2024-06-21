package role_handlers

import (
	"TaamResan/internal/role"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
)

type UpdateRequestBody struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func UpdateRole(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams UpdateRequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateUpdateInputs(conn, reqParams)

		//userId := request.GetUserID() // TODO: check that user has permission and is ADMIN to do this

		newRole := role.Role{
			ID:   reqParams.ID,
			Name: reqParams.Name,
		}

		err = app.RoleService().Update(request.Context(), &newRole)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseData := map[string]any{
			"message": "role updated successfully",
		}
		tcp.RespondJsonSuccess(conn, responseData)
		return
	}
}

func validateUpdateInputs(conn net.Conn, reqParams UpdateRequestBody) {
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
