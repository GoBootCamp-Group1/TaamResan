package block_restaurant_handlers

import (
	"TaamResan/internal/block_restaurant"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
)

type CreateRequestBody struct {
	RestaurantId uint `json:"restaurant_id"`
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

		userId := request.GetUserID()

		newBlockRestaurant := block_restaurant.BlockRestaurant{
			UserId:       userId,
			RestaurantId: reqParams.RestaurantId,
		}

		id, err := app.BlockRestaurantService().Create(request.Context(), &newBlockRestaurant)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"data":    map[string]any{"block_restaurant_id": id},
			"message": "restaurant blocked successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateCreateInputs(conn net.Conn, reqParams CreateRequestBody) {
	var errors []string

	if reqParams.RestaurantId <= 0 {
		errors = append(errors, "RestaurantId must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
