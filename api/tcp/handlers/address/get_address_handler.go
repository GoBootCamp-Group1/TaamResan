package address

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

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

		responseData := map[string]any{
			"data": address,
		}
		tcp.RespondJsonSuccess(conn, responseData)
		return
	}
}
