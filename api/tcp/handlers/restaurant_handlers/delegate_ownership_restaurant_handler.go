package restaurant_handlers

import (
	"TaamResan/internal/action_log"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"encoding/json"
	"fmt"
	"net"
)

type ApproveRequestBody struct {
	RestaurantId uint `json:"restaurant_id"`
	NewOwnerId   uint `json:"new_owner_id"`
}

func Delegate(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams ApproveRequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateApproveInputs(conn, reqParams)

		userId := request.GetUserID() // TODO: check permission
		if err = app.AccessService().CheckRestaurantOwner(request.Context(), userId, reqParams.RestaurantId); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		err = app.RestaurantService().DelegateOwnership(request.Context(), reqParams.RestaurantId, reqParams.NewOwnerId)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		if err = logDelegateRestaurantRequest(app, request, reqParams.RestaurantId); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "restaurant ownership delegated successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateApproveInputs(conn net.Conn, reqParams ApproveRequestBody) {
	var errors []string

	if reqParams.RestaurantId <= 0 {
		errors = append(errors, "RestaurantId must be greater than 0")
	}
	if reqParams.NewOwnerId <= 0 {
		errors = append(errors, "NewOwnerId must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}

func logDelegateRestaurantRequest(app *service.AppContainer, request *tcp.Request, restaurantId uint) error {
	userId := request.GetUserID()
	var payload map[string]any
	err := json.Unmarshal([]byte(request.Body), &payload)
	if err != nil && request.Body != "" {
		fmt.Printf(err.Error() + "\n")
	}
	log := action_log.ActionLog{
		UserID:     &userId,
		Action:     "Delegate Restaurant",
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
