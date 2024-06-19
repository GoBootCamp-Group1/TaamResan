package tcp_http_server

import (
	"encoding/json"
	"fmt"
	"net"
)

// RespondJson writes a JSON response to the client.
func RespondJson(conn net.Conn, data any, statusCode int) {
	responseBody, err := json.Marshal(data)
	if err != nil {
		RespondJsonError(conn, "Internal Server Error", INTERNAL_SERVER_ERROR)
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
func RespondJsonValidateError(conn net.Conn, errors []string, statusCode int) {

	if statusCode == 0 {
		statusCode = 422
	}

	responseBody := map[string][]string{"errors": errors}
	RespondJson(conn, responseBody, statusCode)
}

// RespondJsonSuccess writes an error response to the client.
func RespondJsonSuccess(conn net.Conn, data any) {
	responseBody := map[string]any{"data": data}
	RespondJson(conn, responseBody, OK)
}

// RespondJsonPaginate writes an error response to the client.
func RespondJsonPaginate(conn net.Conn, data any) {
	//TODO: implement
}
