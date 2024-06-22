package wallet

import (
	"TaamResan/internal/wallet"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
	"strconv"
)

type TopUpWalletRequest struct {
	Amount     float64 `json:"amount"`
	CardNumber string  `json:"card_number"`
}

func TopUp(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		//parse input data
		var reqParams TopUpWalletRequest
		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}
		//validate date
		validateTopUpWalletInputs(conn, &reqParams)

		newWalletTopUp := wallet.WalletTopUp{
			Amount:     reqParams.Amount,
			CardNumber: reqParams.CardNumber,
		}

		err = app.WalletService().TopUp(request.Context(), &newWalletTopUp)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		//show response result
		responseBody := map[string]any{
			"message": "your card had been top-upped successfully",
		}

		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateTopUpWalletInputs(conn net.Conn, reqParams *TopUpWalletRequest) {
	var errors []string

	amount := strconv.FormatFloat(reqParams.Amount, 'f', -1, 64)
	amountValidator := validator.Validate(amount).IsFloat()
	errors = append(errors, amountValidator.Errors()...)

	cardNumberValidator := validator.Validate(reqParams.CardNumber).RegMatch(`^\d{16}$`)
	errors = append(errors, cardNumberValidator.Errors()...)

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
