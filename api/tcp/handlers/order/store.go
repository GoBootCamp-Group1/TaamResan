package order

import (
	"TaamResan/internal/order"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
)

type storeOrderRequestBody struct {
	CartID    uint `json:"cart_id"`
	AddressID uint `json:"address_id"`
}

func StoreHandler(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		//parse body
		var requestParams storeOrderRequestBody

		errParseParams := request.ExtractBodyParamsInto(&requestParams)
		if errParseParams != nil {
			tcp.RespondJsonError(conn, errParseParams.Error(), tcp.INVALID_INPUT)
			return
		}

		//validate data
		validateStore(conn, requestParams)

		//generate order data object
		orderInputData := order.InputData{
			CartID:    requestParams.CartID,
			AddressID: requestParams.AddressID,
		}

		//create order using service
		newOrder, err := app.OrderService().Create(request.Context(), &orderInputData)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		//show response
		responseBody := map[string]any{
			"data":    map[string]any{"order": newOrder},
			"message": "order created successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateStore(conn net.Conn, reqParams storeOrderRequestBody) {
	var errors []string

	if reqParams.CartID <= 0 {
		errors = append(errors, "Cart ID must be greater than 0")
	}

	if reqParams.AddressID <= 0 {
		errors = append(errors, "Address ID must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
