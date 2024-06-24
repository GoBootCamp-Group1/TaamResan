package food_handlers

import (
	"TaamResan/internal/category"
	"TaamResan/internal/food"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
)

type CreateRequestBody struct {
	RestaurantId       uint     `json:"restaurant_id"`
	Name               string   `json:"name"`
	Price              float64  `json:"price"`
	CancelRate         float64  `json:"cancel_rate"`
	PreparationMinutes uint     `json:"preparation_minutes"`
	Categories         []string `json:"categories,omitempty"`
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

		userId := request.GetUserID() // TODO: check permission and is OWNER to do this

		categories := make([]*category.Category, 0, len(reqParams.Categories))
		if len(reqParams.Categories) > 0 {
			for _, c := range reqParams.Categories {
				categories = append(categories, &category.Category{Name: c})
			}
		}

		newFood := food.Food{
			RestaurantId:       reqParams.RestaurantId,
			CreatedBy:          userId,
			Name:               reqParams.Name,
			Price:              reqParams.Price,
			CancelRate:         reqParams.CancelRate,
			PreparationMinutes: reqParams.PreparationMinutes,
			Categories:         categories,
		}

		id, err := app.FoodService().Create(request.Context(), &newFood)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		responseBody := map[string]any{
			"data":    map[string]any{"food_id": id},
			"message": "food created successfully",
		}
		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateCreateInputs(conn net.Conn, reqParams CreateRequestBody) {
	var errors []string

	if reqParams.RestaurantId <= 0 {
		errors = append(errors, "RestaurantId must be greater than 0")
	}

	nameValidator := validator.Validate(reqParams.Name).MinLength(3)
	errors = append(errors, nameValidator.Errors()...)

	if reqParams.Price <= 0 {
		errors = append(errors, "Price must be greater than 0")
	}

	if reqParams.CancelRate < 0 {
		errors = append(errors, "CancelRate must be equal or greater than 0")
	}

	if reqParams.PreparationMinutes <= 0 {
		errors = append(errors, "PreparationMinutes must be greater than 0")
	}

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
