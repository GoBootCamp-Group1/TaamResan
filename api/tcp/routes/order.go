package routes

import (
	"TaamResan/api/tcp/handlers/order"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/internal/role"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitOrderRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("POST /orders", tcp_http_server.HandlerChain(
		order.StoreHandler(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.ORDER, []uint{role.Customer}),
	))

	router.HandleFunc("PUT /orders/:orderId/cancel", tcp_http_server.HandlerChain(
		order.CancelHandler(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /orders/:orderId", tcp_http_server.HandlerChain(
		order.InfoHandler(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("PUT /orders/:orderId/approve", tcp_http_server.HandlerChain(
		order.CustomerApproveHandler(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))
}
