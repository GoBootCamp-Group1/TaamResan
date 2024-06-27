package middlewares

import (
	"TaamResan/internal/role"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"errors"
	"net"
)

var (
	NotAllowed = errors.New("not allowed")
)

func PermissionCheck(app *service.AppContainer, permission string, roles []uint) tcp.MiddlewareFunc {
	return func(next tcp.HandlerFunc) tcp.HandlerFunc {
		return func(conn net.Conn, request *tcp.Request) {
			userId := request.GetUserID()

			// get roles
			roleIds, err := app.UserService().GetUserRoles(request.Context(), userId)
			if err != nil {
				tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
				return
			}

			// check roles
			haveRole := false
			for _, r1 := range roleIds {
				for _, r2 := range roles {
					if r1 == r2 {
						haveRole = true
					}
				}
			}

			if !haveRole {
				tcp.RespondJsonError(conn, NotAllowed.Error(), tcp.FORBIDDEN)
				return
			}

			// check permissions
			var permissions []string
			for _, roleId := range roleIds {
				permissions = append(permissions, role.RolePermissions[roleId]...)
			}

			havePermission := false
			for _, p := range permissions {
				if permission == p {
					havePermission = true
				}
			}

			if !havePermission {
				tcp.RespondJsonError(conn, NotAllowed.Error(), tcp.FORBIDDEN)
				return
			}

			next(conn, request)
		}
	}
}
