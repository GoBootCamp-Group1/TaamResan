package cart_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
)

func GetAllByUser(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		userId := request.GetUserID() // TODO: check that user has permission and is OWNER to do this

		cartModels, err := app.CartService().GetByUserId(request.Context(), userId)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "carts loaded successfully",
			"data":    map[string]any{"carts": &cartModels},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
