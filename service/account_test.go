package service

import "testing"

type mockAccountHandler struct{}

func (handler *mockAccountHandler) HandleChannels(account *Account) *AccountWithChannels {
	return &AccountWithChannels{Account: account}
}

type mockAccountRepository struct{}

func (repo *mockAccountRepository) CreateNewAccount(request CreateAccountRequest) *Account {
	if request.Forename == "exists" && request.Surname == "exists" {
		return nil
	}

	account := NewAccount("123", request.Forename, request.Surname, 0)
	return &account
}

func (repo *mockAccountRepository) GetAccount(id string) *Account {
	if id == "exists" {
		account := NewAccount(id, "Test", "Account", 10)
		return &account
	}

	return nil
}

func TestCreateAccount(t *testing.T) {
	mockRepo := mockAccountRepository{}
	mockHandler := mockAccountHandler{}
	accountService := NewAccountService(&mockHandler, &mockRepo)

	t.Run("CreateFailure", func(t *testing.T) {
		request := NewCreateAccountRequest("exists", "exists")
		result := accountService.CreateAccount(request)

		if result != nil {
			t.Error("expected nil as the account could not be created")
		}

	})

	t.Run("CreateSuccess", func(t *testing.T) {
		request := NewCreateAccountRequest("Test", "123")
		result := accountService.CreateAccount(request)

		if result == nil {
			t.Error("expected an account handler to be created but got nil")
		}

		if result.Account == nil {
			t.Error("expected an account to be referenced but got nil")
		}
	})
}

func TestGetAccount(t *testing.T) {
	mockRepo := mockAccountRepository{}
	mockHandler := mockAccountHandler{}
	accountService := NewAccountService(&mockHandler, &mockRepo)

	t.Run("GetNotExists", func(t *testing.T) {
		result := accountService.GetAccount("Abc")

		if result != nil {
			t.Error("expected nil as the account does not exist.")
		}
	})

	t.Run("GetExists", func(t *testing.T) {
		result := accountService.GetAccount("exists")

		if result == nil {
			t.Error("expected an account handler to be created but got nil")
		}
	})
}
