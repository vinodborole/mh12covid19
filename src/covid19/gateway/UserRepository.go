package gateway

import (
	"covid19/config"
	"covid19/config/domain"
	"covid19/infra/database"
	u "covid19/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//CreateUser create account
func (dbRepo *DatabaseRepository) CreateUser(account *domain.User) (*database.User, error) {

	var DBAccount database.User
	u.Copy(&DBAccount, account)

	return dbRepo.CreateUserDB(&DBAccount)
}

//CreateUserDB create account db
func (dbRepo *DatabaseRepository) CreateUserDB(DBUser *database.User) (*database.User, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(DBUser.Password), bcrypt.DefaultCost)
	DBUser.Password = string(hashedPassword)

	if DBUser.ID != 0 {
		u.Infoln("account already exist, save the info")
		return DBUser, dbRepo.GetDBHandle().Save(&DBUser).Error
	}

	err := dbRepo.GetDBHandle().Create(&DBUser).Error

	if err != nil {
		u.Errorln(err.Error())
		return nil, errors.New("failed to create account,connection error: " + err.Error())
	}

	expireToken := time.Now().Add(time.Hour * 72).Unix()

	//Create new JWT token for the newly registered account
	tk := &domain.Token{UserID: DBUser.ID, StandardClaims: jwt.StandardClaims{
		ExpiresAt: expireToken,
		Issuer:    "bredec-auth",
	},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(config.GetConfig().MH12Config.Application.Token))
	DBUser.Token = tokenString

	DBUser.Password = "" //delete password

	return DBUser, nil
}

//GetUserByEmail get account by email
func (dbRepo *DatabaseRepository) GetUserByEmail(email string) (*database.User, error) {
	var DBUser database.User
	err := dbRepo.GetDBHandle().Model(domain.User{}).Where("email = ?", email).Find(&DBUser).Error
	DBUser.Password = ""
	return &DBUser, err
}

//Authenticate authenticate account
func (dbRepo *DatabaseRepository) Authenticate(email string, password string) (*database.User, error) {
	var DBUser database.User
	err := dbRepo.GetDBHandle().Model(domain.User{}).Where("email = ?", email).Find(&DBUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("email address not found")
		}
		return nil, errors.New("connection error. please retry")
	}
	err = bcrypt.CompareHashAndPassword([]byte(DBUser.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return nil, errors.New("invalid login credentials. Please try again")
	}
	DBUser.Password = ""
	return &DBUser, err
}
