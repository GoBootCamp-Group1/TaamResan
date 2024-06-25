package restaurant_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
)

func GetAllRestaurants(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		//userId := request.GetUserID() // TODO: check permission and is Admin

		restaurantModels, err := app.RestaurantService().GetAll(request.Context())
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "restaurants loaded successfully",
			"data":    map[string]any{"restaurants": &restaurantModels},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
