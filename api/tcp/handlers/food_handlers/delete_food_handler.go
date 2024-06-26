package food_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func Delete(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		id, parseErr := strconv.ParseUint(request.UrlParams["food_id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		userId := request.GetUserID() // TODO: check permission
		f, err := app.FoodService().GetById(request.Context(), uint(id))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}
		restaurantId := f.RestaurantId
		if err = app.AccessService().CheckRestaurantStaff(request.Context(), userId, restaurantId); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		err = app.FoodService().Delete(request.Context(), uint(id))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "food deleted successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
