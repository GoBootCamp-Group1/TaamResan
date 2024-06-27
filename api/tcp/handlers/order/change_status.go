package order

import (
	"TaamResan/internal/order"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"fmt"
	"net"
	"strconv"
)

type StatusRequest struct {
	Status uint `json:"status"`
}

func ChangeStatusHandler(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		orderIdStr, ok := request.UrlParams["orderId"]

		if !ok {
			tcp.RespondJsonError(conn, "order id is required", tcp.INVALID_INPUT)
			return
		}

		orderId, err := strconv.Atoi(orderIdStr)

		var reqParams StatusRequest

		err = request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		if reqParams.Status <= 0 {
			tcp.RespondJsonValidateError(conn, []string{"not valid input"}, tcp.INVALID_INPUT)
			return
		}

		order := order.Order{
			ID:     uint(orderId),
			Status: reqParams.Status,
		}

		err = app.OrderService().ChangeStatusByRestaurant(request.Context(), &order)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": fmt.Sprintf("order status changed successfully"),
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
