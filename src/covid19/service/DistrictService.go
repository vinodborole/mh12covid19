package service

import (
	"covid19/config"
	"covid19/config/domain"
	"encoding/json"
	"io/ioutil"
)

//AddDistrictCaseSummary add district summary
func (sh *Interactor) AddDistrictCaseSummary(districtCase *domain.DistrictCase) (*domain.DistrictCase, error) {
	dbDistrictCase, err := sh.Db.AddDistrictCaseSummary(districtCase)
	if err != nil {
		return nil, err
	}
	districtCase.TotalPositive = dbDistrictCase.Confirmed + dbDistrictCase.Death + dbDistrictCase.Recovered
	districtCase.Active = districtCase.TotalPositive - (dbDistrictCase.Death + dbDistrictCase.Recovered)
	defer sh.GenerateDistrictSummaryJSON()
	return districtCase, nil
}

//GetDistrictSummaryByCreateDate get district summary by create date
func (sh *Interactor) GetDistrictSummaryByCreateDate(createDate string, endDate string) ([]domain.DistrictCase, error) {
	dbDistrictCase, err := sh.Db.GetDistrictSummaryByCreateDate(createDate, endDate)
	if err != nil {
		return nil, err
	}
	districtCases := make([]domain.DistrictCase, 0)
	for _, district := range dbDistrictCase {
		totalPositive := district.Confirmed + district.Death + district.Recovered
		active := totalPositive - (district.Death + district.Recovered)
		date := district.CreatedAt.Format("2006-01-02")
		districtCase := domain.DistrictCase{TotalPositive: totalPositive, Active: active, Confirmed: district.Confirmed, Recovered: district.Recovered, Death: district.Death, Date: date}
		districtCases = append(districtCases, districtCase)
	}
	return districtCases, nil
}

//GenerateDistrictSummaryJSON generate district summary json
func (sh *Interactor) GenerateDistrictSummaryJSON() error {
	districtCase, err := sh.GetDistrictSummaryByCreateDate("2019-12-01", "2040-12-31")
	if err != nil {
		return err
	}
	file, _ := json.MarshalIndent(districtCase, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"district-summary.json", file, 0644)
	return nil
}
