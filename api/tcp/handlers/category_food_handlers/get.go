package category_food_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func Get(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		id, parseErr := strconv.ParseUint(request.UrlParams["category_food_id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		categoryFood, err := app.CategoryFoodService().GetById(request.Context(), uint(id))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "category_food loaded successfully",
			"data":    map[string]any{"category_food": categoryFood},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
