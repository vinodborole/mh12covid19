package gateway

import (
	"covid19/config/domain"
	"covid19/infra/database"
	u "covid19/utils"
	"errors"
	"time"
)

//AddNews add news
func (dbRepo *DatabaseRepository) AddNews(news *domain.News) (*database.News, error) {
	var DBNews database.News
	u.Copy(&DBNews, news)
	if len(news.Date) > 0 {
		DBNews.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", news.Date)
	}
	return dbRepo.AddNewsDB(&DBNews)
}

//AddNewsDB add news db
func (dbRepo *DatabaseRepository) AddNewsDB(news *database.News) (*database.News, error) {
	err := dbRepo.GetDBHandle().Create(&news).Error
	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to add news,connection error: " + err.Error())
	}
	return news, nil
}

//GetNews get news by create date
func (dbRepo *DatabaseRepository) GetNews(createDate string, endDate string) ([]database.News, error) {
	var DBNews []database.News
	err := dbRepo.GetDBHandle().Where("date(created_at) >= ? and date(created_at) <= ?", createDate, endDate).Order("created_at desc").Limit(10).Find(&DBNews).Error
	return DBNews, err
}
