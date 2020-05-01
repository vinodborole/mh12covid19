package interactorinterface

import (
	"covid19/config/domain"
	"covid19/infra/database"
	u "covid19/utils"
)

//DatabaseRepository database repository
type DatabaseRepository interface {
	//Transaction Operations
	Ping() error
	OpenTransaction() error
	CommitTransaction() error
	RollBackTransaction() error

	//User
	Authenticate(email string, password string) (*database.User, error)
	CreateUser(account *domain.User) (*database.User, error)
	GetUserByEmail(email string) (*database.User, error)

	//Ward
	CreateWard(ward *domain.Ward) (*database.Ward, error)
	GetWards(page string) *u.Data
}
