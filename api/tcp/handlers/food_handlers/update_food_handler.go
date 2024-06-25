package food_handlers

import (
	"TaamResan/internal/food"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
	"strconv"
)

type UpdateRequestBody struct {
	Name               string  `json:"name"`
	Price              float64 `json:"price"`
	CancelRate         float64 `json:"cancel_rate"`
	PreparationMinutes uint    `json:"preparation_minutes"`
}

func Update(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		id, parseErr := strconv.ParseUint(request.UrlParams["food_id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		var reqParams UpdateRequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateUpdateInputs(conn, reqParams)

		userId := request.GetUserID() // TODO: check permission
		f, err := app.FoodService().GetById(request.Context(), uint(id))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}
		restaurantId := f.RestaurantId
		if err = app.AccessService().CheckRestaurantOwner(request.Context(), userId, restaurantId); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		newFood := food.Food{
			ID:                 uint(id),
			CreatedBy:          userId,
			Name:               reqParams.Name,
			Price:              reqParams.Price,
			CancelRate:         reqParams.CancelRate,
			PreparationMinutes: reqParams.PreparationMinutes,
		}

		err = app.FoodService().Update(request.Context(), &newFood)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "food updated successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateUpdateInputs(conn net.Conn, reqParams UpdateRequestBody) {
	var errors []string

	nameValidator := validator.Validate(reqParams.Name).MinLength(3)
	errors = append(errors, nameValidator.Errors()...)

	if reqParams.Price <= 0 {
		errors = append(errors, "Price must be greater than 0")
	}

	if reqParams.CancelRate < 0 {
		errors = append(errors, "CancelRate must be equal or greater than 0")
	}

	if reqParams.PreparationMinutes <= 0 {
		errors = append(errors, "PreparationMinutes must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
