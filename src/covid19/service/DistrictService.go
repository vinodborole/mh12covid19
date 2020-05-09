package service

import (
	"covid19/config"
	"covid19/config/domain"
	"covid19/infra/database"
	"encoding/json"
	"io/ioutil"
	"time"
)

//AddDistrictCaseSummary add district summary
func (sh *Interactor) AddDistrictCaseSummary(districtCase *domain.DistrictCase) (*domain.DistrictCase, error) {
	dbDistrictCase, err := sh.Db.AddDistrictCaseSummary(districtCase)
	if err != nil {
		return nil, err
	}
	districtCase.Active = districtCase.Confirmed - (dbDistrictCase.Death + dbDistrictCase.Recovered)
	sh.AddDistrictDetail()
	defer sh.GenerateDistrictSummaryLatestJSON()
	defer sh.GenerateDistrictSummaryJSON()
	defer sh.GenerateMetaJSON()
	return districtCase, nil
}

//AddDistrictDetail add district detail
func (sh *Interactor) AddDistrictDetail() (domain.DistrictDetail, error) {
	var districtDetail domain.DistrictDetail
	today := time.Now().Format("2006-01-02")
	todayDistrictCase, err := sh.Db.GetDistrictSummaryByDate(today)
	if err == nil {
		previousDay := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		previousDayDistrictCase, err := sh.Db.GetDistrictSummaryByDate(previousDay)
		if err == nil {
			todayActive := todayDistrictCase.Confirmed - (todayDistrictCase.Death + todayDistrictCase.Recovered)
			previousDayActive := previousDayDistrictCase.Confirmed - (previousDayDistrictCase.Death + previousDayDistrictCase.Recovered)
			districtDetail.DeltaActive = todayActive - previousDayActive
			districtDetail.DeltaConfirmed = todayDistrictCase.Confirmed - previousDayDistrictCase.Confirmed
			districtDetail.DeltaDeath = todayDistrictCase.Death - previousDayDistrictCase.Death
			districtDetail.DeltaRecovered = todayDistrictCase.Recovered - previousDayDistrictCase.Recovered
			_, err := sh.Db.AddDistrictDetail(&districtDetail)
			if err != nil {
				return domain.DistrictDetail{}, err
			}
		}
	}
	defer sh.GenerateDistrictDetailJSON()
	return districtDetail, nil
}

//GetDistrictSummaryByCreateDate get district summary by create date
func (sh *Interactor) GetDistrictSummaryByCreateDate(createDate string, endDate string) ([]domain.DistrictCase, error) {
	dbDistrictCase, err := sh.Db.GetDistrictSummaryByCreateDate(createDate, endDate)
	if err != nil {
		return nil, err
	}
	districtCases := make([]domain.DistrictCase, 0)
	for _, district := range dbDistrictCase {
		active := district.Confirmed - (district.Death + district.Recovered)
		date := district.CreatedAt.Format("2006-01-02")
		districtCase := domain.DistrictCase{Active: active, Confirmed: district.Confirmed, Recovered: district.Recovered, Death: district.Death, Date: date}
		districtCases = append(districtCases, districtCase)
	}
	return districtCases, nil
}

//GetLatestDistrictSummary get latest district summary
func (sh *Interactor) GetLatestDistrictSummary() (domain.DistrictDetail, error) {
	var districtDetail domain.DistrictDetail
	today := time.Now().Format("2006-01-02")
	eDate := "2040-01-30"
	dbDistrictCase, err := sh.Db.GetDistrictSummaryByCreateDate(today, eDate)
	if err == nil {
		districtCases := make([]domain.DistrictCase, 0)
		for _, district := range dbDistrictCase {
			active := district.Confirmed - (district.Death + district.Recovered)
			date := district.CreatedAt.Format("2006-01-02")
			districtCase := domain.DistrictCase{Active: active, Confirmed: district.Confirmed, Recovered: district.Recovered, Death: district.Death, Date: date}
			districtCases = append(districtCases, districtCase)
		}
		districtDetail.DistrictCases = districtCases

		todayDistrictCase, err := sh.Db.GetDistrictSummaryByDate(today)
		if err == nil {
			previousDay := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
			previousDayDistrictCase, err := sh.Db.GetDistrictSummaryByDate(previousDay)
			if err == nil {
				todayActive := todayDistrictCase.Confirmed - (todayDistrictCase.Death + todayDistrictCase.Recovered)
				previousDayActive := previousDayDistrictCase.Confirmed - (previousDayDistrictCase.Death + previousDayDistrictCase.Recovered)

				districtDetail.DeltaActive = todayActive - previousDayActive
				districtDetail.DeltaConfirmed = todayDistrictCase.Confirmed - previousDayDistrictCase.Confirmed
				districtDetail.DeltaDeath = todayDistrictCase.Death - previousDayDistrictCase.Death
				districtDetail.DeltaRecovered = todayDistrictCase.Recovered - previousDayDistrictCase.Recovered
			}
		}
	}
	return districtDetail, nil
}

//GenerateDistrictDetailJSON generate district detail JSON
func (sh *Interactor) GenerateDistrictDetailJSON() error {
	sDate := "2019-12-30"
	eDate := time.Now().Format("2006-01-02")
	districtDetail, err := sh.GetDistrictDetailByDateRange(sDate, eDate)
	if err != nil {
		return err
	}
	file, _ := json.MarshalIndent(districtDetail, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"district-delta-graph.json", file, 0644)
	return nil
}

//GetDistrictDetailByDateRange ger district detail by date range
func (sh *Interactor) GetDistrictDetailByDateRange(startDate string, endDate string) ([]database.DistrictDetail, error) {
	return sh.Db.GetDistrictDetailByCreateDate(startDate, endDate)
}

//GenerateDistrictSummaryLatestJSON generate district summary latest json
func (sh *Interactor) GenerateDistrictSummaryLatestJSON() error {
	districtCase, err := sh.GetLatestDistrictSummary()
	if err != nil {
		return err
	}
	file, _ := json.MarshalIndent(districtCase, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"district-summary-latest.json", file, 0644)
	return nil
}

//GenerateDistrictSummaryJSON generate district summary with date range json
func (sh *Interactor) GenerateDistrictSummaryJSON() error {
	districtCase, err := sh.GetDistrictSummaryByCreateDate("2019-12-30", "2040-12-30")
	if err != nil {
		return err
	}
	file, _ := json.MarshalIndent(districtCase, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"district-summary-range.json", file, 0644)
	return nil
}
