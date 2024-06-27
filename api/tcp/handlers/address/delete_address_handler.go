package address

import (
	"TaamResan/internal/address"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func DeleteAddress(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		//get address id from url params
		addressId, parseErr := strconv.ParseUint(request.UrlParams["addressId"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		deletedAddress := address.Address{
			ID: uint(addressId),
		}

		err := app.AddressService().DeleteAddress(request.Context(), &deletedAddress)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "your address had been deleted successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
