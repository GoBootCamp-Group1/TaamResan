package role_handlers

import (
	"TaamResan/internal/role"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
)

type RequestBody struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateRole(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams RequestBody

		err := request.ExtractParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateInputs(conn, reqParams)

		//userId := request.GetUserID() // TODO: check that user has permission and is ADMIN to do this

		newRole := role.Role{
			ID:   reqParams.ID,
			Name: reqParams.Name,
		}

		err = app.RoleService().Create(request.Context(), &newRole)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseData := map[string]any{
			"message": "role created successfully",
		}
		tcp.RespondJsonSuccess(conn, responseData)
		return
	}
}

func UpdateRole(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams RequestBody

		err := request.ExtractParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateInputs(conn, reqParams)

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

func DeleteRole(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		//userId := request.GetUserID() // TODO: check that user has permission and is ADMIN to do this

		var id uint = 11 // TODO

		err := app.RoleService().Delete(request.Context(), id)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseData := map[string]any{
			"message": "role deleted successfully",
		}
		tcp.RespondJsonSuccess(conn, responseData)
		return
	}
}

func GetRole(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		name := request.QueryParams["name"]

		var errors []string
		nameValidator := validator.Validate(name).MinLength(3)
		errors = append(errors, nameValidator.Errors()...)

		if len(errors) > 0 {
			tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
			return
		}

		//userId := request.GetUserID() // TODO: check that user has permission and is ADMIN to do this

		roleModel, err := app.RoleService().GetByName(request.Context(), name)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseData := map[string]any{
			"message": "role loaded successfully",
			"role":    roleModel,
		}
		tcp.RespondJsonSuccess(conn, responseData)
		return
	}
}

func validateInputs(conn net.Conn, reqParams RequestBody) {
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
