package restaurant_handlers

import (
	"TaamResan/internal/address"
	"TaamResan/internal/restaurant"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
	"strconv"
)

type UpdateRequestBody struct {
	Name         string  `json:"name"`
	AddressTitle string  `json:"address_title"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	CourierSpeed float64 `json:"courier_speed"`
}

func UpdateRestaurant(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		id, parseErr := strconv.ParseUint(request.UrlParams["id"], 10, 64)

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

		userId := request.GetUserID() // TODO: check permission
		if err = app.AccessService().CheckRestaurantOwner(request.Context(), userId, uint(id)); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		newRestaurant := restaurant.Restaurant{
			ID:      uint(id),
			Name:    reqParams.Name,
			OwnedBy: userId,
			Address: address.Address{
				Title: reqParams.AddressTitle,
				Lat:   reqParams.Lat,
				Lng:   reqParams.Lng,
			},
			CourierSpeed: reqParams.CourierSpeed,
		}

		err = app.RestaurantService().Update(request.Context(), &newRestaurant)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "restaurant updated successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateUpdateInputs(conn net.Conn, reqParams UpdateRequestBody) {
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
