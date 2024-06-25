package category_food_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func GetAllByFood(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		foodId, parseErr := strconv.ParseUint(request.UrlParams["food_id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		userId := request.GetUserID() // TODO: check permission
		f, err := app.FoodService().GetById(request.Context(), uint(foodId))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}
		restaurantId := f.RestaurantId
		if err = app.AccessService().CheckRestaurantOwner(request.Context(), userId, restaurantId); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		categoryFoodModels, err := app.CategoryFoodService().GetAllByFoodId(request.Context(), uint(foodId))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"message": "category_foods loaded successfully",
			"data":    map[string]any{"category_foods": &categoryFoodModels},
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
