package repository

import (
	"bank/service"

	"github.com/google/uuid"
)

type accountRepository struct {
	accounts map[string]*service.Account
}

func NewAccountRepository() service.AccountRepository {
	return &accountRepository{
		accounts: make(map[string]*service.Account),
	}
}

func (repo *accountRepository) CreateNewAccount(request service.CreateAccountRequest) *service.Account {
	account := service.NewAccount(uuid.NewString(), request.Forename, request.Surname, 0)
	repo.accounts[account.Id] = &account

	return &account
}

func (repo *accountRepository) GetAccount(id string) *service.Account {
	account, exists := repo.accounts[id]

	if exists {
		return account
	}

	return nil
}
