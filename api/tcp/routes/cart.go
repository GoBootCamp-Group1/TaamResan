package routes

import (
	"TaamResan/api/tcp/handlers/cart_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitCartRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("DELETE /carts/:cart_id", tcp_http_server.HandlerChain(
		cart_handlers.Delete(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /users/:user_id/carts", tcp_http_server.HandlerChain(
		cart_handlers.GetAllByUser(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))
}
