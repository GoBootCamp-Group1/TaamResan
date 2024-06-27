package restaurant_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func GetRestaurant(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		id, parseErr := strconv.ParseUint(request.UrlParams["id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		restaurantModel, err := app.RestaurantService().GetById(request.Context(), uint(id))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "restaurant loaded successfully",
			"data":    map[string]any{"restaurant": restaurantModel},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
