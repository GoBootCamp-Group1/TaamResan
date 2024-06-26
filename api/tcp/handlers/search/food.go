package search

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"fmt"
	"net"
)

type searchResponse struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Lat   float64 `json:"lat"`
	Lng   float64 `json:"lng"`
}

func SearchFood(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {

		queryParams := request.QueryParams

		fmt.Println(queryParams)

		//foods, err := app.SearchService().SearchFoods(request.Context())
		//if err != nil {
		//	tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
		//	return
		//}
		//
		//responseBody := map[string]any{
		//	"data": map[string]any{"foods": getAddressesResponse(conn, foods)},
		//}
		//tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

//func getAddressesResponse(conn net.Conn, address []*address.Address) []GetAddressesResponse {
//	var result []GetAddressesResponse
//	marshalled, _ := json.Marshal(address)
//	if err := json.Unmarshal(marshalled, &result); err != nil {
//		tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
//		return result
//	}
//
//	return result
//}
