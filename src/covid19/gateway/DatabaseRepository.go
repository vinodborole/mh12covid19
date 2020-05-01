package gateway

import (
	"covid19/infra/database"
	"github.com/jinzhu/gorm"
)

//DatabaseRepository represents the Application Database Repository
type DatabaseRepository struct {
	Database    *database.Database
	Transaction *gorm.DB
}

//GetDBHandle returns a handle to DB, using which the database operations can be performed
func (dbRepo *DatabaseRepository) GetDBHandle() *gorm.DB {
	if dbRepo.Transaction != nil {
		return dbRepo.Transaction
	}
	return dbRepo.Database.Instance
}

//Ping underlying DB
func (dbRepo *DatabaseRepository) Ping() error {
	err := dbRepo.GetDBHandle().DB().Ping()
	return err
}

//OpenTransaction begins the database transaction
func (dbRepo *DatabaseRepository) OpenTransaction() error {
	dbRepo.Transaction = dbRepo.GetDBHandle().Begin()
	return dbRepo.Transaction.Error
}

//CommitTransaction commits the database transaction
func (dbRepo *DatabaseRepository) CommitTransaction() error {

	err := dbRepo.GetDBHandle().Commit().Error
	if err != nil {
		panic(err)
	}
	dbRepo.Transaction = nil
	return err
}

//RollBackTransaction rolls back the database transaction
func (dbRepo *DatabaseRepository) RollBackTransaction() error {

	err := dbRepo.GetDBHandle().Rollback().Error
	dbRepo.Transaction = nil
	return err
}
