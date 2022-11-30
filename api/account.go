package api

import (
	"bank/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AccountApi struct {
	service service.AccountService
}

func NewAccountApi(service service.AccountService) AccountApi {
	return AccountApi{
		service: service,
	}
}

func (api AccountApi) CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request service.CreateAccountRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		log.Info().Msgf("create account with forename %q and surname %q", request.Forename, request.Surname)

		handler := api.service.CreateAccount(request)

		if handler == nil {
			log.Info().Msgf("failed to create account with forename %q and surname %q", request.Forename, request.Surname)
			c.JSON(http.StatusBadRequest, gin.H{"error": "account could not be created."})
			return
		}

		c.JSON(http.StatusOK, handler.Account)
	}
}

func (api AccountApi) DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		result := api.service.DeleteAccount(id)

		if !result {
			accountNotFound(c, id)
		}
	}
}

func (api *AccountApi) GetAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		handler := api.service.GetAccount(id)

		if handler == nil {
			accountNotFound(c, id)
			return
		}

		log.Info().Msgf("account with id %q was found.", id)
		c.JSON(http.StatusOK, handler.Account)
	}
}

func (api *AccountApi) GetBalance() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		handler := api.service.GetAccount(id)

		if handler == nil {
			accountNotFound(c, id)
			return
		}

		request := service.NewBalanceRequest()

		handler.BalanceChannel <- request

		balance := <-request.Response

		log.Info().Msgf("account balance of %f obtained for id %q", balance, id)
		c.JSON(http.StatusOK, balance)
	}
}

func (api *AccountApi) DepositIntoAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var request service.DepositRequest
		request.Response = make(chan service.DepositResult)

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		handler := api.service.GetAccount(id)

		if handler == nil {
			accountNotFound(c, id)
			return
		}

		handler.DepositChannel <- request

		response := <-request.Response

		log.Info().Msgf("deposited %f into account with id %q. New balance is %f.", request.Amount, id, response.Balance)
		c.JSON(http.StatusOK, response.Balance)
	}
}

func (api *AccountApi) WithdrawFromAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var request service.WithdrawRequest
		request.Response = make(chan service.WithdrawResult)

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		handler := api.service.GetAccount(id)

		if handler == nil {
			accountNotFound(c, id)
			return
		}

		handler.WithdrawChannel <- request

		response := <-request.Response

		if !response.HadEnoughMoney {
			log.Info().Msgf("account %q did not have enough money to perform a withdraw of %f. Account balance is %f",
				id, request.Amount, response.Balance)

			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("account %q does not have enough money to perform a withdraw of %f. Account balance is %f.",
					id, request.Amount, response.Balance),
			})

			return
		}

		log.Info().Msgf("Withdrawed %f from account with id %q. New balance is %f.", request.Amount, id)
		c.JSON(http.StatusOK, response.Balance)
	}
}

func accountNotFound(c *gin.Context, id string) {
	log.Info().Msgf("account with id %q was not found.", id)
	c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("account with id %q could not be found.", id)})
}
