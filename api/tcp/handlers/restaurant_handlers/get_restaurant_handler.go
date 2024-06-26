package restaurant_handlers

import (
	"TaamResan/internal/action_log"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"encoding/json"
	"fmt"
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

		userId := request.GetUserID() // TODO: check permission

		restaurantModel, err := app.RestaurantService().GetById(request.Context(), uint(id))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		if err = app.AccessService().CheckRestaurantOwner(request.Context(), userId, uint(id)); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		if err = logGetRestaurantRequest(app, request, uint(id)); err != nil {
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

func logGetRestaurantRequest(app *service.AppContainer, request *tcp.Request, restaurantId uint) error {
	userId := request.GetUserID()
	var payload map[string]any
	err := json.Unmarshal([]byte(request.Body), &payload)
	if err != nil && request.Body != "" {
		fmt.Printf(err.Error() + "\n")
	}
	log := action_log.ActionLog{
		UserID:     &userId,
		Action:     "Load Restaurant",
		IP:         request.IP,
		Endpoint:   request.Uri,
		Payload:    payload,
		Method:     request.Method,
		EntityType: action_log.RestaurantEntityType,
		EntityID:   restaurantId,
	}
	_, err = app.ActionLogService().Create(request.Context(), &log)

	return err
}
