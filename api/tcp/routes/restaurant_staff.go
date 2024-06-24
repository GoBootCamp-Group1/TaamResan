package routes

import (
	"TaamResan/api/tcp/handlers/restaurant_staff_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
)

func InitRestaurantStaffRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("POST /restaurant-staff", tcp_http_server.HandlerChain(
		restaurant_staff_handlers.Create(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("DELETE /restaurant-staff/:id", tcp_http_server.HandlerChain(
		restaurant_staff_handlers.Delete(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("GET /restaurant-staff/:restaurant_id", tcp_http_server.HandlerChain(
		restaurant_staff_handlers.GetAllByRestaurant(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))
}
