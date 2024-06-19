package tcp

import (
	"TaamResan/api/tcp/handlers"
	"TaamResan/api/tcp/handlers/authentication_handlers"
	"TaamResan/api/tcp/handlers/signup_handlers"
	"TaamResan/api/tcp/middlewares"
	"TaamResan/cmd/api/config"
	"TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"fmt"
	"net"
)

func Run(cfg config.Server, app *service.AppContainer) {
	// Define listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting TCP server:", err)
		return
	}
	defer listener.Close()

	// Define routes
	router := tcp_http_server.NewRouter()

	// register global routes
	registerGlobalRoutes(router, app)

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

func registerGlobalRoutes(router *tcp_http_server.Router, app *service.AppContainer) {
	router.HandleFunc("GET /", tcp_http_server.HandlerChain(
		handlers.HomeHandler,
		middlewares.LoggingMiddleware,
	))
	router.HandleFunc("POST /todo", tcp_http_server.HandlerChain(
		handlers.TodoHandler,
		middlewares.LoggingMiddleware,
		//middlewares.AuthMiddleware,
	))

	router.HandleFunc("POST /signup", tcp_http_server.HandlerChain(
		signup_handlers.SignUp(app),
		middlewares.LoggingMiddleware,
	))

	router.HandleFunc("POST /login", tcp_http_server.HandlerChain(
		authentication_handlers.Login(app),
		middlewares.LoggingMiddleware,
	))

}
