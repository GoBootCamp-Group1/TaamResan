package routes

import (
	"TaamResan/api/tcp/handlers/restaurant_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitRestaurantRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("POST /restaurants", tcp_http_server.HandlerChain(
		restaurant_handlers.CreateRestaurant(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("PUT /restaurants/:id", tcp_http_server.HandlerChain(
		restaurant_handlers.UpdateRestaurant(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("DELETE /restaurants/:id", tcp_http_server.HandlerChain(
		restaurant_handlers.DeleteRestaurant(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /restaurants/:id", tcp_http_server.HandlerChain(
		restaurant_handlers.GetRestaurant(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /restaurants", tcp_http_server.HandlerChain(
		restaurant_handlers.GetAllRestaurants(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("POST /restaurants/approve/:id", tcp_http_server.HandlerChain(
		restaurant_handlers.Approve(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("POST /restaurants/delegate", tcp_http_server.HandlerChain(
		restaurant_handlers.Delegate(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))
}
