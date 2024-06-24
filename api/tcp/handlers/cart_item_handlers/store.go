package cart_item_handlers

import (
	"TaamResan/internal/cart_item"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
)

type CreateRequestBody struct {
	CartId uint   `json:"cart_id"`
	FoodId uint   `json:"food_id"`
	Amount uint   `json:"amount"`
	Note   string `json:"note"`
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

		//userId := request.GetUserID() // TODO: permission and User

		newCartItem := cart_item.CartItem{
			CartId: reqParams.CartId,
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

	if reqParams.CartId <= 0 {
		errors = append(errors, "CartId must be greater than 0")
	}

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
