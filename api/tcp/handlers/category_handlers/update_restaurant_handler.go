package category_handlers

import (
	"TaamResan/internal/category"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
)

type UpdateRequestBody struct {
	ID       uint   `json:"id"`
	ParentId *uint  `json:"parent_id"`
	Name     string `json:"name"`
}

func Update(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams UpdateRequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateUpdateInputs(conn, reqParams)

		userId := request.GetUserID() // TODO: check that user has permission and is OWNER to do this

		newCategory := category.Category{
			ID:        reqParams.ID,
			ParentId:  reqParams.ParentId,
			CreatedBy: userId,
			Name:      reqParams.Name,
		}

		err = app.CategoryService().Update(request.Context(), &newCategory)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "category updated successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateUpdateInputs(conn net.Conn, reqParams UpdateRequestBody) {
	var errors []string

	nameValidator := validator.Validate(reqParams.Name).MinLength(3)
	errors = append(errors, nameValidator.Errors()...)

	if reqParams.ID <= 0 {
		errors = append(errors, "ID must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
