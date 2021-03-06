package interactorinterface

import (
	"covid19/config/domain"
	"covid19/infra/database"
	u "covid19/utils"
)

//DatabaseRepository database repository
type DatabaseRepository interface {
	//Transaction Operations
	Ping() error
	OpenTransaction() error
	CommitTransaction() error
	RollBackTransaction() error

	//User
	Authenticate(email string, password string) (*database.User, error)
	CreateUser(account *domain.User) (*database.User, error)
	GetUserByEmail(email string) (*database.User, error)

	//Ward
	CreateWard(ward *domain.Ward) (*database.Ward, error)
	GetWards(page string) *u.Data
	WardExists(code string) bool
	GetWardByCode(name string) (*database.Ward, error)

	//Ward Wise
	AddAundhBannerWardCase(ward *domain.AundhBaner) (*database.AundhBaner, error)
	AddBhawaniPethWardCase(ward *domain.BhawaniPeth) (*database.BhawaniPeth, error)
	AddBibwewadiWardCase(ward *domain.Bibwewadi) (*database.Bibwewadi, error)
	AddDhankawadiSahakarnagarWardCase(ward *domain.DhankawadiSahakarnagar) (*database.DhankawadiSahakarnagar, error)
	AddDholePatilWardCase(ward *domain.DholePatil) (*database.DholePatil, error)
	AddHadapsarMundhwaWardCase(ward *domain.HadapsarMundhwa) (*database.HadapsarMundhwa, error)
	AddKasbaVisharambaghwadaWardCase(ward *domain.KasbaVisharambaghwada) (*database.KasbaVisharambaghwada, error)
	AddKondhwaYewalewadiWardCase(ward *domain.KondhwaYewalewadi) (*database.KondhwaYewalewadi, error)
	AddKothrudBavdhanWardCase(ward *domain.KothrudBavdhan) (*database.KothrudBavdhan, error)
	AddNagarRoadWadgoansheriWardCase(ward *domain.NagarRoadWadgoansheri) (*database.NagarRoadWadgoansheri, error)
	AddShivajiNagarGholeRoadWardCase(ward *domain.ShivajiNagarGholeRoad) (*database.ShivajiNagarGholeRoad, error)
	AddSinhagadRoadWardCase(ward *domain.SinghagadRoad) (*database.SinghagadRoad, error)
	AddWanawadiRamtekdiWardCase(ward *domain.WanawadiRamtekdi) (*database.WanawadiRamtekdi, error)
	AddWarjeKarveNagarWardCase(ward *domain.WarjeKarveNagar) (*database.WarjeKarveNagar, error)
	AddYerwadaKalasDhanoriWardCase(ward *domain.YerwadaKalasDhanori) (*database.YerwadaKalasDhanori, error)
	AddPuneRuralWardCase(ward *domain.PuneRural) (*database.PuneRural, error)
	AddPuneCantonmentWardCase(ward *domain.PuneCantonment) (*database.PuneCantonment, error)

	//existing total confirmed cases
	GetExistingTotalConfirmedAundhBannerWardCase() (int, error)
	GetExistingTotalConfirmedBhawaniPethWardCase() (int, error)
	GetExistingTotalConfirmedBibwewadiWardCase() (int, error)
	GetExistingTotalConfirmedDhankawadiWardCase() (int, error)
	GetExistingTotalConfirmedDholePatilWardCase() (int, error)
	GetExistingTotalConfirmedHadapsarMundhwaWardCase() (int, error)
	GetExistingTotalConfirmedKasbaWardCase() (int, error)
	GetExistingTotalConfirmedKondhwaWardCase() (int, error)
	GetExistingTotalConfirmedKothrudWardCase() (int, error)
	GetExistingTotalConfirmedWadgoansheriWardCase() (int, error)
	GetExistingTotalConfirmedShivajiNagarWardCase() (int, error)
	GetExistingTotalConfirmedSinhagadRoadWardCase() (int, error)
	GetExistingTotalConfirmedWanawadiWardCase() (int, error)
	GetExistingTotalConfirmedWarjeWardCase() (int, error)
	GetExistingTotalConfirmedYerwadaWardCase() (int, error)
	GetExistingTotalConfirmedPuneRuralWardCase() (int, error)
	GetExistingTotalConfirmedPuneCantonmentWardCase() (int, error)

	//existing total recovered cases
	GetExistingTotalRecoveredAundhBannerWardCase() (int, error)
	GetExistingTotalRecoveredBhawaniPethWardCase() (int, error)
	GetExistingTotalRecoveredBibwewadiWardCase() (int, error)
	GetExistingTotalRecoveredDhankawadiWardCase() (int, error)
	GetExistingTotalRecoveredDholePatilWardCase() (int, error)
	GetExistingTotalRecoveredHadapsarMundhwaWardCase() (int, error)
	GetExistingTotalRecoveredKasbaWardCase() (int, error)
	GetExistingTotalRecoveredKondhwaWardCase() (int, error)
	GetExistingTotalRecoveredKothrudWardCase() (int, error)
	GetExistingTotalRecoveredWadgoansheriWardCase() (int, error)
	GetExistingTotalRecoveredShivajiNagarWardCase() (int, error)
	GetExistingTotalRecoveredSinhagadRoadWardCase() (int, error)
	GetExistingTotalRecoveredWanawadiWardCase() (int, error)
	GetExistingTotalRecoveredWarjeWardCase() (int, error)
	GetExistingTotalRecoveredYerwadaWardCase() (int, error)
	GetExistingTotalRecoveredPuneRuralWardCase() (int, error)
	GetExistingTotalRecoveredPuneCantonmentWardCase() (int, error)

	//existing total death cases
	GetExistingTotalDeathAundhBannerWardCase() (int, error)
	GetExistingTotalDeathBhawaniPethWardCase() (int, error)
	GetExistingTotalDeathBibwewadiWardCase() (int, error)
	GetExistingTotalDeathDhankawadiWardCase() (int, error)
	GetExistingTotalDeathDholePatilWardCase() (int, error)
	GetExistingTotalDeathHadapsarMundhwaWardCase() (int, error)
	GetExistingTotalDeathKasbaWardCase() (int, error)
	GetExistingTotalDeathKondhwaWardCase() (int, error)
	GetExistingTotalDeathKothrudWardCase() (int, error)
	GetExistingTotalDeathWadgoansheriWardCase() (int, error)
	GetExistingTotalDeathShivajiNagarWardCase() (int, error)
	GetExistingTotalDeathSinhagadRoadWardCase() (int, error)
	GetExistingTotalDeathWanawadiWardCase() (int, error)
	GetExistingTotalDeathWarjeWardCase() (int, error)
	GetExistingTotalDeathYerwadaWardCase() (int, error)
	GetExistingTotalDeathPuneRuralWardCase() (int, error)
	GetExistingTotalDeathPuneCantonmentWardCase() (int, error)

	GetAundhBannerWardCaseByCreateDate(createDate string, endDate string) ([]database.AundhBaner, error)
	GetBhawaniPethWardByCreateDate(createDate string, endDate string) ([]database.BhawaniPeth, error)
	GetBibwewadiWardByCreateDate(createDate string, endDate string) ([]database.Bibwewadi, error)
	GetDhankawadiWardCaseByCreateDate(createDate string, endDate string) ([]database.DhankawadiSahakarnagar, error)
	GetDholePatilWardCaseByCreateDate(createDate string, endDate string) ([]database.DholePatil, error)
	GetHadapsarMundhwaWardCaseByCreateDate(createDate, endDate string) ([]database.HadapsarMundhwa, error)
	GetKasbaWardCaseByCreateDate(createDate string, endDate string) ([]database.KasbaVisharambaghwada, error)
	GetKondhwaWardCaseByCreateDate(createDate string, endDate string) ([]database.KondhwaYewalewadi, error)
	GetKothrudWardCaseByCreateDate(createDate string, endDate string) ([]database.KothrudBavdhan, error)
	GetWadgoansheriWardCaseByCreateDate(createDate string, endDate string) ([]database.NagarRoadWadgoansheri, error)
	GetShivajiNagarWardCaseByCreateDate(createDate string, endDate string) ([]database.ShivajiNagarGholeRoad, error)
	GetSinhagadRoadWardCaseByCreateDate(createDate string, endDate string) ([]database.SinghagadRoad, error)
	GetWanawadiWardCaseByCreateDate(createDate string, endDate string) ([]database.WanawadiRamtekdi, error)
	GetWarjeWardCaseByCreateDate(createDate string, endDate string) ([]database.WarjeKarveNagar, error)
	GetYerwadaWardCaseByCreateDate(createDate string, endDate string) ([]database.YerwadaKalasDhanori, error)
	GetPuneRuralWardCaseByCreateDate(createDate string, endDate string) ([]database.PuneRural, error)
	GetPuneCantonmentWardCaseByCreateDate(createDate string, endDate string) ([]database.PuneCantonment, error)

	//District
	AddDistrictCaseSummary(ward *domain.DistrictCase) (*database.DistrictCase, error)
	GetDistrictSummaryByCreateDate(createDate string, endDate string) ([]database.DistrictCase, error)
	AddDistrictDetail(districtDetail *domain.DistrictDetail) (*database.DistrictDetail, error)
	GetDistrictDetailByCreateDate(createDate string, endDate string) ([]database.DistrictDetail, error)

	//Patient Summary
	AddPatientSummary(patientSummary *domain.PatientSummary) (*database.PatientSummary, error)
	AddPatientSummaryDB(DBPatientSummary *database.PatientSummary) (*database.PatientSummary, error)
	GetPatientSummaryByCreateDate(createDate string, endDate string) ([]database.PatientSummary, error)
	GetPatientSummaryByDate(date string) (database.PatientSummary, error)

	//News
	AddNews(news *domain.News) (*database.News, error)
	GetNews(createDate string, endDate string) ([]database.News, error)
	GetDistrictSummaryByDate(date string) (database.DistrictCase, error)
}
