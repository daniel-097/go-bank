package repository

import (
	"bank/service"
	"testing"
)

func TestCreateNewAccount(t *testing.T) {
	repo := NewAccountRepository()
	request := service.NewCreateAccountRequest("Some", "Name")

	account := repo.CreateNewAccount(request)

	if account == nil {
		t.Error("expected an account but got nil")
		return
	}

	if account.Forename != request.Forename {
		t.Errorf("expected forename to be %q but got %q", account.Forename, request.Forename)
	}

	if account.Surname != request.Surname {
		t.Errorf("expected forname to be %q but got %q", account.Surname, request.Surname)
	}

	if account.Id == "" {
		t.Errorf("expected an id to be generated but got %q", account.Id)
	}
}

func TestGetAccount(t *testing.T) {
	repo := accountRepository{
		accounts: make(map[string]*service.Account),
	}

	const id string = "123"
	repoAccount := service.NewAccount(id, "Test", "Abc", 0)

	repo.accounts[id] = &repoAccount

	account := repo.GetAccount(id)

	if account == nil {
		t.Errorf("expected an account with id %q but got nil", id)
		return
	}

	if account != repo.accounts[id] {
		t.Errorf("expected account to be account with id %q", id)
	}
}

func TestDeleteAccount(t *testing.T) {
	repo := accountRepository{
		accounts: make(map[string]*service.Account),
	}

	const id string = "123"
	repoAccount := service.NewAccount(id, "Test", "Abc", 0)

	repo.accounts[id] = &repoAccount

	result := repo.DeleteAccount(id)

	if result == false {
		t.Error("expected true but got false")
		return
	}

	_, exists := repo.accounts[id]

	if exists {
		t.Errorf("expected account with id %q to be deleted", id)
	}
}
