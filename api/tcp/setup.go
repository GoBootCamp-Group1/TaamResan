package tcp

import (
	"TaamResan/api/tcp/handlers"
	"TaamResan/api/tcp/handlers/authentication_handlers"
	"TaamResan/api/tcp/handlers/signup_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/api/tcp/routes"
	"TaamResan/cmd/api/config"
	"TaamResan/internal/adapters/storage"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
)

func Run(cfg config.Server, app *service.AppContainer) {
	if err := app.RoleService().InitializeRoles(context.Background()); err != nil {
		if !errors.Is(err, storage.ErrRoleExists) {
			log.Fatalf("Error initializing roles: %v", err)
		}
	}

	// Define listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.HttpPort))
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	// Define routes
	router := tcp_http_server.NewRouter()

	// register global routes
	registerGlobalRoutes(router, app, cfg)

	// registering users APIs
	//registerUsersAPI(api, app.UserService(), []byte(cfg.TokenSecret))

	fmt.Printf("🌏 Listening on %s:%d\n", cfg.Host, cfg.HttpPort)

	// run server
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go router.Serve(conn)
	}
}

func registerGlobalRoutes(router *tcp_http_server.Router, app *service.AppContainer, cfg config.Server) {
	router.HandleFunc("GET /", tcp_http_server.HandlerChain(
		handlers.HomeHandler,
		middlewares.LoggingMiddleware(app.ActionLogService()),
	))
	router.HandleFunc("POST /todo", tcp_http_server.HandlerChain(
		handlers.TodoHandler,
		middlewares.LoggingMiddleware(app.ActionLogService()),
		middlewares.AuthMiddleware(cfg.TokenSecret),
	))

	router.HandleFunc("POST /signup", tcp_http_server.HandlerChain(
		signup_handlers.SignUp(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
	))

	router.HandleFunc("POST /login", tcp_http_server.HandlerChain(
		authentication_handlers.Login(app),
		middlewares.LoggingMiddleware(app.ActionLogService()),
	))

	routes.InitUserRoutes(router, app, cfg)
	routes.InitAddressRoutes(router, app, cfg)
	routes.InitRoleRoutes(router, app, cfg)
	routes.InitWalletRoutes(router, app, cfg)
	routes.InitRestaurantRoutes(router, app, cfg)
	routes.InitRestaurantStaffRoutes(router, app, cfg)
	routes.InitCategoryRoutes(router, app, cfg)
	routes.InitFoodRoutes(router, app, cfg)
	routes.InitCategoryFoodRoutes(router, app, cfg)
	routes.InitCartRoutes(router, app, cfg)
	routes.InitCartItemRoutes(router, app, cfg)
}
