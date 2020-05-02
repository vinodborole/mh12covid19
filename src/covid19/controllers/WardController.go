package controllers

import (
	"covid19/config/domain"
	"covid19/infra"
	u "covid19/utils"
	"encoding/json"
	"net/http"
	"time"
)

//GetWards retrieve customer lists rest api handler
var GetWards = func(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	wards := infra.GetUseCaseInteractor().GetWards(page)
	if wards == nil {
		u.Respond(w, u.Message(false, "No Wards list present, please add wards"))
		w.WriteHeader(http.StatusOK)
		return
	}
	u.Infoln("get wards : ", wards)
	resp := map[string]interface{}{"status": true, "message": "Retrieved wards list", "wards": wards}
	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}

//AddWardCase add ward case
var AddWardCase = func(w http.ResponseWriter, r *http.Request) {
	wardCase := &domain.WardCase{}
	err := json.NewDecoder(r.Body).Decode(wardCase) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid add ward case request: "+err.Error()))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	u.Infoln("Add ward case : ", wardCase)
	dbWardCase, err := infra.GetUseCaseInteractor().AddWardCase(wardCase)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := u.Message(true, "Added ward cases successfully")
	resp["wardCase"] = dbWardCase
	w.WriteHeader(http.StatusCreated)
	u.Respond(w, resp)
}

//GetWardDetails get ward details
var GetWardDetails = func(w http.ResponseWriter, r *http.Request) {
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
	code := r.URL.Query().Get("code")
	wardDetails := make([]*domain.WardCase, 0)
	if len(code) > 0 {
		wardDetails, err = infra.GetUseCaseInteractor().GetWardDetailsByCreateDateAndCode(sdate, edate, code)
		if err != nil {
			u.Respond(w, u.Message(false, err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		wardDetails, err = infra.GetUseCaseInteractor().GetAllWardDetailsByCreateDate(sdate, edate)
		if err != nil {
			u.Respond(w, u.Message(false, err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	resp := u.Message(true, "Ward details fetch successfully")
	resp["wardDetails"] = wardDetails
	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}
