package service

import (
	"covid19/config"
	"covid19/config/domain"
	"covid19/infra/database"
	u "covid19/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
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

//AddWardCase add ward case
func (sh *Interactor) AddWardCase(wardCase *domain.WardCase) (*domain.WardCase, error) {
	var newWardCase domain.WardCase
	defer sh.GenerateWardCasesJSON()
	switch wardCase.Code {
	case "AB":
		{
			var aundhBanner domain.AundhBaner
			u.Copy(&aundhBanner, wardCase)
			dbAundhBannerCase, err := sh.Db.AddAundhBannerWardCase(&aundhBanner)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbAundhBannerCase)
			return &newWardCase, nil
		}
	case "BP":
		{
			var bhawaniPeth domain.BhawaniPeth
			u.Copy(&bhawaniPeth, wardCase)
			dbBhawaniPethCase, err := sh.Db.AddBhawaniPethWardCase(&bhawaniPeth)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbBhawaniPethCase)
			return &newWardCase, nil
		}
	case "BW":
		{
			var bibwewadi domain.Bibwewadi
			u.Copy(&bibwewadi, wardCase)
			dbBhawaniPethCase, err := sh.Db.AddBibwewadiWardCase(&bibwewadi)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbBhawaniPethCase)
			return &newWardCase, nil
		}
	case "DS":
		{
			var dhankawadiSahakarnagar domain.DhankawadiSahakarnagar
			u.Copy(&dhankawadiSahakarnagar, wardCase)
			dbdhankawadiSahakarnagarCase, err := sh.Db.AddDhankawadiSahakarnagarWardCase(&dhankawadiSahakarnagar)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbdhankawadiSahakarnagarCase)
			return &newWardCase, nil
		}
	case "DP":
		{
			var dholePatil domain.DholePatil
			u.Copy(&dholePatil, wardCase)
			dbdholePatilCase, err := sh.Db.AddDholePatilWardCase(&dholePatil)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbdholePatilCase)
			return &newWardCase, nil

		}
	case "HM":
		{
			var hadapsarMundhwa domain.HadapsarMundhwa
			u.Copy(&hadapsarMundhwa, wardCase)
			dbhadapsarMundhwaCase, err := sh.Db.AddHadapsarMundhwaWardCase(&hadapsarMundhwa)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbhadapsarMundhwaCase)
			return &newWardCase, nil

		}
	case "KV":
		{
			var kasbaVisharambaghwada domain.KasbaVisharambaghwada
			u.Copy(&kasbaVisharambaghwada, wardCase)
			dbkasbaVisharambaghwadaCase, err := sh.Db.AddKasbaVisharambaghwadaWardCase(&kasbaVisharambaghwada)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbkasbaVisharambaghwadaCase)
			return &newWardCase, nil

		}
	case "KY":
		{
			var kondhwaYewalewadi domain.KondhwaYewalewadi
			u.Copy(&kondhwaYewalewadi, wardCase)
			dbkondhwaYewalewadiCase, err := sh.Db.AddKondhwaYewalewadiWardCase(&kondhwaYewalewadi)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbkondhwaYewalewadiCase)
			return &newWardCase, nil

		}
	case "KB":
		{
			var kothrudBavdhandomain domain.KothrudBavdhan
			u.Copy(&kothrudBavdhandomain, wardCase)
			dbkothrudBavdhandomainCase, err := sh.Db.AddKothrudBavdhanWardCase(&kothrudBavdhandomain)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbkothrudBavdhandomainCase)
			return &newWardCase, nil

		}
	case "NW":
		{
			var nagarRoadWadgoansheri domain.NagarRoadWadgoansheri
			u.Copy(&nagarRoadWadgoansheri, wardCase)
			dbnagarRoadWadgoansheriCase, err := sh.Db.AddNagarRoadWadgoansheriWardCase(&nagarRoadWadgoansheri)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbnagarRoadWadgoansheriCase)
			return &newWardCase, nil

		}
	case "SN":
		{
			var shivajiNagarGholeRoad domain.ShivajiNagarGholeRoad
			u.Copy(&shivajiNagarGholeRoad, wardCase)
			dbshivajiNagarGholeRoadCase, err := sh.Db.AddShivajiNagarGholeRoadWardCase(&shivajiNagarGholeRoad)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbshivajiNagarGholeRoadCase)
			return &newWardCase, nil

		}
	case "SR":
		{
			var singhagadRoad domain.SinghagadRoad
			u.Copy(&singhagadRoad, wardCase)
			dbsinghagadRoadCase, err := sh.Db.AddSinhagadRoadWardCase(&singhagadRoad)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbsinghagadRoadCase)
			return &newWardCase, nil

		}
	case "WR":
		{
			var wanawadiRamtekdi domain.WanawadiRamtekdi
			u.Copy(&wanawadiRamtekdi, wardCase)
			dbwanawadiRamtekdiCase, err := sh.Db.AddWanawadiRamtekdiWardCase(&wanawadiRamtekdi)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbwanawadiRamtekdiCase)
			return &newWardCase, nil

		}
	case "WK":
		{
			var warjeKarveNagar domain.WarjeKarveNagar
			u.Copy(&warjeKarveNagar, wardCase)
			dbwarjeKarveNagarCase, err := sh.Db.AddWarjeKarveNagarWardCase(&warjeKarveNagar)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbwarjeKarveNagarCase)
			return &newWardCase, nil

		}
	case "YKD":
		{
			var yerwadaKalasDhanori domain.YerwadaKalasDhanori
			u.Copy(&yerwadaKalasDhanori, wardCase)
			dbyerwadaKalasDhanoriCase, err := sh.Db.AddYerwadaKalasDhanoriWardCase(&yerwadaKalasDhanori)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbyerwadaKalasDhanoriCase)
			return &newWardCase, nil

		}
	case "PR":
		{
			var puneRural domain.PuneRural
			u.Copy(&puneRural, wardCase)
			dbpuneRuralCase, err := sh.Db.AddPuneRuralWardCase(&puneRural)
			if err != nil {
				return nil, err
			}
			u.Copy(&newWardCase, dbpuneRuralCase)
			return &newWardCase, nil
		}
	}
	return nil, errors.New("Invalid ward code")
}

