package gateway

import (
	"covid19/config/domain"
	"covid19/infra/database"
	u "covid19/utils"
	"errors"
	"time"
)

//CreateWard create account
func (dbRepo *DatabaseRepository) CreateWard(ward *domain.Ward) (*database.Ward, error) {
	var DBWard database.Ward
	u.Copy(&DBWard, ward)
	return dbRepo.CreateWardDB(&DBWard)
}

//CreateWardDB create ward db
func (dbRepo *DatabaseRepository) CreateWardDB(DBWard *database.Ward) (*database.Ward, error) {
	err := dbRepo.GetDBHandle().Create(&DBWard).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to create ward,connection error: " + err.Error())
	}
	return DBWard, nil
}

//GetWards get wards
func (dbRepo *DatabaseRepository) GetWards(page string) *u.Data {
	var wards []database.Ward
	// need this limit somewhere in constant/env file
	var limit string = "100"
	orderBy := []string{"ID"}
	paginator := u.Paginator{DB: dbRepo.GetDBHandle(), OrderBy: orderBy, Page: page, PerPage: limit}
	data := paginator.Paginate(&wards)
	return data
}

//GetWardByCode get ward by code
func (dbRepo *DatabaseRepository) GetWardByCode(code string) (*database.Ward, error) {
	var DBWard database.Ward
	err := dbRepo.GetDBHandle().Model(domain.Ward{}).Where("code = ?", code).Find(&DBWard).Error
	return &DBWard, err
}

//WardExists ward exists
func (dbRepo *DatabaseRepository) WardExists(code string) bool {
	var DBWard database.Ward
	if dbRepo.GetDBHandle().Where("code = ?", code).First(&DBWard).RecordNotFound() {
		return false
	}
	return true
}

//AddAundhBannerWardCase add case to aundh banner ward
func (dbRepo *DatabaseRepository) AddAundhBannerWardCase(ward *domain.AundhBaner) (*database.AundhBaner, error) {
	var DBAundhBannerWard database.AundhBaner
	u.Copy(&DBAundhBannerWard, ward)
	if len(ward.Date) > 0 {
		DBAundhBannerWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddAundhBannerWardDBCase(&DBAundhBannerWard)
}

//AddAundhBannerWardDBCase add case to aundh banner ward db
func (dbRepo *DatabaseRepository) AddAundhBannerWardDBCase(DBAundhBannerWard *database.AundhBaner) (*database.AundhBaner, error) {
	err := dbRepo.GetDBHandle().Create(&DBAundhBannerWard).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add aundh banner ward case,connection error: " + err.Error())
	}
	return DBAundhBannerWard, nil
}

//AddBhawaniPethWardCase add case to bhawani peth ward
func (dbRepo *DatabaseRepository) AddBhawaniPethWardCase(ward *domain.BhawaniPeth) (*database.BhawaniPeth, error) {
	var DBWard database.BhawaniPeth
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddBhawaniPethWardDBCase(&DBWard)
}

