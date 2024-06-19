package tcp_http_server

// StatusText returns the text for the HTTP status code.
func StatusText(statusCode int) string {
	switch statusCode {
	case 200:
		return "OK"
	case 400:
		return "Bad Request"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	default:
		return "Unknown Status"
	}
}

const (
	OK                    = 200
	BAD_REQUEST           = 400
	UNAUTHORIZED          = 401
	INVALID_INPUT         = 422
	NOT_FOUND             = 404
	INTERNAL_SERVER_ERROR = 500
)
