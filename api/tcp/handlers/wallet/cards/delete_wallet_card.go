package cards

import (
	"TaamResan/internal/wallet"
	tcp "TaamResan/pkg/tcp_http_server"
	"TaamResan/service"
	"net"
	"strconv"
)

func DeleteWalletCard(app *service.AppContainer) tcp.HandlerFunc {
	return func(conn net.Conn, request *tcp.Request) {

		//get card id
		if request.UrlParams["cardId"] == "" {
			tcp.RespondJsonError(conn, "card id not entered", tcp.INVALID_INPUT)
			return
		}

		cardId, _ := strconv.ParseUint(request.UrlParams["cardId"], 10, 64)

		cardToBeDelete := wallet.WalletCard{
			ID: uint(cardId),
		}

		//store in database using service
		err := app.WalletService().DeleteWalletCard(request.Context(), &cardToBeDelete)
		if err != nil {
			tcp.RespondJsonError(conn, err.Error(), tcp.INTERNAL_SERVER_ERROR)
			return
		}

		//show response result
		responseBody := map[string]any{
			"message": "your card had been deleted successfully",
		}

		tcp.RespondJsonSuccess(conn, responseBody)
		return
	}
}