//AddAundhBannerCases add aundh banner case
func (sh *Interactor) AddAundhBannerCases(aundhBanner *domain.AundhBaner) (*database.AundhBaner, error) {
	if sh.Db.WardExists(aundhBanner.Code) {
		dbAundhBannerCase, err := sh.Db.AddAundhBannerWardCase(aundhBanner)
		if err != nil {
			return nil, err
		}
		return dbAundhBannerCase, nil
	}
	return nil, errors.New("invalid ward code")
}

//GetWardDetailsByCreateDateAndCode get ward details by create date and code
func (sh *Interactor) GetWardDetailsByCreateDateAndCode(createDate string, endDate string, code string) ([]*domain.WardCase, error) {
	wardCases := make([]*domain.WardCase, 0)
	var newWardCase domain.WardCase
	if sh.Db.WardExists(code) {
		switch code {
		case "AB":
			{
				dbAundhBannerCase, err := sh.Db.GetAundhBannerWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbAundhBannerCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Aundh - Baner"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "BP":
			{
				dbbhawaniPethWardCase, err := sh.Db.GetBhawaniPethWardByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbbhawaniPethWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Bhawani Peth"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "BW":
			{
				dbbibwewadiWardCase, err := sh.Db.GetBibwewadiWardByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbbibwewadiWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Bibwewadi"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "DS":
			{
				dbdhankawadiWard, err := sh.Db.GetDhankawadiWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbdhankawadiWard {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Dhankawadi - Sahakarnagar"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "DP":
			{
				dbdholePatilWard, err := sh.Db.GetDholePatilWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbdholePatilWard {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Dhole Patil"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "HM":
			{
				dbhadapsarMundhwaWardCase, err := sh.Db.GetHadapsarMundhwaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbhadapsarMundhwaWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Hadapsar Mundhwa"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "KV":
			{
				dbkasbaWardCase, err := sh.Db.GetKasbaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbkasbaWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Kasba - Visharambaghwada"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "KY":
			{
				dbkondhwaWardCase, err := sh.Db.GetKondhwaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbkondhwaWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Kondhwa - Yewalewadi"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "KB":
			{
				dbkothrudWardCase, err := sh.Db.GetKothrudWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbkothrudWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Kothrud - Bavdhan"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "NW":
			{
				dbwadgoansheriWardCase, err := sh.Db.GetWadgoansheriWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbwadgoansheriWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Nagar Road - Wadgoansheri"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "SN":
			{
				dbshivajiNagarWardCase, err := sh.Db.GetShivajiNagarWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbshivajiNagarWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Shivaji Nagar - Ghole Road"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "SR":
			{
				dbsinhagadRoadWardCase, err := sh.Db.GetSinhagadRoadWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbsinhagadRoadWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Singhagad Road"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "WR":
			{
				dbwanawadiWardCase, err := sh.Db.GetWanawadiWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbwanawadiWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Wanawadi - Ramtekdi"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "WK":
			{
				dbwarjeWardCase, err := sh.Db.GetWarjeWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbwarjeWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Warje - Karve Nagar"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "YKD":
			{
				dbyerwadaWardCase, err := sh.Db.GetYerwadaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbyerwadaWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Yerwada - Kalas - Dhanori"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		case "PR":
			{
				dbpuneRuralWardCase, err := sh.Db.GetPuneRuralWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbpuneRuralWardCase {
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Pune Rural"
					wardCases = append(wardCases, &newWardCase)
				}
			}
		}
		return wardCases, nil
	}
	return nil, errors.New("invalid ward code")
}

//GetAllWardDetailsByCreateDate get all ward details by create date
func (sh *Interactor) GetAllWardDetailsByCreateDate(createDate string, endDate string) ([]*domain.WardCase, error) {
	wardCases := make([]*domain.WardCase, 0)

	dbAundhBannerCase, err := sh.Db.GetAundhBannerWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbAundhBannerCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Aundh - Baner"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbbhawaniPethWardCase, err := sh.Db.GetBhawaniPethWardByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbbhawaniPethWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Bhawani Peth"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbbibwewadiWardCase, err := sh.Db.GetBibwewadiWardByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbbibwewadiWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Bibwewadi"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbdhankawadiWard, err := sh.Db.GetDhankawadiWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbdhankawadiWard {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Dhankawadi - Sahakarnagar"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbdholePatilWard, err := sh.Db.GetDholePatilWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbdholePatilWard {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Dhole Patil"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbhadapsarMundhwaWardCase, err := sh.Db.GetHadapsarMundhwaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbhadapsarMundhwaWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Hadapsar Mundhwa"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbkasbaWardCase, err := sh.Db.GetKasbaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbkasbaWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Kasba - Visharambaghwada"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbkondhwaWardCase, err := sh.Db.GetKondhwaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbkondhwaWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Kondhwa - Yewalewadi"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbkothrudWardCase, err := sh.Db.GetKothrudWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbkothrudWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Kothrud - Bavdhan"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbwadgoansheriWardCase, err := sh.Db.GetWadgoansheriWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbwadgoansheriWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Nagar Road - Wadgoansheri"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbshivajiNagarWardCase, err := sh.Db.GetShivajiNagarWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbshivajiNagarWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Shivaji Nagar - Ghole Road"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbsinhagadRoadWardCase, err := sh.Db.GetSinhagadRoadWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbsinhagadRoadWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Singhagad Road"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbwanawadiWardCase, err := sh.Db.GetWanawadiWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbwanawadiWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Wanawadi - Ramtekdi"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbwarjeWardCase, err := sh.Db.GetWarjeWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbwarjeWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Warje - Karve Nagar"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbyerwadaWardCase, err := sh.Db.GetYerwadaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbyerwadaWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Yerwada - Kalas - Dhanori"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	dbpuneRuralWardCase, err := sh.Db.GetPuneRuralWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		for _, wardCase := range dbpuneRuralWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Pune Rural"
			wardCases = append(wardCases, &newWardCase)
		}
	}
	return wardCases, nil
}

//GenerateWardCasesJSON generate district summary json file
func (sh *Interactor) GenerateWardCasesJSON() error {
	wardCases, err := sh.GetAllWardDetailsByCreateDate("2019-12-01", "2040-12-31")
	if err != nil {
		return err
	}
	file, _ := json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-all-cases.json", file, 0644)
	return nil
}
