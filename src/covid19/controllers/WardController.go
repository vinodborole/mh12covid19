package controllers

import (
	"covid19/infra"
	u "covid19/utils"
	"net/http"
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
