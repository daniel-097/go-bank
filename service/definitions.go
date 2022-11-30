package service

type Account struct {
	Id       string `json:"id"`
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
	balance  float32
}

type AccountWithChannels struct {
	Account         *Account
	DepositChannel  chan DepositRequest
	WithdrawChannel chan WithdrawRequest
	BalanceChannel  chan BalanceRequest
}

type CreateAccountRequest struct {
	Forename string
	Surname  string
}

type BalanceRequest struct {
	Response chan float32
}

type DepositRequest struct {
	Amount   float32 `json:"amount"`
	Response chan DepositResult
}

type DepositResult struct {
	Balance float32
}

type WithdrawRequest struct {
	Amount   float32 `json:"amount"`
	Response chan WithdrawResult
}

type WithdrawResult struct {
	Balance        float32
	HadEnoughMoney bool
}

func NewAccount(id string, forename string, surname string, balance float32) Account {
	return Account{
		Id:       id,
		Forename: forename,
		Surname:  surname,
		balance:  balance,
	}
}

func NewBalanceRequest() BalanceRequest {
	return BalanceRequest{
		Response: make(chan float32),
	}
}

func NewDepositRequest(amount float32) DepositRequest {
	return DepositRequest{
		Amount:   amount,
		Response: make(chan DepositResult),
	}
}

func NewWithdrawRequest(amount float32) WithdrawRequest {
	return WithdrawRequest{
		Amount:   amount,
		Response: make(chan WithdrawResult),
	}
}

func NewCreateAccountRequest(forename string, surname string) CreateAccountRequest {
	return CreateAccountRequest{
		Forename: forename,
		Surname:  surname,
	}
}
