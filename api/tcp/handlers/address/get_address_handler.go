package address

import (
	"TaamResan/internal/address"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"encoding/json"
	"net"
	"strconv"
)

type GetAddressResponse struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
}

func GetAddress(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		//get address id from url params
		addressId, parseErr := strconv.ParseUint(request.UrlParams["addressId"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		//TODO: check for permission and ownership of resource

		fetchedAddress, err := app.AddressService().GetAddressByID(request.Context(), uint(addressId))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"data": map[string]any{"address": getAddressResponse(conn, fetchedAddress)},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func getAddressResponse(conn net.Conn, address *address.Address) GetAddressResponse {
	var result GetAddressResponse
	marshalled, _ := json.Marshal(address)
	if err := json.Unmarshal(marshalled, &result); err != nil {
		tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
		return result
	}

	return result
}
