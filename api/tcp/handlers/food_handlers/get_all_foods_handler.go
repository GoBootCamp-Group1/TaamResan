package food_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func GetAll(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		restaurantId, parseErr := strconv.ParseUint(request.UrlParams["restaurant_id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		userId := request.GetUserID() // TODO: check permission
		if err := app.AccessService().CheckRestaurantOwner(request.Context(), userId, uint(restaurantId)); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		foodModel, err := app.FoodService().GetAll(request.Context(), uint(restaurantId))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "foods loaded successfully",
			"data":    map[string]any{"foods": &foodModel},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
