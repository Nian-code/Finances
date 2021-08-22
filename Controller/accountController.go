package Controller

import (
	"github.com/gin-gonic/gin"
	d "github/nian-code/finances/DTO"
	"github/nian-code/finances/Services"
	"net/http"
)

type AccountController interface {
	SaveAccount(ctx *gin.Context)
	SaveAccounts(ctx *gin.Context)
	AddMoneyAccountByName(ctx *gin.Context)
	FindAllAccount(ctx *gin.Context)
}

type controller struct {
	service Services.AccountService
}

func (c *controller) AddMoneyAccountByName(ctx *gin.Context) {
	var acc d.AccountDTO
	ctx.BindJSON(&acc)

	account, err := c.service.AddMoneyAccountByName(acc.Title, acc.Money)

	if err != nil{
		ctx.JSON(404, gin.H{
			"Error" : err.Error(),
		})
		return
	}

	ctx.JSON(200, account)
}

func (c *controller) SaveAccounts(ctx *gin.Context) {
	var accDTOS []interface{}
	ctx.BindJSON(&accDTOS)
	res, err := c.service.SaveAccounts(accDTOS)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		} )
	}
	ctx.JSON(200, res)
}

func (c *controller) SaveAccount(ctx *gin.Context){
	var acc d.AccountDTO
	ctx.BindJSON(&acc)
	res , err :=c.service.SaveAccount(acc)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		} )
	}
	ctx.JSON(200, res)
}

func (c *controller) FindAllAccount(ctx *gin.Context) {
	accounts, err := c.service.FindAll()

	if err != nil{
		 ctx.JSON(http.StatusBadRequest, gin.H{
		 	"Error": err,
		 } )
	}

	ctx.JSON(200, accounts)
}

func New(service Services.AccountService) controller {
	return controller{
		service: service,
	}
}
