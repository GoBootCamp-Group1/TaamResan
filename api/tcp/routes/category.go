package routes

import (
	"TaamResan/api/tcp/handlers/category_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/internal/role"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitCategoryRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("POST /categories", tcp_http_server.HandlerChain(
		category_handlers.Create(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.CATEGORY, []uint{role.RestaurantOwner, role.RestaurantOperator}),
	))

	router.HandleFunc("PUT /categories/:category_id", tcp_http_server.HandlerChain(
		category_handlers.Update(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.CATEGORY, []uint{role.RestaurantOwner, role.RestaurantOperator}),
	))

	router.HandleFunc("DELETE /categories/:category_id", tcp_http_server.HandlerChain(
		category_handlers.Delete(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.CATEGORY, []uint{role.RestaurantOwner, role.RestaurantOperator}),
	))

	router.HandleFunc("GET /categories/:category_id", tcp_http_server.HandlerChain(
		category_handlers.Get(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.CATEGORY, []uint{role.RestaurantOwner, role.RestaurantOperator}),
	))

	router.HandleFunc("GET /restaurants/:restaurant_id/categories", tcp_http_server.HandlerChain(
		category_handlers.GetAll(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.CATEGORY, []uint{role.RestaurantOwner, role.RestaurantOperator}),
	))

}
