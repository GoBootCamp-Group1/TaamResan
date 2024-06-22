package cards

import (
	"TaamResan/internal/wallet"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
)

type StoreWalletCardRequest struct {
	Title    string `json:"title"`
	BankName string `json:"bank_name"`
	Number   string `json:"number"`
}

func StoreWalletCard(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		//parse input data
		var reqParams StoreWalletCardRequest
		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}
		//validate date
		validateStoreWalletInputs(conn, &reqParams)

		newCard := wallet.WalletCard{
			Title:    reqParams.Title,
			BankName: reqParams.BankName,
			Number:   reqParams.Number,
		}

		//store in database using service
		err = app.WalletService().CreateWalletCard(request.Context(), &newCard)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		//show response result
		responseBody := map[string]any{
			"message": "your address had been created successfully",
		}

		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateStoreWalletInputs(conn net.Conn, reqParams *StoreWalletCardRequest) {
	var errors []string

	bankNameValidator := validator.Validate(reqParams.BankName).MinLength(3)
	errors = append(errors, bankNameValidator.Errors()...)
	titleValidator := validator.Validate(reqParams.Title).MinLength(3)
	errors = append(errors, titleValidator.Errors()...)
	numberValidator := validator.Validate(reqParams.Number).RegMatch(`^d{16}$`)
	errors = append(errors, numberValidator.Errors()...)

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
