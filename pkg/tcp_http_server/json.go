package tcp_http_server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

// RespondJson writes a JSON response to the client.
func RespondJson(conn net.Conn, data any, statusCode int) {
	responseBody, err := json.Marshal(data)
	if err != nil {
		RespondJsonError(conn, "Internal Server Error", 500)
		return
	}

	response := fmt.Sprintf("HTTP/1.1 %d %s\r\n", statusCode, StatusText(statusCode))
	response += "Content-Type: application/json\r\n"
	response += fmt.Sprintf("Content-Length: %d\r\n", len(responseBody))
	response += "\r\n"
	response += string(responseBody)

	conn.Write([]byte(response))
}

// RespondJsonError writes an error response to the client.
func RespondJsonError(conn net.Conn, message string, statusCode int) {

	if statusCode == 0 {
		statusCode = 500
	}

	responseBody := map[string]string{"error": message}
	RespondJson(conn, responseBody, statusCode)
}

// RespondJsonValidateError writes validation errors response to the client.
func RespondJsonValidateError(conn net.Conn, messages []string, statusCode int) {

	if statusCode == 0 {
		statusCode = 422
	}

	responseBody := map[string][]string{"error": messages}
	RespondJson(conn, responseBody, statusCode)
}

// RespondJsonSuccess writes an error response to the client.
func RespondJsonSuccess(conn net.Conn, data any, statusCode int) {

	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	responseBody := map[string]any{"data": data}
	RespondJson(conn, responseBody, statusCode)
}
