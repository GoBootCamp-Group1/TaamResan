package address

import (
	"TaamResan/internal/address"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
	"strconv"
)

type RequestBody struct {
	Title string  `json:"title"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
}

func StoreAddress(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams RequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateInputs(conn, reqParams)

		//get authenticated user claims
		//userId := request.GetUserID()

		newAddress := address.Address{
			Title: reqParams.Title,
			Lat:   reqParams.Lat,
			Lng:   reqParams.Lng,
		}

		err = app.AddressService().CreateAddress(request.Context(), &newAddress)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "your address had been created successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateInputs(conn net.Conn, reqParams RequestBody) {
	var errors []string

	titleValidator := validator.Validate(reqParams.Title).MinLength(3)
	errors = append(errors, titleValidator.Errors()...)
	latValidator := validator.Validate(strconv.FormatFloat(reqParams.Lat, 'f', -1, 64)).IsFloat()
	errors = append(errors, latValidator.Errors()...)
	lngValidator := validator.Validate(strconv.FormatFloat(reqParams.Lng, 'f', -1, 64)).IsFloat()
	errors = append(errors, lngValidator.Errors()...)

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
