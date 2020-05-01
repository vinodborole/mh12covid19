package controllers

import (
	"covid19/config"
	"covid19/config/domain"
	"covid19/infra"
	u "covid19/utils"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

//CreateUser create user rest api handler
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	User := &domain.User{}
	err := json.NewDecoder(r.Body).Decode(User) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid create User request: "+err.Error()))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	u.Infoln("Create User : ", User)
	dbUser, err := infra.GetUseCaseInteractor().CreateUser(User)
	if err != nil {
		u.Respond(w, u.Message(false, err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := u.Message(true, "User created successfully")
	resp["User"] = dbUser
	w.WriteHeader(http.StatusCreated)
	u.Respond(w, resp)
}

//Authenticate authentic rest api handler
var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	User := &domain.User{}
	err := json.NewDecoder(r.Body).Decode(User) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid authenticate request: "+err.Error()))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	u.Infoln("Authenticate User : ", User)
	dbUser, err := infra.GetUseCaseInteractor().Login(User.Email, User.Password)
	if err != nil {
		u.Errorln("Incorrect login details - ", err.Error())
		u.Respond(w, u.Message(false, "Incorrect login details: "+err.Error()))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	u.Infoln("login successful")
	resp := u.Message(true, "Login successfully")
	resp["User"] = dbUser

	i := config.GetConfig().MH12Config.Application.TokenExpiryHrs
	expirationTime := time.Now().Add(time.Duration(i) * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   dbUser.Token,
		Expires: expirationTime,
	})
	u.Infoln("token : ", dbUser.Token)
	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}

//RefreshToken refresh token rest api handler
var RefreshToken = func(w http.ResponseWriter, r *http.Request) {
	var token string
	User := &domain.User{}
	err := json.NewDecoder(r.Body).Decode(User) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid User details for refresh token: "+err.Error()))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	token = User.Token
	if len(token) == 0 {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token = c.Value
	}
	tk := &domain.Token{}

	refreshToken, err := jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().MH12Config.Application.Token), nil
	})

	if err != nil { //Malformed token, returns with http code 403 as usual
		u.Respond(w, u.Message(false, "Malformed authentication token: "+err.Error()))
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if !refreshToken.Valid { //Token is invalid, maybe not signed on this server
		u.Respond(w, u.Message(false, "Token is not valid."))
		w.WriteHeader(http.StatusForbidden)
		return
	}

	dbUser, err := infra.GetUseCaseInteractor().RefreshToken(User.Email)
	if err != nil {
		u.Respond(w, u.Message(false, "Incorrect login details: "+err.Error()))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	resp := u.Message(true, "Refresh token received successfully")
	resp["User"] = dbUser

	i := config.GetConfig().MH12Config.Application.TokenExpiryHrs

	expirationTime := time.Now().Add(time.Duration(i) * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   dbUser.Token,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
	u.Respond(w, resp)
}
