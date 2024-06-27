package cart_item_handlers

import (
	"TaamResan/internal/cart_item"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
	"strconv"
)

type UpdateRequestBody struct {
	Amount float64 `json:"amount"`
	Note   string  `json:"note"`
}

func Update(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		id, parseErr := strconv.ParseUint(request.UrlParams["cart_item_id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		var reqParams UpdateRequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateUpdateInputs(conn, reqParams)

		cartItem := cart_item.CartItem{
			ID:     uint(id),
			Amount: reqParams.Amount,
			Note:   reqParams.Note,
		}

		err = app.CartItemService().Update(request.Context(), &cartItem)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "cart_item updated successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateUpdateInputs(conn net.Conn, reqParams UpdateRequestBody) {
	var errors []string

	noteValidator := validator.Validate(reqParams.Note).MinLength(3)
	errors = append(errors, noteValidator.Errors()...)

	if reqParams.Amount <= 0 {
		errors = append(errors, "Amount must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
