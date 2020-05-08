package service

import (
	"covid19/config"
	"covid19/config/domain"
	u "covid19/utils"
	"encoding/json"
	"io/ioutil"
	"time"
)

//AddPatientSummary add patient summary
func (sh *Interactor) AddPatientSummary(patientSummary *domain.PatientSummary) (*domain.PatientSummary, error) {
	_, err := sh.Db.AddPatientSummary(patientSummary)
	if err != nil {
		return nil, err
	}
	defer sh.GeneratePatientDeltaSummaryJSON()
	defer sh.GeneratePatientSummaryJSON()
	defer sh.GenerateMetaJSON()
	return patientSummary, nil
}

//GetPatientSummaryByCreateDate get patient summary by create date
func (sh *Interactor) GetPatientSummaryByCreateDate(createDate string, endDate string) ([]domain.PatientSummary, error) {
	dbPatientSummary, err := sh.Db.GetPatientSummaryByCreateDate(createDate, endDate)
	if err != nil {
		return nil, err
	}
	patients := make([]domain.PatientSummary, 0)
	for _, patient := range dbPatientSummary {
		var patientSummary domain.PatientSummary
		u.Copy(&patientSummary, patient)
		patients = append(patients, patientSummary)
	}
	return patients, nil
}

//GetPatientDeltaSummary get delta of patient summary
func (sh *Interactor) GetPatientDeltaSummary() (domain.PatientDeltaSummary, error) {
	var patientDeltaSummary domain.PatientDeltaSummary
	today := time.Now().Format("2006-01-02")
	eDate := "2040-01-30"
	dbPatientSummary, err := sh.Db.GetPatientSummaryByCreateDate(today, eDate)
	if err == nil {
		patients := make([]domain.PatientSummary, 0)
		for _, patient := range dbPatientSummary {
			var patientSummary domain.PatientSummary
			u.Copy(&patientSummary, patient)
			patients = append(patients, patientSummary)
		}
		patientDeltaSummary.PatientSummary = patients
		todayPatientSummary, err := sh.Db.GetPatientSummaryByDate(today)
		if err == nil {
			previousDay := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
			previousDayPatientSummary, err := sh.Db.GetPatientSummaryByDate(previousDay)
			if err == nil {
				patientDeltaSummary.DeltaTotalTests = todayPatientSummary.TotalTests - previousDayPatientSummary.TotalTests
				patientDeltaSummary.DeltaInQuarantine = todayPatientSummary.InQuarantine - previousDayPatientSummary.InQuarantine
				patientDeltaSummary.DeltaInICU = todayPatientSummary.InICU - previousDayPatientSummary.InICU
				patientDeltaSummary.DeltaOnVentilator = todayPatientSummary.OnVentilator - previousDayPatientSummary.OnVentilator
			}
		}
	}
	return patientDeltaSummary, nil
}

//GeneratePatientDeltaSummaryJSON generate patient delta summary json
func (sh *Interactor) GeneratePatientDeltaSummaryJSON() error {
	districtCase, err := sh.GetPatientDeltaSummary()
	if err != nil {
		return err
	}
	file, _ := json.MarshalIndent(districtCase, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"patient-summary-delta.json", file, 0644)
	return nil
}

//GeneratePatientSummaryJSON generate patient summary with date range json
func (sh *Interactor) GeneratePatientSummaryJSON() error {
	districtCase, err := sh.GetPatientSummaryByCreateDate("2019-12-30", "2040-12-30")
	if err != nil {
		return err
	}
	file, _ := json.MarshalIndent(districtCase, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"patient-summary-range.json", file, 0644)
	return nil
}
