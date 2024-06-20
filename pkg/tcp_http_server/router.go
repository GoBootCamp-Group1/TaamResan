package tcp_http_server

import (
	"TaamResan/pkg/jwt"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	Method  string
	Uri     string
	Body    string
	Headers map[string]string
	ctx     context.Context
}

func (r *Request) Context() context.Context {
	if r.ctx != nil {
		return r.ctx
	}
	return context.Background()
}

func (r *Request) WithContext(ctx context.Context) *Request {
	if ctx == nil {
		panic("nil context")
	}
	r2 := new(Request)
	*r2 = *r
	r2.ctx = ctx
	return r2
}

func (r *Request) ExtractParamsInto(mock any) error {
	return json.Unmarshal([]byte(r.Body), &mock)
}

func (r *Request) GetClaims() *jwt.UserClaims {
	return r.Context().Value(jwt.UserClaimKey).(*jwt.UserClaims)
}

func (r *Request) GetUserID() uint {
	return r.Context().Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID
}

// HandlerFunc is the type for HTTP request handlers.
type HandlerFunc func(conn net.Conn, request *Request)

// MiddlewareFunc is the type for middleware functions.
type MiddlewareFunc func(HandlerFunc) HandlerFunc

// HandlerChain Utility function to apply middleware to a handler.
func HandlerChain(handler HandlerFunc, middlewares ...MiddlewareFunc) HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

type route struct {
	handler HandlerFunc
}

// Router holds the route mappings and middleware.
type Router struct {
	routes map[string]route
}

// NewRouter creates a new Router instance.
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]route),
	}
}

// HandleFunc registers a handler with middlewares for the given method and pattern.
func (r *Router) HandleFunc(methodAndPattern string, handler HandlerFunc) {
	r.routes[methodAndPattern] = route{handler: handler}
}

// Serve handles incoming connections and routes them to the appropriate handler.
func (r *Router) Serve(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	requestLine, err := reader.ReadString('\n')
	if err != nil {
		//fmt.Println("Error reading request line:", err)
		return
	}
	requestLine = strings.TrimSpace(requestLine)
	//fmt.Println("Request Line:", requestLine)

	// Parse the request line
	parts := strings.Split(requestLine, " ")
	if len(parts) < 3 {
		fmt.Println("Invalid request line")
		return
	}
	method := parts[0]
	uri := parts[1]
	//protocol := parts[2]

	//fmt.Println("Method:", method)
	//fmt.Println("URI:", uri)
	//fmt.Println("Protocol:", protocol)

	// Read headers
	headers := make(map[string]string)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading header line:", err)
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		headerParts := strings.SplitN(line, ": ", 2)
		if len(headerParts) == 2 {
			headers[headerParts[0]] = headerParts[1]
		}
	}
	//fmt.Println("Headers:", headers)

	// Handle the request body if present ONLY for POST and PUT
	var body string
	if method == "POST" || method == "PUT" {
		contentLength := headers["Content-Length"]
		if contentLength != "" {
			bodyLength, err := strconv.Atoi(contentLength)
			if err != nil {
				fmt.Println("Invalid Content-Length:", err)
				return
			}
			bodyBytes := make([]byte, bodyLength)
			_, err = reader.Read(bodyBytes)
			if err != nil {
				fmt.Println("Error reading body:", err)
				return
			}
			body = string(bodyBytes)
			//fmt.Println("Body:", body)
		}
	}

	// Route the request
	uriWithoutQueryParams := strings.Split(uri, "?")[0]
	key := method + " " + uriWithoutQueryParams

	if route, ok := r.routes[key]; ok {

		//sending parsed request
		request := Request{
			Method:  method,
			Uri:     uri,
			Body:    body,
			Headers: headers,
			ctx:     context.Background(),
		}

		route.handler(conn, &request)
	} else {
		HttpNotFound(conn)
	}
}

// HttpNotFound writes a 404 Not Found response.
func HttpNotFound(conn net.Conn) {
	RespondJsonError(conn, "Not Found", NOT_FOUND)
}
