package routes

import (
	"TaamResan/api/tcp/handlers/category_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitCategoryRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("POST /category", tcp_http_server.HandlerChain(
		category_handlers.Create(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("PUT /category", tcp_http_server.HandlerChain(
		category_handlers.Update(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("DELETE /category/:id", tcp_http_server.HandlerChain(
		category_handlers.Delete(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /category/:id", tcp_http_server.HandlerChain(
		category_handlers.Get(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /categories/:restaurant_id", tcp_http_server.HandlerChain(
		category_handlers.GetAll(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

}
