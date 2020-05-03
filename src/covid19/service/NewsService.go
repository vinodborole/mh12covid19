package service

import (
	"covid19/config"
	"covid19/config/domain"
	"covid19/infra/database"
	"encoding/json"
	"io/ioutil"
	"time"
)

//AddNews add news
func (sh *Interactor) AddNews(news *domain.News) (*database.News, error) {
	dbNews, err := sh.Db.AddNews(news)
	if err != nil {
		return nil, err
	}
	defer sh.GenerateNewsJSON()
	defer sh.GenerateMetaJSON()
	return dbNews, nil
}

//GetNews get news
func (sh *Interactor) GetNews(sdate string, edate string) ([]domain.News, error) {
	newsList := make([]domain.News, 0)
	dbNews, err := sh.Db.GetNews(sdate, edate)
	if err != nil {
		return nil, err
	}
	for _, news := range dbNews {
		var domainNews domain.News
		domainNews.Source = news.Source
		domainNews.Summary = news.Summary
		domainNews.Date = news.CreatedAt.Format("2006-01-02 15:04:05")
		newsList = append(newsList, domainNews)
	}
	return newsList, nil
}

//GenerateNewsJSON generate news JSON
func (sh *Interactor) GenerateNewsJSON() error {
	news, err := sh.GetNews("2019-12-01", "2040-12-31")
	if err != nil {
		return err
	}
	file, _ := json.MarshalIndent(news, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"punenews.json", file, 0644)
	return nil
}

//GenerateMetaJSON generate meta JSON
func (sh *Interactor) GenerateMetaJSON() error {
	metaDetails := domain.MetaDetails{LastUpdateDateTime: time.Now().Format("2006-01-02 15:04:05")}
	file, _ := json.MarshalIndent(metaDetails, "", " ")
	_ = ioutil.WriteFile(config.GetConfig().MH12Config.Application.ExportJSONPath+"meta-details.json", file, 0644)
	return nil
}
