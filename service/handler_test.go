package service

import "testing"

func TestDepositRequest(t *testing.T) {
	account := NewAccount("123", "Test", "Account", 0)

	handler := NewAccountHandlerService()
	channels := handler.HandleChannels(&account)

	const amount float32 = 10.5

	request := NewDepositRequest(amount)

	channels.DepositChannel <- request

	response := <-request.Response

	if response.Balance != amount {
		t.Errorf("expected response balance to be %f but got %f", amount, response.Balance)
	}

	if account.balance != amount {
		t.Errorf("expected account balance to be %f but got %f", amount, account.balance)
	}
}

func TestWithdrawRequest(t *testing.T) {
	const balance float32 = 20

	account := NewAccount("123", "Test", "Account", 20)

	handler := NewAccountHandlerService()
	channels := handler.HandleChannels(&account)

	const amount float32 = 15.5
	const expect float32 = balance - amount

	request := NewWithdrawRequest(amount)

	channels.WithdrawChannel <- request

	response := <-request.Response

	if response.Balance != expect {
		t.Errorf("expected response balance to be %f but got %f", expect, response.Balance)
	}

	if account.balance != expect {
		t.Errorf("expected account balance to be %f but got %f", expect, account.balance)
	}
}

func TestBalanceRequest(t *testing.T) {
	const balance float32 = 20

	account := NewAccount("123", "Test", "Account", 20)

	handler := NewAccountHandlerService()
	channels := handler.HandleChannels(&account)

	request := NewBalanceRequest()

	channels.BalanceChannel <- request

	response := <-request.Response

	if response != balance {
		t.Errorf("expected response balance to be %f but got %f", balance, response)
	}

	if account.balance != balance {
		t.Errorf("expected account balance to be %f but got %f", balance, account.balance)
	}
}
