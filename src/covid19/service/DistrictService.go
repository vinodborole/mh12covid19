package service

import (
	"covid19/config/domain"
)

//AddDistrictCaseSummary add district summary
func (sh *Interactor) AddDistrictCaseSummary(districtCase *domain.DistrictCase) (*domain.DistrictCase, error) {
	dbDistrictCase, err := sh.Db.AddDistrictCaseSummary(districtCase)
	if err != nil {
		return nil, err
	}
	districtCase.TotalPositive = dbDistrictCase.Confirmed + dbDistrictCase.Death + dbDistrictCase.Recovered
	districtCase.Active = districtCase.TotalPositive - (dbDistrictCase.Death + dbDistrictCase.Recovered)
	return districtCase, nil
}

//GetDistrictSummaryByCreateDate get district summary by create date
func (sh *Interactor) GetDistrictSummaryByCreateDate(createDate string, endDate string) (*domain.DistrictCase, error) {
	dbDistrictCase, err := sh.Db.GetDistrictSummaryByCreateDate(createDate, endDate)
	if err != nil {
		return nil, err
	}
	totalPositive := dbDistrictCase.Confirmed + dbDistrictCase.Death + dbDistrictCase.Recovered
	active := totalPositive - (dbDistrictCase.Death + dbDistrictCase.Recovered)
	date := dbDistrictCase.CreatedAt.Format("2006-01-02")
	districtCase := domain.DistrictCase{TotalPositive: totalPositive, Active: active, Confirmed: dbDistrictCase.Confirmed, Recovered: dbDistrictCase.Recovered, Death: dbDistrictCase.Death, Date: date}
	return &districtCase, nil
}
