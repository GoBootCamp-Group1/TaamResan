package restaurant_staff_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func GetAllByRestaurant(app *service.AppContainer) tcp.HandlerFunc {
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

		staffs, err := app.RestaurantStaffService().GetAllByRestaurantId(request.Context(), uint(restaurantId))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "restaurant_staffs loaded successfully",
			"data":    map[string]any{"restaurant_staffs": &staffs},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
