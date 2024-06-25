package routes

import (
	"TaamResan/api/tcp/handlers/block_restaurant_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitBlockRestaurantRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("POST /block_restaurants", tcp_http_server.HandlerChain(
		block_restaurant_handlers.Create(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("DELETE /block_restaurants/:block_restaurant_id", tcp_http_server.HandlerChain(
		block_restaurant_handlers.Delete(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /users/:user_id/block_restaurants", tcp_http_server.HandlerChain(
		block_restaurant_handlers.GetAllByRestaurant(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))
}
