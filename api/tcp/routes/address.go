package routes

import (
	"TaamResan/api/tcp/handlers/address"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitAddressRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("POST /addresses", tcp_http_server.HandlerChain(
		address.StoreAddress(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("PUT /addresses/:addressId", tcp_http_server.HandlerChain(
		address.UpdateAddress(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /addresses", tcp_http_server.HandlerChain(
		address.GetAllAddresses(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /addresses/:addressId", tcp_http_server.HandlerChain(
		address.GetAddress(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("DELETE /addresses/:addressId", tcp_http_server.HandlerChain(
		address.DeleteAddress(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))
}
