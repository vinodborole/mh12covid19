package service

import (
	"covid19/config/domain"
	"covid19/infra/database"
	u "covid19/utils"
	"errors"
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
				u.Copy(&newWardCase, dbAundhBannerCase)
				newWardCase.Date = dbAundhBannerCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Aundh - Baner"
				wardCases = append(wardCases, &newWardCase)
			}
		case "BP":
			{
				dbbhawaniPethWardCase, err := sh.Db.GetBhawaniPethWardByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbbhawaniPethWardCase)
				newWardCase.Date = dbbhawaniPethWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Bhawani Peth"
				wardCases = append(wardCases, &newWardCase)
			}
		case "BW":
			{
				dbbibwewadiWardCase, err := sh.Db.GetBibwewadiWardByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbbibwewadiWardCase)
				newWardCase.Date = dbbibwewadiWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Bibwewadi"
				wardCases = append(wardCases, &newWardCase)
			}
		case "DS":
			{
				dbdhankawadiWard, err := sh.Db.GetDhankawadiWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbdhankawadiWard)
				newWardCase.Date = dbdhankawadiWard.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Dhankawadi - Sahakarnagar"
				wardCases = append(wardCases, &newWardCase)
			}
		case "DP":
			{
				dbdholePatilWard, err := sh.Db.GetDholePatilWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbdholePatilWard)
				newWardCase.Date = dbdholePatilWard.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Dhole Patil"
				wardCases = append(wardCases, &newWardCase)
			}
		case "HM":
			{
				dbhadapsarMundhwaWardCase, err := sh.Db.GetHadapsarMundhwaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbhadapsarMundhwaWardCase)
				newWardCase.Date = dbhadapsarMundhwaWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Hadapsar Mundhwa"
				wardCases = append(wardCases, &newWardCase)
			}
		case "KV":
			{
				dbkasbaWardCase, err := sh.Db.GetKasbaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbkasbaWardCase)
				newWardCase.Date = dbkasbaWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Kasba - Visharambaghwada"
				wardCases = append(wardCases, &newWardCase)
			}
		case "KY":
			{
				dbkondhwaWardCase, err := sh.Db.GetKondhwaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbkondhwaWardCase)
				newWardCase.Date = dbkondhwaWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Kondhwa - Yewalewadi"
				wardCases = append(wardCases, &newWardCase)
			}
		case "KB":
			{
				dbkothrudWardCase, err := sh.Db.GetKothrudWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbkothrudWardCase)
				newWardCase.Date = dbkothrudWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Kothrud - Bavdhan"
				wardCases = append(wardCases, &newWardCase)
			}
		case "NW":
			{
				dbwadgoansheriWardCase, err := sh.Db.GetWadgoansheriWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbwadgoansheriWardCase)
				newWardCase.Date = dbwadgoansheriWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Nagar Road - Wadgoansheri"
				wardCases = append(wardCases, &newWardCase)
			}
		case "SN":
			{
				dbshivajiNagarWardCase, err := sh.Db.GetShivajiNagarWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbshivajiNagarWardCase)
				newWardCase.Date = dbshivajiNagarWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Shivaji Nagar - Ghole Road"
				wardCases = append(wardCases, &newWardCase)
			}
		case "SR":
			{
				dbsinhagadRoadWardCase, err := sh.Db.GetSinhagadRoadWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbsinhagadRoadWardCase)
				newWardCase.Date = dbsinhagadRoadWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Singhagad Road"
				wardCases = append(wardCases, &newWardCase)
			}
		case "WR":
			{
				dbwanawadiWardCase, err := sh.Db.GetWanawadiWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbwanawadiWardCase)
				newWardCase.Date = dbwanawadiWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Wanawadi - Ramtekdi"
				wardCases = append(wardCases, &newWardCase)
			}
		case "WK":
			{
				dbwarjeWardCase, err := sh.Db.GetWarjeWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbwarjeWardCase)
				newWardCase.Date = dbwarjeWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Warje - Karve Nagar"
				wardCases = append(wardCases, &newWardCase)
			}
		case "YKD":
			{
				dbyerwadaWardCase, err := sh.Db.GetYerwadaWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbyerwadaWardCase)
				newWardCase.Date = dbyerwadaWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Yerwada - Kalas - Dhanori"
				wardCases = append(wardCases, &newWardCase)
			}
		case "PR":
			{
				dbpuneRuralWardCase, err := sh.Db.GetPuneRuralWardCaseByCreateDate(createDate, endDate)
				if err != nil {
					return nil, err
				}
				u.Copy(&newWardCase, dbpuneRuralWardCase)
				newWardCase.Date = dbpuneRuralWardCase.CreatedAt.Format("2006-01-02")
				newWardCase.Name = "Pune Rural"
				wardCases = append(wardCases, &newWardCase)
			}
		}
		return wardCases, nil
	}
	return nil, errors.New("invalid ward code")
}

