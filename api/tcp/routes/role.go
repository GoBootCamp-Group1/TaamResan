package routes

import (
	"TaamResan/api/tcp/handlers/role_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/internal/role"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitRoleRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("POST /role", tcp_http_server.HandlerChain(
		role_handlers.CreateRole(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.ROLE, []uint{role.Admin}),
	))

	router.HandleFunc("PUT /role", tcp_http_server.HandlerChain(
		role_handlers.UpdateRole(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.ROLE, []uint{role.Admin}),
	))

	router.HandleFunc("DELETE /role/:id", tcp_http_server.HandlerChain(
		role_handlers.DeleteRole(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.ROLE, []uint{role.Admin}),
	))

	router.HandleFunc("GET /role/:id", tcp_http_server.HandlerChain(
		role_handlers.GetRole(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.ROLE, []uint{role.Admin}),
	))

	router.HandleFunc("GET /roles", tcp_http_server.HandlerChain(
		role_handlers.GetAllRoles(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.ROLE, []uint{role.Admin}),
	))
}
