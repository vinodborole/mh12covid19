package gateway

import (
	"covid19/config/domain"
	"covid19/infra/database"
	u "covid19/utils"
	"errors"
)

//CreateWard create account
func (dbRepo *DatabaseRepository) CreateWard(ward *domain.Ward) (*database.Ward, error) {
	var DBWard database.Ward
	u.Copy(&DBWard, ward)
	return dbRepo.CreateWardDB(&DBWard)
}

//CreateWardDB create ward db
func (dbRepo *DatabaseRepository) CreateWardDB(DBWard *database.Ward) (*database.Ward, error) {
	err := dbRepo.GetDBHandle().Create(&DBWard).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to create ward,connection error: " + err.Error())
	}
	return DBWard, nil
}

//GetWards get wards
func (dbRepo *DatabaseRepository) GetWards(page string) *u.Data {
	var wards []database.Ward
	// need this limit somewhere in constant/env file
	var limit string = "100"
	orderBy := []string{"ID"}
	paginator := u.Paginator{DB: dbRepo.GetDBHandle(), OrderBy: orderBy, Page: page, PerPage: limit}
	data := paginator.Paginate(&wards)
	return data
}
