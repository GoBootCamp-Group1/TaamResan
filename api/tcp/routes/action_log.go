package routes

import (
	"TaamResan/api/tcp/handlers/action_log_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/internal/role"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitActionLogRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("GET /users/:user_id/logs", tcp_http_server.HandlerChain(
		action_log_handlers.GetByUser(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.LOG, []uint{role.Admin}),
	))

	router.HandleFunc("GET /restaurants/:restaurant_id/logs", tcp_http_server.HandlerChain(
		action_log_handlers.GetByRestaurant(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.LOG, []uint{role.Admin}),
	))
}
