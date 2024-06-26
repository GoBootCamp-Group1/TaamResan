package routes

import (
	"TaamResan/api/tcp/handlers/search"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitSearchRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("GET /search/food", tcp_http_server.HandlerChain(
		search.SearchFood(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	//router.HandleFunc("GET /search/restaurant", tcp_http_server.HandlerChain(
	//	search.SearchRestaurant(app),
	//	middlewares.LoggingMiddleware(app.ActionLogService()),
	//	middlewares.AuthMiddleware(cfg.TokenSecret),
	//))
}
