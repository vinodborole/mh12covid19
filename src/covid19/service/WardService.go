package service

import (
	"covid19/config/domain"
	"covid19/infra/database"
	u "covid19/utils"
)

//CreateWard create ward service
func (sh *Interactor) CreateWard(ward *domain.Ward) (*database.Ward, error) {
	dbWard, err := sh.Db.CreateWard(ward)
	if err != nil {
		return nil, err
	}
	return dbWard, nil
}

//GetWards get wards service
func (sh *Interactor) GetWards(page string) *u.Data {
	return sh.Db.GetWards(page)
}
