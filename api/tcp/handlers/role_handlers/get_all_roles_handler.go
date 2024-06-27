package role_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
)

func GetAllRoles(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		roleModels, err := app.RoleService().GetAll(request.Context())
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "roles loaded successfully",
			"data":    map[string]any{"roles": &roleModels},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
