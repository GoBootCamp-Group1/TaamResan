package restaurant_handlers

import (
	"TaamResan/internal/action_log"
	"TaamResan/internal/address"
	"TaamResan/internal/restaurant"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

type CreateRequestBody struct {
	Name         string  `json:"name"`
	AddressTitle string  `json:"address_title"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	CourierSpeed float64 `json:"courier_speed"`
}

func CreateRestaurant(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams CreateRequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateCreateInputs(conn, reqParams)

		userId := request.GetUserID()

		newRestaurant := restaurant.Restaurant{
			Name:           reqParams.Name,
			OwnedBy:        userId,
			ApprovalStatus: restaurant.NotApproved,
			Address: address.Address{
				Title: reqParams.AddressTitle,
				Lat:   reqParams.Lat,
				Lng:   reqParams.Lng,
			},
			CourierSpeed: reqParams.CourierSpeed,
		}

		id, err := app.RestaurantService().Create(request.Context(), &newRestaurant)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		if err = logCreateRestaurantRequest(app, request, id); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"data":    map[string]any{"restaurant_id": id},
			"message": "restaurant created successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateCreateInputs(conn net.Conn, reqParams CreateRequestBody) {
	var errors []string

	nameValidator := validator.Validate(reqParams.Name).MinLength(3)
	addressTitleValidator := validator.Validate(reqParams.AddressTitle).MinLength(3)
	latValidator := validator.Validate(strconv.FormatFloat(reqParams.Lat, 'f', -1, 64)).IsFloat()
	lngValidator := validator.Validate(strconv.FormatFloat(reqParams.Lng, 'f', -1, 64)).IsFloat()
	errors = append(errors, latValidator.Errors()...)
	errors = append(errors, lngValidator.Errors()...)
	errors = append(errors, nameValidator.Errors()...)
	errors = append(errors, addressTitleValidator.Errors()...)
	if reqParams.CourierSpeed <= 0 {
		errors = append(errors, "CourierSpeed must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}

func logCreateRestaurantRequest(app *service.AppContainer, request *tcp.Request, restaurantId uint) error {
	userId := request.GetUserID()
	var payload map[string]any
	err := json.Unmarshal([]byte(request.Body), &payload)
	if err != nil && request.Body != "" {
		fmt.Printf(err.Error() + "\n")
	}
	log := action_log.ActionLog{
		UserID:     &userId,
		Action:     "Create Restaurant",
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
