package address

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
)

func GetAllAddresses(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {

		addresses, err := app.AddressService().GetAll(request.Context())
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseData := map[string]any{
			"data": addresses,
		}
		tcp.RespondJsonSuccess(conn, responseData)
		return
	}
}
