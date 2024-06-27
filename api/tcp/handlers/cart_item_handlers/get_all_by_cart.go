package cart_item_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func GetAllByCart(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		cartId, parseErr := strconv.ParseUint(request.UrlParams["cart_id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		cartItemModels, err := app.CartItemService().GetAllByCartId(request.Context(), uint(cartId))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "cart_items loaded successfully",
			"data":    map[string]any{"cart_items": &cartItemModels},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
