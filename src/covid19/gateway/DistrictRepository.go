package gateway

import (
	"covid19/config/domain"
	"covid19/infra/database"
	u "covid19/utils"
	"errors"
	"time"
)

//AddDistrictCaseSummary add case to aundh banner ward
func (dbRepo *DatabaseRepository) AddDistrictCaseSummary(ward *domain.DistrictCase) (*database.DistrictCase, error) {
	var DBDistrictCase database.DistrictCase
	DBDistrictCase.Recovered = ward.Recovered
	DBDistrictCase.Death = ward.Death
	DBDistrictCase.Confirmed = ward.Confirmed
	if len(ward.Date) > 0 {
		DBDistrictCase.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddDistrictCaseSummaryDB(&DBDistrictCase)
}

//AddDistrictCaseSummaryDB add case to aundh banner ward db
func (dbRepo *DatabaseRepository) AddDistrictCaseSummaryDB(DBDistrictCase *database.DistrictCase) (*database.DistrictCase, error) {
	err := dbRepo.GetDBHandle().Create(&DBDistrictCase).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add district case summary,connection error: " + err.Error())
	}
	return DBDistrictCase, nil
}

//GetDistrictSummaryByCreateDate get district summary by create date
func (dbRepo *DatabaseRepository) GetDistrictSummaryByCreateDate(createDate string, endDate string) (*database.DistrictCase, error) {
	var DBDistrictCase database.DistrictCase
	err := dbRepo.GetDBHandle().Model(domain.DistrictCase{}).Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Find(&DBDistrictCase).Error
	return &DBDistrictCase, err
}
