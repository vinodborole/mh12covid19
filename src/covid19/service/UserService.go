package service

import (
	"covid19/config"
	"covid19/config/domain"
	"covid19/infra/database"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

//CreateUser create user account service
func (sh *Interactor) CreateUser(account *domain.User) (*database.User, error) {

	//validate
	if len(account.Name) < 2 {
		return nil, errors.New("name is required")
	}
	if !strings.Contains(account.Email, "@") {
		return nil, errors.New("email address is required")
	}

	if len(account.Password) < 2 {
		return nil, errors.New("password is required")
	}

	dbaccount, err := sh.Db.CreateUser(account)
	if err != nil {
		return nil, err
	}
	expireToken := time.Now().Add(time.Hour * 72).Unix()

	//Create new JWT token for the newly registered account
	tk := &domain.Token{UserID: dbaccount.ID, StandardClaims: jwt.StandardClaims{
		ExpiresAt: expireToken,
		Issuer:    "bredec-auth",
	},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(config.GetConfig().MH12Config.Application.Token))
	dbaccount.Token = tokenString

	dbaccount.Password = "" //delete password

	return dbaccount, nil
}

//Login login service
func (sh *Interactor) Login(email string, password string) (*database.User, error) {
	dbaccount, err := sh.Db.Authenticate(email, password)
	if err != nil {
		return dbaccount, err
	}
	i := config.GetConfig().MH12Config.Application.TokenExpiryHrs
	expirationTime := time.Now().Add(time.Duration(i) * time.Hour)
	//Create JWT token
	tk := &domain.Token{UserID: dbaccount.ID, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(config.GetConfig().MH12Config.Application.Token))
	if err != nil {
		return nil, err
	}
	dbaccount.Token = tokenString //Store the token in the response
	return dbaccount, nil
}

//RefreshToken refresh token service
func (sh *Interactor) RefreshToken(email string) (*database.User, error) {
	dbUser, err := sh.Db.GetUserByEmail(email)
	if err != nil {
		return dbUser, err
	}
	i := config.GetConfig().MH12Config.Application.TokenExpiryHrs
	expirationTime := time.Now().Add(time.Duration(i) * time.Hour)
	//Create JWT token
	tk := &domain.Token{UserID: dbUser.ID, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(config.GetConfig().MH12Config.Application.Token))
	if err != nil {
		return nil, err
	}
	dbUser.Token = tokenString //Store the token in the response

	return dbUser, nil
}
