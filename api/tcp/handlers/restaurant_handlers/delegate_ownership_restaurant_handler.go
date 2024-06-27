package restaurant_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
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

		err = app.RestaurantService().DelegateOwnership(request.Context(), reqParams.RestaurantId, reqParams.NewOwnerId)
		if err != nil {
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
