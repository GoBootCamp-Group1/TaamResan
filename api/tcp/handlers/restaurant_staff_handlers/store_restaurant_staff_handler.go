package restaurant_staff_handlers

import (
	"TaamResan/internal/restaurant_staff"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
)

type CreateRequestBody struct {
	UserId       uint   `json:"user_id"`
	RestaurantId uint   `json:"restaurant_id"`
	Position     string `json:"position"`
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

		//userId := request.GetUserID() // TODO: Check permission and Owner

		newResStaff := restaurant_staff.RestaurantStaff{
			UserId:       reqParams.UserId,
			RestaurantId: reqParams.RestaurantId,
			Position:     restaurant_staff.GetPosition(reqParams.Position),
		}

		id, err := app.RestaurantStaffService().Create(request.Context(), &newResStaff)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"data":    map[string]any{"restaurant_staff_id": id},
			"message": "restaurant staff registered successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateCreateInputs(conn net.Conn, reqParams CreateRequestBody) {
	var errors []string

	positionValidator := validator.Validate(reqParams.Position).MinLength(3)
	errors = append(errors, positionValidator.Errors()...)
	if reqParams.UserId <= 0 {
		errors = append(errors, "UserId must be greater than 0")
	}

	if reqParams.RestaurantId <= 0 {
		errors = append(errors, "RestaurantId must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