//AddBhawaniPethWardDBCase add case to bhawani peth ward db
func (dbRepo *DatabaseRepository) AddBhawaniPethWardDBCase(ward *database.BhawaniPeth) (*database.BhawaniPeth, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add bhawani peth ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddBibwewadiWardCase add case to bibwewadi ward
func (dbRepo *DatabaseRepository) AddBibwewadiWardCase(ward *domain.Bibwewadi) (*database.Bibwewadi, error) {
	var DBWard database.Bibwewadi
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddBibwewadiWardDBCase(&DBWard)
}

//AddBibwewadiWardDBCase add case to bibwewadi ward db
func (dbRepo *DatabaseRepository) AddBibwewadiWardDBCase(ward *database.Bibwewadi) (*database.Bibwewadi, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add bibwewadi ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddDhankawadiSahakarnagarWardCase add case to Dhankawadi Sahakarnagar ward
func (dbRepo *DatabaseRepository) AddDhankawadiSahakarnagarWardCase(ward *domain.DhankawadiSahakarnagar) (*database.DhankawadiSahakarnagar, error) {
	var DBWard database.DhankawadiSahakarnagar
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddDhankawadiSahakarnagarWardDBCase(&DBWard)
}

//AddDhankawadiSahakarnagarWardDBCase add case to Dhankawadi Sahakarnagar ward db
func (dbRepo *DatabaseRepository) AddDhankawadiSahakarnagarWardDBCase(ward *database.DhankawadiSahakarnagar) (*database.DhankawadiSahakarnagar, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Dhankawadi Sahakarnagar ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddDholePatilWardCase add case to Dhole Patil ward
func (dbRepo *DatabaseRepository) AddDholePatilWardCase(ward *domain.DholePatil) (*database.DholePatil, error) {
	var DBWard database.DholePatil
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddDholePatilWardDBCase(&DBWard)
}

//AddDholePatilWardDBCase add case to Dhole Patil ward db
func (dbRepo *DatabaseRepository) AddDholePatilWardDBCase(ward *database.DholePatil) (*database.DholePatil, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Dhole Patil ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddHadapsarMundhwaWardCase add case to Hadapsar Mundhwa ward
func (dbRepo *DatabaseRepository) AddHadapsarMundhwaWardCase(ward *domain.HadapsarMundhwa) (*database.HadapsarMundhwa, error) {
	var DBWard database.HadapsarMundhwa
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddHadapsarMundhwaWardDBCase(&DBWard)
}

//AddHadapsarMundhwaWardDBCase add case to Hadapsar Mundhwa ward db
func (dbRepo *DatabaseRepository) AddHadapsarMundhwaWardDBCase(ward *database.HadapsarMundhwa) (*database.HadapsarMundhwa, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Hadapsar Mundhwa ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddKasbaVisharambaghwadaWardCase add case to Kasba Visharambaghwada ward
func (dbRepo *DatabaseRepository) AddKasbaVisharambaghwadaWardCase(ward *domain.KasbaVisharambaghwada) (*database.KasbaVisharambaghwada, error) {
	var DBWard database.KasbaVisharambaghwada
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddKasbaVisharambaghwadaWardDBCase(&DBWard)
}

//AddKasbaVisharambaghwadaWardDBCase add case to Kasba Visharambaghwada ward db
func (dbRepo *DatabaseRepository) AddKasbaVisharambaghwadaWardDBCase(ward *database.KasbaVisharambaghwada) (*database.KasbaVisharambaghwada, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Kasba Visharambaghwada ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddKondhwaYewalewadiWardCase add case to Kondhwa Yewalewadi ward
func (dbRepo *DatabaseRepository) AddKondhwaYewalewadiWardCase(ward *domain.KondhwaYewalewadi) (*database.KondhwaYewalewadi, error) {
	var DBWard database.KondhwaYewalewadi
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddKondhwaYewalewadiWardDBCase(&DBWard)
}

//AddKondhwaYewalewadiWardDBCase add case to Kondhwa Yewalewadi ward db
func (dbRepo *DatabaseRepository) AddKondhwaYewalewadiWardDBCase(ward *database.KondhwaYewalewadi) (*database.KondhwaYewalewadi, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Kondhwa Yewalewadi ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddKothrudBavdhanWardCase add case to Kothrud Bavdhan ward
func (dbRepo *DatabaseRepository) AddKothrudBavdhanWardCase(ward *domain.KothrudBavdhan) (*database.KothrudBavdhan, error) {
	var DBWard database.KothrudBavdhan
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddKothrudBavdhanWardDBCase(&DBWard)
}

//AddKothrudBavdhanWardDBCase add case to Kothrud Bavdhan ward db
func (dbRepo *DatabaseRepository) AddKothrudBavdhanWardDBCase(ward *database.KothrudBavdhan) (*database.KothrudBavdhan, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Kothrud Bavdhan ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddNagarRoadWadgoansheriWardCase add case to Nagar Road Wadgoansheri ward
func (dbRepo *DatabaseRepository) AddNagarRoadWadgoansheriWardCase(ward *domain.NagarRoadWadgoansheri) (*database.NagarRoadWadgoansheri, error) {
	var DBWard database.NagarRoadWadgoansheri
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddNagarRoadWadgoansheriWardDBCase(&DBWard)
}

//AddNagarRoadWadgoansheriWardDBCase add case to Nagar Road Wadgoansheri ward db
func (dbRepo *DatabaseRepository) AddNagarRoadWadgoansheriWardDBCase(ward *database.NagarRoadWadgoansheri) (*database.NagarRoadWadgoansheri, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Nagar Road Wadgoansheri ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddShivajiNagarGholeRoadWardCase add case to Shivaji Nagar Ghole Road ward
func (dbRepo *DatabaseRepository) AddShivajiNagarGholeRoadWardCase(ward *domain.ShivajiNagarGholeRoad) (*database.ShivajiNagarGholeRoad, error) {
	var DBWard database.ShivajiNagarGholeRoad
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddShivajiNagarGholeRoadWardDBCase(&DBWard)
}

//AddShivajiNagarGholeRoadWardDBCase add case to Shivaji Nagar Ghole Road ward db
func (dbRepo *DatabaseRepository) AddShivajiNagarGholeRoadWardDBCase(ward *database.ShivajiNagarGholeRoad) (*database.ShivajiNagarGholeRoad, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Shivaji Nagar Ghole Road ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddSinhagadRoadWardCase add case to Sinhagad Road ward
func (dbRepo *DatabaseRepository) AddSinhagadRoadWardCase(ward *domain.SinghagadRoad) (*database.SinghagadRoad, error) {
	var DBWard database.SinghagadRoad
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddSinhagadRoadWardDBCase(&DBWard)
}

//AddSinhagadRoadWardDBCase add case to Sinhagad Road ward db
func (dbRepo *DatabaseRepository) AddSinhagadRoadWardDBCase(ward *database.SinghagadRoad) (*database.SinghagadRoad, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Sinhagad Road ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddWanawadiRamtekdiWardCase add case to Wanawadi Ramtekdi ward
func (dbRepo *DatabaseRepository) AddWanawadiRamtekdiWardCase(ward *domain.WanawadiRamtekdi) (*database.WanawadiRamtekdi, error) {
	var DBWard database.WanawadiRamtekdi
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddWanawadiRamtekdiWardDBCase(&DBWard)
}

//AddWanawadiRamtekdiWardDBCase add case to Wanawadi Ramtekdi ward db
func (dbRepo *DatabaseRepository) AddWanawadiRamtekdiWardDBCase(ward *database.WanawadiRamtekdi) (*database.WanawadiRamtekdi, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Wanawadi Ramtekdi ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddWarjeKarveNagarWardCase add case to Warje Karve Nagar ward
func (dbRepo *DatabaseRepository) AddWarjeKarveNagarWardCase(ward *domain.WarjeKarveNagar) (*database.WarjeKarveNagar, error) {
	var DBWard database.WarjeKarveNagar
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddWarjeKarveNagarWardDBCase(&DBWard)
}

//AddWarjeKarveNagarWardDBCase add case to Warje Karve Nagar ward db
func (dbRepo *DatabaseRepository) AddWarjeKarveNagarWardDBCase(ward *database.WarjeKarveNagar) (*database.WarjeKarveNagar, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Warje Karve Nagar ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddYerwadaKalasDhanoriWardCase add case to Yerwada Kalas Dhanori ward
func (dbRepo *DatabaseRepository) AddYerwadaKalasDhanoriWardCase(ward *domain.YerwadaKalasDhanori) (*database.YerwadaKalasDhanori, error) {
	var DBWard database.YerwadaKalasDhanori
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddYerwadaKalasDhanoriWardDBCase(&DBWard)
}

//AddYerwadaKalasDhanoriWardDBCase add case to Yerwada Kalas Dhanori ward db
func (dbRepo *DatabaseRepository) AddYerwadaKalasDhanoriWardDBCase(ward *database.YerwadaKalasDhanori) (*database.YerwadaKalasDhanori, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Yerwada Kalas Dhanori ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//AddPuneRuralWardCase add case to Pune Rural ward
func (dbRepo *DatabaseRepository) AddPuneRuralWardCase(ward *domain.PuneRural) (*database.PuneRural, error) {
	var DBWard database.PuneRural
	u.Copy(&DBWard, ward)
	if len(ward.Date) > 0 {
		DBWard.CreatedAt, _ = time.Parse("2006-01-02", ward.Date)
	}
	return dbRepo.AddPuneRuralWardDBCase(&DBWard)
}

//AddPuneRuralWardDBCase add case to Pune Rural ward db
func (dbRepo *DatabaseRepository) AddPuneRuralWardDBCase(ward *database.PuneRural) (*database.PuneRural, error) {
	err := dbRepo.GetDBHandle().Create(&ward).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add Pune Rural ward case,connection error: " + err.Error())
	}
	return ward, nil
}

//GetAundhBannerWardCaseByCreateDate get aundh banner ward cases by create date
func (dbRepo *DatabaseRepository) GetAundhBannerWardCaseByCreateDate(createDate string, endDate string) ([]database.AundhBaner, error) {
	var DBWardCase []database.AundhBaner
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetBhawaniPethWardByCreateDate get bhawani peth ward cases by create date
func (dbRepo *DatabaseRepository) GetBhawaniPethWardByCreateDate(createDate string, endDate string) ([]database.BhawaniPeth, error) {
	var DBWardCase []database.BhawaniPeth
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetBibwewadiWardByCreateDate get bibwewadi ward cases by create date
func (dbRepo *DatabaseRepository) GetBibwewadiWardByCreateDate(createDate string, endDate string) ([]database.Bibwewadi, error) {
	var DBWardCase []database.Bibwewadi
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetDhankawadiWardCaseByCreateDate get Dhankawadi Ward cases by create date
func (dbRepo *DatabaseRepository) GetDhankawadiWardCaseByCreateDate(createDate string, endDate string) ([]database.DhankawadiSahakarnagar, error) {
	var DBWardCase []database.DhankawadiSahakarnagar
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetDholePatilWardCaseByCreateDate get dhole patil ward cases by create date
func (dbRepo *DatabaseRepository) GetDholePatilWardCaseByCreateDate(createDate string, endDate string) ([]database.DholePatil, error) {
	var DBWardCase []database.DholePatil
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetHadapsarMundhwaWardCaseByCreateDate get hadapsar mundhwa ward cases by create date
func (dbRepo *DatabaseRepository) GetHadapsarMundhwaWardCaseByCreateDate(createDate, endDate string) ([]database.HadapsarMundhwa, error) {
	var DBWardCase []database.HadapsarMundhwa
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetKasbaWardCaseByCreateDate get Kasba Visharambagh wada ward cases by create date
func (dbRepo *DatabaseRepository) GetKasbaWardCaseByCreateDate(createDate string, endDate string) ([]database.KasbaVisharambaghwada, error) {
	var DBWardCase []database.KasbaVisharambaghwada
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetKondhwaWardCaseByCreateDate get Kondhwa Yewalewadi ward cases by create date
func (dbRepo *DatabaseRepository) GetKondhwaWardCaseByCreateDate(createDate string, endDate string) ([]database.KondhwaYewalewadi, error) {
	var DBWardCase []database.KondhwaYewalewadi
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetKothrudWardCaseByCreateDate get Kothrud Bavdhan ward cases by create date
func (dbRepo *DatabaseRepository) GetKothrudWardCaseByCreateDate(createDate string, endDate string) ([]database.KothrudBavdhan, error) {
	var DBWardCase []database.KothrudBavdhan
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetWadgoansheriWardCaseByCreateDate get Nagar Road Wadgoansheri ward cases by create date
func (dbRepo *DatabaseRepository) GetWadgoansheriWardCaseByCreateDate(createDate string, endDate string) ([]database.NagarRoadWadgoansheri, error) {
	var DBWardCase []database.NagarRoadWadgoansheri
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetShivajiNagarWardCaseByCreateDate get Shivaji Nagar Ghole Road ward cases by create date
func (dbRepo *DatabaseRepository) GetShivajiNagarWardCaseByCreateDate(createDate string, endDate string) ([]database.ShivajiNagarGholeRoad, error) {
	var DBWardCase []database.ShivajiNagarGholeRoad
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetSinhagadRoadWardCaseByCreateDate get Singhagad Road ward cases by create date
func (dbRepo *DatabaseRepository) GetSinhagadRoadWardCaseByCreateDate(createDate string, endDate string) ([]database.SinghagadRoad, error) {
	var DBWardCase []database.SinghagadRoad
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetWanawadiWardCaseByCreateDate get Wanawadi Ramtekdi ward cases by create date
func (dbRepo *DatabaseRepository) GetWanawadiWardCaseByCreateDate(createDate string, endDate string) ([]database.WanawadiRamtekdi, error) {
	var DBWardCase []database.WanawadiRamtekdi
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetWarjeWardCaseByCreateDate get Warje Karve Nagar ward cases by create date
func (dbRepo *DatabaseRepository) GetWarjeWardCaseByCreateDate(createDate string, endDate string) ([]database.WarjeKarveNagar, error) {
	var DBWardCase []database.WarjeKarveNagar
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetYerwadaWardCaseByCreateDate get Yerwada Kalas Dhanori ward cases by create date
func (dbRepo *DatabaseRepository) GetYerwadaWardCaseByCreateDate(createDate string, endDate string) ([]database.YerwadaKalasDhanori, error) {
	var DBWardCase []database.YerwadaKalasDhanori
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}

//GetPuneRuralWardCaseByCreateDate get Pune Rural ward cases by create date
func (dbRepo *DatabaseRepository) GetPuneRuralWardCaseByCreateDate(createDate string, endDate string) ([]database.PuneRural, error) {
	var DBWardCase []database.PuneRural
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Find(&DBWardCase).Error
	return DBWardCase, err
}
