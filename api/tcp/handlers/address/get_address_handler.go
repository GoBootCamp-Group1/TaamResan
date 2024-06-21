package address

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

//type GetAddressResponse struct {
//	ID    int     `json:"id"`
//	Title string  `json:"title"`
//	Lat   float64 `json:"lat"`
//	Lng   float64 `json:"lng"`
//}

func GetAddress(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		//get address id from url params
		addressId, parseErr := strconv.ParseUint(request.UrlParams["addressId"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		//TODO: check for permission and ownership of resource

		address, err := app.AddressService().GetAddressByID(request.Context(), uint(addressId))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		//var result GetAddressResponse
		//marshalled, _ := json.Marshal(address)
		//if err := json.Unmarshal(marshalled, &result); err != nil {
		//	fmt.Printf("Can not unmarshal JSON, %v", err.Error())
		//}

		responseData := map[string]any{
			"address": address,
		}
		tcp.RespondJsonSuccess(conn, responseData)
		return
	}
}
