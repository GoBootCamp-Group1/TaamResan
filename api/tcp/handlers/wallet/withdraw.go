package wallet

import (
	"TaamResan/internal/wallet"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/pkg/validator"
	"TaamResan/service"
	"net"
	"strconv"
)

type WithdrawWalletRequest struct {
	Amount float64 `json:"amount"`
	CardID uint    `json:"card_id"`
}

func Withdraw(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {
		//parse input data
		var reqParams WithdrawWalletRequest
		err := request.ExtractBodyParamsInto(&reqParams)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}
		//validate date
		validateWithdrawWalletInputs(conn, &reqParams)

		newWalletWithdraw := wallet.WalletWithdraw{
			Amount: reqParams.Amount,
			CardID: reqParams.CardID,
		}

		err = app.WalletService().Withdraw(request.Context(), &newWalletWithdraw)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		//show response result
		responseBody := map[string]any{
			"message": "successfully withdraw",
		}

		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}

func validateWithdrawWalletInputs(conn net.Conn, reqParams *WithdrawWalletRequest) {
	var errors []string

	amount := strconv.FormatFloat(reqParams.Amount, 'f', -1, 64)
	amountValidator := validator.Validate(amount).IsFloat()
	errors = append(errors, amountValidator.Errors()...)

	cardNumberValidator := validator.Validate(strconv.Itoa(int(reqParams.CardID))).NonEmpty()
	errors = append(errors, cardNumberValidator.Errors()...)

	if len(errors) > 0 {
		tcp.RespondJsonValidateError(conn, errors, tcp.INVALID_INPUT)
		return
	}
}
