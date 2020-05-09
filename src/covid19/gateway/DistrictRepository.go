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

//AddDistrictDetail add delta
func (dbRepo *DatabaseRepository) AddDistrictDetail(districtDetail *domain.DistrictDetail) (*database.DistrictDetail, error) {
	var DBDistrictDetail database.DistrictDetail
	DBDistrictDetail.DeltaActive = districtDetail.DeltaActive
	DBDistrictDetail.DeltaRecovered = districtDetail.DeltaRecovered
	DBDistrictDetail.DeltaDeath = districtDetail.DeltaDeath
	DBDistrictDetail.DeltaConfirmed = districtDetail.DeltaConfirmed
	return dbRepo.AddDistrictDetailDB(&DBDistrictDetail)
}

//AddDistrictDetailDB add delta to db
func (dbRepo *DatabaseRepository) AddDistrictDetailDB(DBDistrictDetail *database.DistrictDetail) (*database.DistrictDetail, error) {
	err := dbRepo.GetDBHandle().Create(&DBDistrictDetail).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add district detail ,connection error: " + err.Error())
	}
	return DBDistrictDetail, nil
}

//GetDistrictSummaryByCreateDate get district summary by create date
func (dbRepo *DatabaseRepository) GetDistrictSummaryByCreateDate(createDate string, endDate string) ([]database.DistrictCase, error) {
	var DBDistrictCase []database.DistrictCase
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at asc").Limit(100).Find(&DBDistrictCase).Error
	return DBDistrictCase, err
}

//GetDistrictDetailByCreateDate get district detail by create date
func (dbRepo *DatabaseRepository) GetDistrictDetailByCreateDate(createDate string, endDate string) ([]database.DistrictDetail, error) {
	var DBDistrictDetail []database.DistrictDetail
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at asc").Limit(100).Find(&DBDistrictDetail).Error
	return DBDistrictDetail, err
}

//GetDistrictSummaryByDate get district by date
func (dbRepo *DatabaseRepository) GetDistrictSummaryByDate(date string) (database.DistrictCase, error) {
	var DBDistrictCase database.DistrictCase
	err := dbRepo.GetDBHandle().Where("date(created_at) = ? ", date).Order("created_at desc").Find(&DBDistrictCase).Error
	return DBDistrictCase, err
}
