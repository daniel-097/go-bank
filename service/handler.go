package service

import "github.com/rs/zerolog/log"

type accountHandler struct {
}

func NewAccountHandlerService() AccountHandler {
	return &accountHandler{}
}

func (handler *accountHandler) HandleChannels(account *Account) *AccountWithChannels {
	accountWithChannels := &AccountWithChannels{
		Account:         account,
		BalanceChannel:  make(chan BalanceRequest),
		DepositChannel:  make(chan DepositRequest),
		WithdrawChannel: make(chan WithdrawRequest),
	}

	go accountWithChannels.monitor()

	return accountWithChannels
}

func (handler *AccountWithChannels) monitor() {
	for {
		select {
		case request := <-handler.DepositChannel:
			handleDepositRequest(request, handler.Account)
		case request := <-handler.WithdrawChannel:
			handleWithdrawRequest(request, handler.Account)
		case request := <-handler.BalanceChannel:
			handleBalanceRequest(request, handler.Account)
		}
	}
}

func handleBalanceRequest(request BalanceRequest, account *Account) {
	log.Info().Msgf("request recieved to obtain balance for account %q", account.Id)
	request.Response <- account.balance
}

func handleDepositRequest(request DepositRequest, account *Account) {
	log.Info().Msgf("request recieved to deposit %f into account %q", request.Amount, account.Id)

	account.balance += request.Amount

	request.Response <- DepositResult{Balance: account.balance}
}

func handleWithdrawRequest(request WithdrawRequest, account *Account) {
	log.Info().Msgf("request recieved to withdraw %f from account %q", request.Amount, account.Id)

	if account.balance > request.Amount {
		account.balance -= request.Amount
		request.Response <- WithdrawResult{HadEnoughMoney: true, Balance: account.balance}
	} else {
		request.Response <- WithdrawResult{HadEnoughMoney: false, Balance: account.balance}
	}
}
