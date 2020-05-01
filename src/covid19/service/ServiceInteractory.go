package service

import (
	"context"
	InteractorIntf "covid19/service/interactorinterface"
	"sync"
)

//Interactor is the receiver object
type Interactor struct {
	Db         InteractorIntf.DatabaseRepository
	AppContext context.Context
	DBMutex    sync.Mutex
	Refresh    bool
}
