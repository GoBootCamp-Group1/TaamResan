package search

import (
	"TaamResan/internal/restaurant"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func SearchRestaurant(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {

		queryParams := request.QueryParams

		if (queryParams["lat"] != "" && queryParams["lng"] == "") || (queryParams["lat"] == "" && queryParams["lng"] != "") {
			tcp.RespondJsonError(conn, "invalid location, need both valid lat and lng", tcp.INVALID_INPUT)
			return
		}

		var categoryID *uint
		var lat, lng *float64

		if queryParams["category_id"] != "" {
			id, err := strconv.ParseUint(queryParams["category_id"], 10, 64)
			if err == nil {
				castID := uint(id)
				categoryID = &castID
			} else {
				tcp.RespondJsonError(conn, "invalid category", tcp.INVALID_INPUT)
				return
			}
		}

		if queryParams["lat"] != "" {
			latitude, err := strconv.ParseFloat(queryParams["lat"], 64)
			if err == nil {
				lat = &latitude
			} else {
				tcp.RespondJsonError(conn, "invalid latitude", tcp.INVALID_INPUT)
				return
			}
		}

		if queryParams["lng"] != "" {
			longitude, err := strconv.ParseFloat(queryParams["lng"], 64)
			if err == nil {
				lng = &longitude
			} else {
				tcp.RespondJsonError(conn, "invalid longitude", tcp.INVALID_INPUT)
				return
			}
		}

		searchData := restaurant.RestaurantSearch{
			Name:       queryParams["name"],
			CategoryID: categoryID,
			Lat:        lat,
			Lng:        lng,
		}

		restaurants, err := app.SearchService().SearchRestaurants(request.Context(), &searchData)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"data": map[string]any{"restaurants": &restaurants},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
