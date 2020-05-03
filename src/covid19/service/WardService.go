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
	defer sh.GenerateMetaJSON()
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
func (sh *Interactor) GetWardDetailsByCreateDateAndCode(createDate string, endDate string, code string) ([]domain.WardDetail, error) {
	wardDetails := make([]domain.WardDetail, 0)
	if sh.Db.WardExists(code) {
		switch code {
		case "AB":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbAundhBannerCase, err := sh.Db.GetAundhBannerWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbAundhBannerCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Aundh - Baner"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Aundh - Baner"
				wardDetail.Code = "AB"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "BP":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbbhawaniPethWardCase, err := sh.Db.GetBhawaniPethWardByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbbhawaniPethWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Bhawani Peth"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Bhawani Peth"
				wardDetail.Code = "BP"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "BW":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbbibwewadiWardCase, err := sh.Db.GetBibwewadiWardByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbbibwewadiWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Bibwewadi"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Bibwewadi"
				wardDetail.Code = "BW"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "DS":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbdhankawadiWard, err := sh.Db.GetDhankawadiWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbdhankawadiWard {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Dhankawadi - Sahakarnagar"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Dhankawadi - Sahakarnagar"
				wardDetail.Code = "DS"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "DP":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbdholePatilWard, err := sh.Db.GetDholePatilWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbdholePatilWard {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Dhole Patil"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Dhole Patil"
				wardDetail.Code = "DP"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "HM":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbhadapsarMundhwaWardCase, err := sh.Db.GetHadapsarMundhwaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbhadapsarMundhwaWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Hadapsar Mundhwa"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Hadapsar Mundhwa"
				wardDetail.Code = "HM"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "KV":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbkasbaWardCase, err := sh.Db.GetKasbaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbkasbaWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Kasba - Visharambaghwada"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Kasba - Visharambaghwada"
				wardDetail.Code = "KV"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "KY":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbkondhwaWardCase, err := sh.Db.GetKondhwaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbkondhwaWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Kondhwa - Yewalewadi"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Kondhwa - Yewalewadi"
				wardDetail.Code = "KY"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "KB":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbkothrudWardCase, err := sh.Db.GetKothrudWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbkothrudWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Kothrud - Bavdhan"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Kothrud - Bavdhan"
				wardDetail.Code = "KB"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "NW":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbwadgoansheriWardCase, err := sh.Db.GetWadgoansheriWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbwadgoansheriWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Nagar Road - Wadgoansheri"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Nagar Road - Wadgoansheri"
				wardDetail.Code = "NW"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "SN":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbshivajiNagarWardCase, err := sh.Db.GetShivajiNagarWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbshivajiNagarWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Shivaji Nagar - Ghole Road"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Shivaji Nagar - Ghole Road"
				wardDetail.Code = "SN"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "SR":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbsinhagadRoadWardCase, err := sh.Db.GetSinhagadRoadWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbsinhagadRoadWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Singhagad Road"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Singhagad Road"
				wardDetail.Code = "SR"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "WR":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbwanawadiWardCase, err := sh.Db.GetWanawadiWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbwanawadiWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Wanawadi - Ramtekdi"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Wanawadi - Ramtekdi"
				wardDetail.Code = "WR"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "WK":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbwarjeWardCase, err := sh.Db.GetWarjeWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbwarjeWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Warje - Karve Nagar"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Warje - Karve Nagar"
				wardDetail.Code = "WK"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "YKD":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbyerwadaWardCase, err := sh.Db.GetYerwadaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbyerwadaWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Yerwada - Kalas - Dhanori"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Yerwada - Kalas - Dhanori"
				wardDetail.Code = "YKD"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		case "PR":
			{
				var wardDetail domain.WardDetail
				wardCases := make([]domain.WardCase, 0)
				var total int
				dbpuneRuralWardCase, err := sh.Db.GetPuneRuralWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				for _, wardCase := range dbpuneRuralWardCase {
					var newWardCase domain.WardCase
					u.Copy(&newWardCase, wardCase)
					newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
					newWardCase.Name = "Pune Rural"
					total = total + wardCase.Active
					wardCases = append(wardCases, newWardCase)
				}
				wardDetail.Name = "Pune Rural"
				wardDetail.Code = "PR"
				wardDetail.WardCases = wardCases
				wardDetail.TotalActive = total
				wardDetails = append(wardDetails, wardDetail)
			}
		}
		return wardDetails, nil
	}
	return nil, errors.New("invalid ward code")
}

//GetAllWardDetailsByCreateDate get all ward details by create date
func (sh *Interactor) GetAllWardDetailsByCreateDate(createDate string, endDate string) ([]domain.WardDetail, error) {
	wardDetails := make([]domain.WardDetail, 0)

	dbAundhBannerCase, err := sh.Db.GetAundhBannerWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbAundhBannerCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Aundh - Baner"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Aundh - Baner"
		wardDetail.Code = "AB"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}

	dbbhawaniPethWardCase, err := sh.Db.GetBhawaniPethWardByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbbhawaniPethWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Bhawani Peth"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Bhawani Peth"
		wardDetail.Code = "BP"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}

	dbbibwewadiWardCase, err := sh.Db.GetBibwewadiWardByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbbibwewadiWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Bibwewadi"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Bibwewadi"
		wardDetail.Code = "BW"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbdhankawadiWard, err := sh.Db.GetDhankawadiWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbdhankawadiWard {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Dhankawadi - Sahakarnagar"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Dhankawadi - Sahakarnagar"
		wardDetail.Code = "DS"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbdholePatilWard, err := sh.Db.GetDholePatilWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbdholePatilWard {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Dhole Patil"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Dhole Patil"
		wardDetail.Code = "DP"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbhadapsarMundhwaWardCase, err := sh.Db.GetHadapsarMundhwaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbhadapsarMundhwaWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Hadapsar Mundhwa"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Hadapsar Mundhwa"
		wardDetail.Code = "HM"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbkasbaWardCase, err := sh.Db.GetKasbaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbkasbaWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Kasba - Visharambaghwada"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Kasba - Visharambaghwada"
		wardDetail.Code = "KV"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbkondhwaWardCase, err := sh.Db.GetKondhwaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbkondhwaWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Kondhwa - Yewalewadi"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Kondhwa - Yewalewadi"
		wardDetail.Code = "KY"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbkothrudWardCase, err := sh.Db.GetKothrudWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbkothrudWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Kothrud - Bavdhan"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Kothrud - Bavdhan"
		wardDetail.Code = "KB"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbwadgoansheriWardCase, err := sh.Db.GetWadgoansheriWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbwadgoansheriWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Nagar Road - Wadgoansheri"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Nagar Road - Wadgoansheri"
		wardDetail.Code = "NW"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbshivajiNagarWardCase, err := sh.Db.GetShivajiNagarWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbshivajiNagarWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Shivaji Nagar - Ghole Road"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Shivaji Nagar - Ghole Road"
		wardDetail.Code = "SN"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbsinhagadRoadWardCase, err := sh.Db.GetSinhagadRoadWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbsinhagadRoadWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Singhagad Road"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Singhagad Road"
		wardDetail.Code = "SR"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbwanawadiWardCase, err := sh.Db.GetWanawadiWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbwanawadiWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Wanawadi - Ramtekdi"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Wanawadi - Ramtekdi"
		wardDetail.Code = "WR"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbwarjeWardCase, err := sh.Db.GetWarjeWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbwarjeWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Warje - Karve Nagar"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Warje - Karve Nagar"
		wardDetail.Code = "WK"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbyerwadaWardCase, err := sh.Db.GetYerwadaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbyerwadaWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Yerwada - Kalas - Dhanori"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Yerwada - Kalas - Dhanori"
		wardDetail.Code = "YKD"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	dbpuneRuralWardCase, err := sh.Db.GetPuneRuralWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		var wardDetail domain.WardDetail
		var total int
		wardCases := make([]domain.WardCase, 0)
		for _, wardCase := range dbpuneRuralWardCase {
			var newWardCase domain.WardCase
			u.Copy(&newWardCase, wardCase)
			newWardCase.Date = wardCase.CreatedAt.Format("2006-01-02")
			newWardCase.Name = "Pune Rural"
			total = total + wardCase.Active
			wardCases = append(wardCases, newWardCase)
		}
		wardDetail.Name = "Pune Rural"
		wardDetail.Code = "PR"
		wardDetail.WardCases = wardCases
		wardDetail.TotalActive = total
		wardDetails = append(wardDetails, wardDetail)
	}
	return wardDetails, nil
}

//GenerateWardCasesJSON generate district summary json file
func (sh *Interactor) GenerateWardCasesJSON() error {
	wardCases, err := sh.GetAllWardDetailsByCreateDate("2019-12-01", "2040-12-31")
	if err != nil {
		return err
	}
	file, _ := json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-all-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "AB")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-ab-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "BP")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-bp-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "BW")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-bw-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "DS")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-ds-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "DP")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-dp-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "HM")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-hm-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "KV")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-kv-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "KY")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-ky-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "KB")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-kb-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "NW")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-nw-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "SN")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-sn-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "SR")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-sr-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "WR")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-wr-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "WK")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-wk-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "YKD")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-ykd-cases.json", file, 0644)

	wardCases, err = sh.GetWardDetailsByCreateDateAndCode("2019-12-01", "2040-12-31", "PR")
	if err != nil {
		return err
	}
	file, _ = json.MarshalIndent(wardCases, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"ward-pr-cases.json", file, 0644)

	return nil
}
