package role_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
)

func GetAllRoles(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		//userId := request.GetUserID() // TODO: check that user has permission and is ADMIN to do this

		roleModels, err := app.RoleService().GetAll(request.Context())
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseData := map[string]any{
			"message": "roles loaded successfully",
			"role":    &roleModels,
		}
		tcp.RespondJsonSuccess(conn, responseData)
		return
	}
}
