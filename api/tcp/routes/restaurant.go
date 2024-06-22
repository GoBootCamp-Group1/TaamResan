package routes

import (
	"TaamResan/api/tcp/handlers/restaurant_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitRestaurantRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("POST /restaurant", tcp_http_server.HandlerChain(
		restaurant_handlers.CreateRestaurant(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("PUT /restaurant", tcp_http_server.HandlerChain(
		restaurant_handlers.UpdateRestaurant(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("DELETE /restaurant/:id", tcp_http_server.HandlerChain(
		restaurant_handlers.DeleteRestaurant(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /restaurant/:id", tcp_http_server.HandlerChain(
		restaurant_handlers.GetRestaurant(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /restaurants", tcp_http_server.HandlerChain(
		restaurant_handlers.GetAllRestaurants(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("POST /restaurant/approve/:id", tcp_http_server.HandlerChain(
		restaurant_handlers.Approve(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("POST /restaurant/delegate", tcp_http_server.HandlerChain(
		restaurant_handlers.Delegate(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))
}
