package service

type AccountHandler interface {
	HandleChannels(account *Account) *AccountWithChannels
}

type AccountRepository interface {
	CreateNewAccount(request CreateAccountRequest) *Account
	DeleteAccount(id string) bool
	GetAccount(id string) *Account
}

type AccountService interface {
	CreateAccount(request CreateAccountRequest) *AccountWithChannels
	DeleteAccount(id string) bool
	GetAccount(id string) *AccountWithChannels
}

type accountService struct {
	handler AccountHandler
	repo    AccountRepository
}

func NewAccountService(handler AccountHandler, repo AccountRepository) AccountService {
	return &accountService{handler: handler, repo: repo}
}

func (service *accountService) CreateAccount(request CreateAccountRequest) *AccountWithChannels {
	account := service.repo.CreateNewAccount(request)

	if account == nil {
		return nil
	}

	return service.handler.HandleChannels(account)
}

func (service *accountService) DeleteAccount(id string) bool {
	return service.repo.DeleteAccount(id)
}

func (service *accountService) GetAccount(id string) *AccountWithChannels {
	account := service.repo.GetAccount(id)

	if account != nil {
		return service.handler.HandleChannels(account)
	}

	return nil
}
