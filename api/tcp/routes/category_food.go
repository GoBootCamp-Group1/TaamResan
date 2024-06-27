package routes

import (
	"TaamResan/api/tcp/handlers/category_food_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/internal/role"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitCategoryFoodRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("POST /category-foods", tcp_http_server.HandlerChain(
		category_food_handlers.Create(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.CATEGORY, []uint{role.RestaurantOwner, role.RestaurantOperator}),
	))

	router.HandleFunc("DELETE /category-foods/:category_food_id", tcp_http_server.HandlerChain(
		category_food_handlers.Delete(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.CATEGORY, []uint{role.RestaurantOwner, role.RestaurantOperator}),
	))

	router.HandleFunc("GET /category-foods/:category_food_id", tcp_http_server.HandlerChain(
		category_food_handlers.Get(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.CATEGORY, []uint{role.RestaurantOwner, role.RestaurantOperator}),
	))

	router.HandleFunc("GET /foods/:food_id/category-foods", tcp_http_server.HandlerChain(
		category_food_handlers.GetAllByFood(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.CATEGORY, []uint{role.RestaurantOwner, role.RestaurantOperator}),
	))

	router.HandleFunc("GET /categories/:category_id/category-foods", tcp_http_server.HandlerChain(
		category_food_handlers.GetAllByCategory(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
		middlewares.PermissionCheck(app, role.CATEGORY, []uint{role.RestaurantOwner, role.RestaurantOperator}),
	))
}
