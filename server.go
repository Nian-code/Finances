package main

import (
	"github.com/Nian-code/Finances/Controller"
	"github.com/Nian-code/Finances/Services"
	"github.com/Nian-code/Finances/modules"
	"github.com/gin-gonic/gin"
	"net/http"
)

var(
	accountService = Services.New()
	accountController = Controller.New(accountService)
)

func main()  {
	server := gin.Default()
	server.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, modules.SayHello())
	})
	server.GET("/getAllAccounts/", accountController.FindAllAccount)

	server.POST("/save/account",  accountController.SaveAccount)

	server.POST("/save/accounts", accountController.SaveAccounts)

	server.PUT("/add/money/account", accountController.AddMoneyAccountByName)

	defer server.Run(":8080")


}

