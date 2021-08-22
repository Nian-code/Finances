package Services

import (
	"errors"
	q "github.com/Nian-code/Finances/DB"
	d "github.com/Nian-code/Finances/DTO"
	e "github.com/Nian-code/Finances/Entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountService interface {
	SaveAccount(account d.AccountDTO) (*mongo.InsertManyResult, error)
	SaveAccounts(accounts []interface{}) (*mongo.InsertManyResult, error)
	AddMoneyAccountByName(name string, money float64) (e.Account, error)
	FindAll() (interface{}, error)
}

type accountService struct {
	accounts []e.Account
}

func (a *accountService) AddMoneyAccountByName(name string, money float64)  (account e.Account, err error) {
	account, i , err := a.filterAccountByName(name)

	if err != nil{
		return account, err
	}
	account.Money += money
	a.accounts[i] = account
	return account, err
}

func (a *accountService) filterAccountByName(name string) (account e.Account, i int, err error){
	for i , acc := range a.accounts {
		if acc.Title == name {
			return acc, i, nil
		}
	}
	err = errors.New("account not found")

	return
}


func (a *accountService) SaveAccounts(accounts  []interface{}) (account *mongo.InsertManyResult, err error){
	return q.InsertDocuments("Accounts", accounts)
}

func (a *accountService) SaveAccount(acc d.AccountDTO) (account *mongo.InsertManyResult, err error) {
	return q.InsertDocuments("Accounts", []interface{}{acc})
}

func (a *accountService) FindAll() (accounts interface{}, err error) {
	return q.FindAll("Accounts", "Account")
}

func New() AccountService {
	return &accountService{}
}
