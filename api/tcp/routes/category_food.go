package routes

import (
	"TaamResan/api/tcp/handlers/category_food_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitCategoryFoodRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("POST /category-foods", tcp_http_server.HandlerChain(
		category_food_handlers.Create(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("PUT /category-foods/:category_food_id", tcp_http_server.HandlerChain(
		category_food_handlers.Update(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("DELETE /category-foods/:category_food_id", tcp_http_server.HandlerChain(
		category_food_handlers.Delete(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /category-foods/:category_food_id", tcp_http_server.HandlerChain(
		category_food_handlers.Get(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /restaurants/:restaurant_id/category-foods", tcp_http_server.HandlerChain(
		category_food_handlers.GetAll(app),
		middlewares.LoggingMiddleware,
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))
}
