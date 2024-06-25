package category_handlers

import (
	"TaamResan/internal/category"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
)

type CreateRequestBody struct {
	ParentId     *uint  `json:"parent_id"`
	RestaurantId uint   `json:"restaurant_id"`
	Name         string `json:"name"`
}

func Create(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams CreateRequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateCreateInputs(conn, reqParams)

		userId := request.GetUserID() // TODO: check permission
		if err = app.AccessService().CheckRestaurantOwner(request.Context(), userId, reqParams.RestaurantId); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		newCategory := category.Category{
			ParentId:     reqParams.ParentId,
			RestaurantId: reqParams.RestaurantId,
			CreatedBy:    userId,
			Name:         reqParams.Name,
		}

		id, err := app.CategoryService().Create(request.Context(), &newCategory)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"data":    map[string]any{"category_id": id},
			"message": "category created successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateCreateInputs(conn net.Conn, reqParams CreateRequestBody) {
	var errors []string

	nameValidator := validator.Validate(reqParams.Name).MinLength(3)
	errors = append(errors, nameValidator.Errors()...)

	if reqParams.RestaurantId <= 0 {
		errors = append(errors, "RestaurantId must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
