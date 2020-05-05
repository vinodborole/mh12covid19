package controllers

import (
	"covid19/config/domain"
	"covid19/infra"
	u "covid19/utils"
	"encoding/json"
	"net/http"
	"time"
)

//AddPatientSummary add patient summary
var AddPatientSummary = func(w http.ResponseWriter, r *http.Request) {
	patientSummary := &domain.PatientSummary{}
	err := json.NewDecoder(r.Body).Decode(patientSummary) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid add patient summary request: "+err.Error()))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	u.Infoln("Add patient summary : ", patientSummary)
	patientSummary, err = infra.GetUseCaseInteractor().AddPatientSummary(patientSummary)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := u.Message(true, "Added patient summary successfully")
	resp["patientSummary"] = patientSummary
	w.WriteHeader(http.StatusCreated)
	u.Respond(w, resp)
}

//GetPatientSummary get patient summary
var GetPatientSummary = func(w http.ResponseWriter, r *http.Request) {
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
	districtSummary, err := infra.GetUseCaseInteractor().GetPatientSummaryByCreateDate(sdate, edate)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := u.Message(true, "Patient summary fetch successfully")
	resp["patientSummary"] = districtSummary
	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}

//GetPatientSummaryDelta get delta patient summary
var GetPatientSummaryDelta = func(w http.ResponseWriter, r *http.Request) {
	patientDeltaSummary, err := infra.GetUseCaseInteractor().GetPatientDeltaSummary()
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := u.Message(true, "Patient summary latest fetch successfully")
	resp["patientDeltaSummary"] = patientDeltaSummary
	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}