//GetAllWardDetailsByCreateDate get all ward details by create date
func (sh *Interactor) GetAllWardDetailsByCreateDate(createDate string, endDate string) ([]*domain.WardCase, error) {
	wardCases := make([]*domain.WardCase, 0)
	var newWardCase domain.WardCase
	dbAundhBannerCase, err := sh.Db.GetAundhBannerWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbAundhBannerCase)
		newWardCase.Date = dbAundhBannerCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Aundh - Baner"
		wardCases = append(wardCases, &newWardCase)
	}
	dbbhawaniPethWardCase, err := sh.Db.GetBhawaniPethWardByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbbhawaniPethWardCase)
		newWardCase.Date = dbbhawaniPethWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Bhawani Peth"
		wardCases = append(wardCases, &newWardCase)
	}
	dbbibwewadiWardCase, err := sh.Db.GetBibwewadiWardByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbbibwewadiWardCase)
		newWardCase.Date = dbbibwewadiWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Bibwewadi"
		wardCases = append(wardCases, &newWardCase)
	}
	dbdhankawadiWard, err := sh.Db.GetDhankawadiWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbdhankawadiWard)
		newWardCase.Date = dbdhankawadiWard.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Dhankawadi - Sahakarnagar"
		wardCases = append(wardCases, &newWardCase)
	}
	dbdholePatilWard, err := sh.Db.GetDholePatilWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbdholePatilWard)
		newWardCase.Date = dbdholePatilWard.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Dhole Patil"
		wardCases = append(wardCases, &newWardCase)
	}
	dbhadapsarMundhwaWardCase, err := sh.Db.GetHadapsarMundhwaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbhadapsarMundhwaWardCase)
		newWardCase.Date = dbhadapsarMundhwaWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Hadapsar Mundhwa"
		wardCases = append(wardCases, &newWardCase)
	}
	dbkasbaWardCase, err := sh.Db.GetKasbaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbkasbaWardCase)
		newWardCase.Date = dbkasbaWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Kasba - Visharambaghwada"
		wardCases = append(wardCases, &newWardCase)
	}
	dbkondhwaWardCase, err := sh.Db.GetKondhwaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbkondhwaWardCase)
		newWardCase.Date = dbkondhwaWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Kondhwa - Yewalewadi"
		wardCases = append(wardCases, &newWardCase)
	}
	dbkothrudWardCase, err := sh.Db.GetKothrudWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbkothrudWardCase)
		newWardCase.Date = dbkothrudWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Kothrud - Bavdhan"
		wardCases = append(wardCases, &newWardCase)
	}
	dbwadgoansheriWardCase, err := sh.Db.GetWadgoansheriWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbwadgoansheriWardCase)
		newWardCase.Date = dbwadgoansheriWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Nagar Road - Wadgoansheri"
		wardCases = append(wardCases, &newWardCase)
	}
	dbshivajiNagarWardCase, err := sh.Db.GetShivajiNagarWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbshivajiNagarWardCase)
		newWardCase.Date = dbshivajiNagarWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Shivaji Nagar - Ghole Road"
		wardCases = append(wardCases, &newWardCase)
	}
	dbsinhagadRoadWardCase, err := sh.Db.GetSinhagadRoadWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbsinhagadRoadWardCase)
		newWardCase.Date = dbsinhagadRoadWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Singhagad Road"
		wardCases = append(wardCases, &newWardCase)
	}
	dbwanawadiWardCase, err := sh.Db.GetWanawadiWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbwanawadiWardCase)
		newWardCase.Date = dbwanawadiWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Wanawadi - Ramtekdi"
		wardCases = append(wardCases, &newWardCase)
	}
	dbwarjeWardCase, err := sh.Db.GetWarjeWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbwarjeWardCase)
		newWardCase.Date = dbwarjeWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Warje - Karve Nagar"
		wardCases = append(wardCases, &newWardCase)
	}
	dbyerwadaWardCase, err := sh.Db.GetYerwadaWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbyerwadaWardCase)
		newWardCase.Date = dbyerwadaWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Yerwada - Kalas - Dhanori"
		wardCases = append(wardCases, &newWardCase)
	}
	dbpuneRuralWardCase, err := sh.Db.GetPuneRuralWardCaseByCreateDate(createDate, endDate)
	if err == nil {
		u.Copy(&newWardCase, dbpuneRuralWardCase)
		newWardCase.Date = dbpuneRuralWardCase.CreatedAt.Format("2006-01-02")
		newWardCase.Name = "Pune Rural"
		wardCases = append(wardCases, &newWardCase)
	}
	return wardCases, nil
}
