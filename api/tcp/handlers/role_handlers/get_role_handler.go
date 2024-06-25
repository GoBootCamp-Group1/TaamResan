package role_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func GetRole(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		id, parseErr := strconv.ParseUint(request.UrlParams["id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		userId := request.GetUserID() // TODO: check permission
		if err := app.AccessService().CheckAdminAccess(request.Context(), userId); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		roleModel, err := app.RoleService().Get(request.Context(), uint(id))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "role loaded successfully",
			"data":    map[string]any{"role": roleModel},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
