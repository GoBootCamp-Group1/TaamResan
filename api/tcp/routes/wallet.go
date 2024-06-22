package routes

import (
	"TaamResan/api/tcp/handlers/wallet/cards"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitWalletRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	//store card for user wallet
	router.HandleFunc("POST /wallet/cards", tcp_http_server.HandlerChain(
		cards.StoreWalletCard(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))
}
