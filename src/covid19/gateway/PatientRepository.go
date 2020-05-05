package gateway

import (
	"covid19/config/domain"
	"covid19/infra/database"
	u "covid19/utils"
	"errors"
	"time"
)

//AddPatientSummary add patient summary
func (dbRepo *DatabaseRepository) AddPatientSummary(patientSummary *domain.PatientSummary) (*database.PatientSummary, error) {
	var DBPatientSummary database.PatientSummary
	DBPatientSummary.TotalTests = patientSummary.TotalTests
	DBPatientSummary.InQuarantine = patientSummary.InQuarantine
	DBPatientSummary.InICU = patientSummary.InICU
	DBPatientSummary.OnVentilator = patientSummary.OnVentilator
	if len(patientSummary.Date) > 0 {
		DBPatientSummary.CreatedAt, _ = time.Parse("2006-01-02", patientSummary.Date)
	}
	return dbRepo.AddPatientSummaryDB(&DBPatientSummary)
}

//AddPatientSummaryDB add patient summary to db
func (dbRepo *DatabaseRepository) AddPatientSummaryDB(DBPatientSummary *database.PatientSummary) (*database.PatientSummary, error) {
	err := dbRepo.GetDBHandle().Create(&DBPatientSummary).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add patient summary,connection error: " + err.Error())
	}
	return DBPatientSummary, nil
}

//GetPatientSummaryByCreateDate get patient summary by create date
func (dbRepo *DatabaseRepository) GetPatientSummaryByCreateDate(createDate string, endDate string) ([]database.PatientSummary, error) {
	var DBPatientSummary []database.PatientSummary
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at asc").Limit(100).Find(&DBPatientSummary).Error
	return DBPatientSummary, err
}

//GetPatientSummaryByDate get patient summary by date
func (dbRepo *DatabaseRepository) GetPatientSummaryByDate(date string) (database.PatientSummary, error) {
	var DBPatientSummary database.PatientSummary
	err := dbRepo.GetDBHandle().Where("date(created_at) = ? ", date).Order("created_at desc").Find(&DBPatientSummary).Error
	return DBPatientSummary, err
}
