package controllers

import (
	"covid19/config/domain"
	"covid19/infra"
	u "covid19/utils"
	"encoding/json"
	"net/http"
	"time"
)

//AddDistrictSummary add district summary
var AddDistrictSummary = func(w http.ResponseWriter, r *http.Request) {
	districtCase := &domain.DistrictCase{}
	err := json.NewDecoder(r.Body).Decode(districtCase) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid add district case request: "+err.Error()))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	u.Infoln("Add district case summary: ", districtCase)
	districtCase, err = infra.GetUseCaseInteractor().AddDistrictCaseSummary(districtCase)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := u.Message(true, "Added district case summary successfully")
	resp["districtCase"] = districtCase
	w.WriteHeader(http.StatusCreated)
	u.Respond(w, resp)
}

//GetDistrictSummary get district summary
var GetDistrictSummary = func(w http.ResponseWriter, r *http.Request) {
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
		_, err := time.Parse("2006-01-02", edate)
		if err != nil {
			u.Respond(w, u.Message(false, "Invalid Date format, correct format is YYYY-MM-DD"))
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}
	} else {
		edate = "2040-01-30"
	}
	districtSummary, err := infra.GetUseCaseInteractor().GetDistrictSummaryByCreateDate(sdate, edate)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := u.Message(true, "District case summary fetch successfully")
	resp["districtSummary"] = districtSummary
	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}

//GetDistrictSummaryLatest get latest district summary
var GetDistrictSummaryLatest = func(w http.ResponseWriter, r *http.Request) {
	districtSummaryLatest, err := infra.GetUseCaseInteractor().GetLatestDistrictSummary()
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := u.Message(true, "District case summary latest fetch successfully")
	resp["districtSummaryLatest"] = districtSummaryLatest
	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}
