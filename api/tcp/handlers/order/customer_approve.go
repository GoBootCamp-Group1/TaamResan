package order

import (
	"TaamResan/internal/order"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"fmt"
	"net"
	"strconv"
)

func CustomerApproveHandler(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {

		orderIdStr, ok := request.UrlParams["orderId"]

		if !ok {
			tcp.RespondJsonError(conn, "order id is required", tcp.INVALID_INPUT)
			return
		}

		orderId, err := strconv.Atoi(orderIdStr)

		//generate order data object
		orderModel := order.Order{
			ID: uint(orderId),
		}

		//create order using service
		updatedOrder, err := app.OrderService().ApproveByCustomer(request.Context(), &orderModel)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		//show response
		responseBody := map[string]any{
			"data":    map[string]any{"order": updatedOrder},
			"message": fmt.Sprintf("order approved successfully"),
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
