package block_restaurant_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func GetAllByRestaurant(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		userId, parseErr := strconv.ParseUint(request.UrlParams["user_id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		blockRestaurants, err := app.BlockRestaurantService().GetAllByUserId(request.Context(), uint(userId))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "blocked restaurants loaded successfully",
			"data":    map[string]any{"blocked_restaurants": &blockRestaurants},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
