package database

import (
	"database/sql"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	databaseConfig "github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/response"
	"github.com/ZupIT/horusec-devkit/pkg/utils/logger"
)

type Database struct {
	connectionWrite *gorm.DB
	connectionRead  *gorm.DB
	config          databaseConfig.IConfig
}

func NewDatabaseReadAndWrite(config databaseConfig.IConfig) (IDatabaseRead, IDatabaseWrite, error) {
	if err := config.Validate(); err != nil {
		return nil, nil, err
	}

	database := &Database{config: config}
	database.makeConnection()
	database.setLogMode()
	return database, database, nil
}

func (d *Database) makeConnection() {
	d.makeConnectionWrite()
	d.makeConnectionRead()
}

func (d *Database) makeConnectionWrite() {
	connectionWrite, err := gorm.Open(postgres.Open(d.config.GetURI()), &gorm.Config{})
	if err != nil {
		logger.LogPanic(enums.FailedToConnectToDatabase, err)
	}

	d.connectionWrite = connectionWrite
}

func (d *Database) makeConnectionRead() {
	connectionRead, err := gorm.Open(postgres.Open(d.config.GetURI()), &gorm.Config{})
	if err != nil {
		logger.LogPanic(enums.FailedToConnectToDatabase, err)
	}

	d.connectionRead = connectionRead
}

func (d *Database) setLogMode() {
	if d.config.GetLogMode() {
		d.connectionWrite.Logger = d.connectionWrite.Logger.LogMode(gormLogger.Info)
		d.connectionRead.Logger = d.connectionRead.Logger.LogMode(gormLogger.Info)
		return
	}

	d.connectionWrite.Logger = d.connectionWrite.Logger.LogMode(gormLogger.Error)
	d.connectionRead.Logger = d.connectionRead.Logger.LogMode(gormLogger.Error)
}

func (d *Database) StartTransaction() IDatabaseWrite {
	return &Database{
		connectionWrite: d.connectionWrite.Begin(),
	}
}

func (d *Database) RollbackTransaction() response.IResponse {
	result := d.connectionWrite.Rollback()
	return response.NewResponse(result.RowsAffected, result.Error, nil)
}

func (d *Database) CommitTransaction() response.IResponse {
	result := d.connectionWrite.Commit()
	return response.NewResponse(result.RowsAffected, result.Error, nil)
}

func (d *Database) IsAvailable() bool {
	if d.connectionWrite == nil || d.connectionRead == nil {
		return false
	}

	if !d.pingDatabase(d.connectionWrite.DB()) {
		return false
	}

	return d.pingDatabase(d.connectionRead.DB())
}

func (d *Database) pingDatabase(db *sql.DB, err error) bool {
	if err != nil {
		logger.LogError(enums.FailedToVerifyIsAvailable, err)
		return false
	}

	return db.Ping() == nil
}

func (d *Database) Create(data interface{}, table string) response.IResponse {
	result := d.connectionWrite.Table(table).Create(data)

	return response.NewResponse(result.RowsAffected, result.Error, data)
}

func (d *Database) CreateOrUpdate(entity interface{}, where map[string]interface{}, table string) response.IResponse {
	result := d.connectionWrite.Table(table).Where(where).Save(entity)

	return response.NewResponse(result.RowsAffected, result.Error, entity)
}

func (d *Database) Find(entity interface{}, where map[string]interface{}, table string) response.IResponse {
	result := d.connectionRead.Table(table).Where(where).Find(entity)
	if err := d.verifyNotFoundError(result); err != nil {
		return response.NewResponse(result.RowsAffected, err, nil)
	}

	return response.NewResponse(result.RowsAffected, result.Error, entity)
}

func (d *Database) Update(entity interface{}, where map[string]interface{}, table string) response.IResponse {
	result := d.connectionWrite.Table(table).Where(where).Updates(entity)

	return response.NewResponse(result.RowsAffected, result.Error, entity)
}

func (d *Database) Delete(where map[string]interface{}, table string) response.IResponse {
	result := d.connectionWrite.Table(table).Where(where).Delete(nil)

	return response.NewResponse(result.RowsAffected, result.Error, nil)
}

func (d *Database) First(entity interface{}, where map[string]interface{}, table string) response.IResponse {
	result := d.connectionRead.Table(table).Where(where).First(entity)
	if err := d.verifyNotFoundError(result); err != nil {
		return response.NewResponse(result.RowsAffected, err, nil)
	}

	return response.NewResponse(result.RowsAffected, result.Error, entity)
}

func (d *Database) Raw(rawSQL string, entity interface{}) response.IResponse {
	result := d.connectionRead.Raw(rawSQL).Find(entity)
	if err := d.verifyNotFoundError(result); err != nil {
		return response.NewResponse(result.RowsAffected, err, nil)
	}

	return response.NewResponse(result.RowsAffected, result.Error, entity)
}

func (d *Database) verifyNotFoundError(result *gorm.DB) error {
	if result.Error != nil {
		if strings.EqualFold(result.Error.Error(), "record not found") {
			return enums.ErrNotFoundRecords
		}

		return result.Error
	}

	if result.RowsAffected == 0 {
		return enums.ErrNotFoundRecords
	}

	return nil
}
