package main

import (
	"bank/api"
	"bank/repository"
	"bank/service"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	repo := repository.NewAccountRepository()
	handler := service.NewAccountHandlerService()
	service := service.NewAccountService(handler, repo)
	api := api.NewAccountApi(service)

	router := gin.Default()

	account := router.Group("/api/account")

	account.GET(":id", api.GetAccount())
	account.DELETE(":id", api.DeleteAccount())
	account.PUT(":id/deposit", api.DepositIntoAccount())
	account.PUT(":id/withdraw", api.WithdrawFromAccount())
	account.GET(":id/balance", api.GetBalance())
	account.POST("", api.CreateAccount())

	router.Run(":80")
}
