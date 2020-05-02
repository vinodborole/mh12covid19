package controllers

import (
	"covid19/config/domain"
	"covid19/infra"
	u "covid19/utils"
	"encoding/json"
	"net/http"
	"time"
)

//GetNews get news details
var GetNews = func(w http.ResponseWriter, r *http.Request) {
	sdate := r.URL.Query().Get("sdate")
	if sdate == "" {
		u.Respond(w, u.Message(false, "Missing Date with format YYYY-MM-DD"))
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	_, err := time.Parse("2006-01-02", sdate)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid Date format, correct format is YYYY-MM-DD"))
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	edate := r.URL.Query().Get("edate")
	if len(edate) > 0 {
		_, err = time.Parse("2006-01-02", edate)
		if err != nil {
			u.Respond(w, u.Message(false, "Invalid Date format, correct format is YYYY-MM-DD"))
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
	} else {
		edate = "2040-01-30"
	}
	news, err := infra.GetUseCaseInteractor().GetNews(sdate, edate)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := u.Message(true, "news fetch successfully")
	resp["news"] = news
	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}

//AddNews add news
var AddNews = func(w http.ResponseWriter, r *http.Request) {
	news := &domain.News{}
	err := json.NewDecoder(r.Body).Decode(news) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid add news request: "+err.Error()))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	u.Infoln("Add news : ", news)
	dbNews, err := infra.GetUseCaseInteractor().AddNews(news)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := u.Message(true, "Added news successfully")
	resp["news"] = dbNews
	w.WriteHeader(http.StatusCreated)
	u.Respond(w, resp)
}
