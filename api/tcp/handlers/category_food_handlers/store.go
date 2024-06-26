package category_food_handlers

import (
	"TaamResan/internal/category_food"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
)

type CreateRequestBody struct {
	CategoryId uint `json:"category_id"`
	FoodId     uint `json:"food_id"`
}

func Create(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		var reqParams CreateRequestBody

		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		validateCreateInputs(conn, reqParams)

		userId := request.GetUserID() // TODO: check permission
		cat, err := app.CategoryService().GetById(request.Context(), reqParams.CategoryId)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}
		restaurantId := cat.RestaurantId
		if err = app.AccessService().CheckRestaurantStaff(request.Context(), userId, restaurantId); err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.FORBIDDEN)
			return
		}

		newCategoryFood := category_food.CategoryFood{
			CategoryId: reqParams.CategoryId,
			FoodId:     reqParams.FoodId,
		}

		id, err := app.CategoryFoodService().Create(request.Context(), &newCategoryFood)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"data":    map[string]any{"category_food_id": id},
			"message": "category_food created successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateCreateInputs(conn net.Conn, reqParams CreateRequestBody) {
	var errors []string

	if reqParams.CategoryId <= 0 {
		errors = append(errors, "CategoryId must be greater than 0")
	}

	if reqParams.FoodId <= 0 {
		errors = append(errors, "FoodId must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
