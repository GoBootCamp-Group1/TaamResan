package address

import (
	"TaamResan/internal/address"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"encoding/json"
	"net"
)

type GetAddressesResponse struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
}

func GetAllAddresses(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {

		addresses, err := app.AddressService().GetAll(request.Context())
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"data": map[string]any{"addresses": getAddressesResponse(conn, addresses)},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func getAddressesResponse(conn net.Conn, address []*address.Address) []GetAddressesResponse {
	var result []GetAddressesResponse
	marshalled, _ := json.Marshal(address)
	if err := json.Unmarshal(marshalled, &result); err != nil {
		tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
		return result
	}

	return result
}
