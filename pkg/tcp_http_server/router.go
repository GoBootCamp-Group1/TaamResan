package tcp_http_server

import (
	"TaamResan/pkg/jwt"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type Request struct {
	Method      string
	Uri         string
	Body        string
	Headers     map[string]string
	UrlParams   map[string]string
	QueryParams map[string]string
	ctx         context.Context
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

func (r *Request) ExtractBodyParamsInto(mock any) error {
	return json.Unmarshal([]byte(r.Body), &mock)
}

func (r *Request) GetClaims() *jwt.UserClaims {
	return r.Context().Value(jwt.UserClaimKey).(*jwt.UserClaims)
}

func (r *Request) GetUserID() uint {
	return r.Context().Value(jwt.UserClaimKey).(*jwt.UserClaims).UserID
}

func (r *Request) fillQueryParams() {
	r.QueryParams = make(map[string]string)
	parsedUrl, err := url.Parse(r.Uri)
	if err != nil {
		return
	}
	for key, values := range parsedUrl.Query() {
		if len(values) > 0 {
			r.QueryParams[key] = values[0]
		}
	}
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
	pattern *regexp.Regexp
	method  string
	keys    []string
	handler HandlerFunc
}

// Router holds the route mappings and middleware.
type Router struct {
	routes []route
}

// NewRouter creates a new Router instance.
func NewRouter() *Router {
	return &Router{
		routes: []route{},
	}
}

// HandleFunc registers a handler with middlewares for the given method and pattern.
func (r *Router) HandleFunc(methodAndPattern string, handler HandlerFunc) {
	methodAndPatternParts := strings.SplitN(methodAndPattern, " ", 2)
	if len(methodAndPatternParts) != 2 {
		panic("Invalid method and pattern")
	}
	method := methodAndPatternParts[0]
	pattern := methodAndPatternParts[1]
	regex, keys := patternToRegex(pattern)
	r.routes = append(r.routes, route{pattern: regex, method: method, keys: keys, handler: handler})
}

func patternToRegex(pattern string) (*regexp.Regexp, []string) {
	var keys []string
	regexPattern := "^" + regexp.QuoteMeta(pattern)
	regexPattern = strings.ReplaceAll(regexPattern, `\\:`, `:`)
	regexPattern = regexp.MustCompile(`:[^/]+`).ReplaceAllStringFunc(regexPattern, func(param string) string {
		keys = append(keys, param[1:])
		return `([^/]+)`
	})
	regexPattern += `$`
	return regexp.MustCompile(regexPattern), keys
}

// Serve handles incoming connections and routes them to the appropriate handler.
func (r *Router) Serve(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	requestLine, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	requestLine = strings.TrimSpace(requestLine)

	// Parse the request line
	parts := strings.Split(requestLine, " ")
	if len(parts) < 3 {
		fmt.Println("Invalid request line")
		return
	}
	method := parts[0]
	uri := parts[1]

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

	// Handle the request body if present ONLY for POST and PUT
	var body string

	fmt.Println(body)
	fmt.Println(method)

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
		}
	}

	// Route the request
	for _, route := range r.routes {
		if route.pattern.MatchString(uri) && route.method == method {
			matches := route.pattern.FindStringSubmatch(uri)
			urlParams := make(map[string]string)
			for i, match := range matches[1:] {
				urlParams[route.keys[i]] = match
			}

			request := &Request{
				Method:    method,
				Uri:       uri,
				Body:      body,
				Headers:   headers,
				UrlParams: urlParams,
				ctx:       context.Background(),
			}
			request.fillQueryParams()
			route.handler(conn, request)
			return
		}
	}

	HttpNotFound(conn)
}

// HttpNotFound writes a 404 Not Found response.
func HttpNotFound(conn net.Conn) {
	RespondJsonError(conn, "Not Found", NOT_FOUND)
}
