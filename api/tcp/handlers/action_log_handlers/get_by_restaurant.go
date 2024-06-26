package action_log_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func GetByRestaurant(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		restaurantId, parseErr := strconv.ParseUint(request.UrlParams["restaurant_id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		userId := request.GetUserID() // TODO: check permission

		logs, err := app.ActionLogService().GetAllByRestaurantId(request.Context(), uint(restaurantId))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		if err = app.AccessService().CheckAdminAccess(request.Context(), userId); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		responseBody := map[string]any{
			"message": "logs loaded successfully",
			"data":    map[string]any{"logs": logs},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
