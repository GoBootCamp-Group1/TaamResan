package category_food_handlers

import (
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func GetAllByCategory(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		categoryId, parseErr := strconv.ParseUint(request.UrlParams["category_id"], 10, 64)

		if parseErr != nil {
			tcp.RespondJsonError(conn, parseErr.Error(), tcp.NOT_FOUND)
			return
		}

		userId := request.GetUserID() // TODO: check permission
		cat, err := app.CategoryService().GetById(request.Context(), uint(categoryId))
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}
		restaurantId := cat.RestaurantId
		if err = app.AccessService().CheckRestaurantOwner(request.Context(), userId, restaurantId); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		categoryFoodModels, err := app.CategoryFoodService().GetAllByCategoryId(request.Context(), uint(categoryId))
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
