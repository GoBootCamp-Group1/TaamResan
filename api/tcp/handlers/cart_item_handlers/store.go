package cart_item_handlers

import (
	"TaamResan/internal/cart_item"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
)

type CreateRequestBody struct {
	FoodId uint    `json:"food_id"`
	Amount float64 `json:"amount"`
	Note   string  `json:"note"`
}

func Create(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams CreateRequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateCreateInputs(conn, reqParams)

		newCartItem := cart_item.CartItem{
			FoodId: reqParams.FoodId,
			Amount: reqParams.Amount,
			Note:   reqParams.Note,
		}

		id, err := app.CartItemService().Create(request.Context(), &newCartItem)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"data":    map[string]any{"cart_item_id": id},
			"message": "cart_item created successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateCreateInputs(conn net.Conn, reqParams CreateRequestBody) {
	var errors []string

	noteValidator := validator.Validate(reqParams.Note).MinLength(3)
	errors = append(errors, noteValidator.Errors()...)

	if reqParams.FoodId <= 0 {
		errors = append(errors, "FoodId must be greater than 0")
	}

	if reqParams.Amount <= 0 {
		errors = append(errors, "Amount must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
