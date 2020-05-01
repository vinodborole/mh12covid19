package infra

import (
	"covid19/gateway"
	"covid19/infra/database"
	"covid19/service"
	"sync"
)

//ServiceInteractor  provides reciever object for services
var ServiceInteractor *service.Interactor
var once sync.Once

//GetUseCaseInteractor gets a singleton instance of DeviceInteractor with DatabaseRepository
func GetUseCaseInteractor() *service.Interactor {
	once.Do(func() {
		DatabaseRepository := gateway.DatabaseRepository{Database: database.GetWorkingInstance()}
		ServiceInteractor = &service.Interactor{Db: &DatabaseRepository}
	})
	return ServiceInteractor
}
