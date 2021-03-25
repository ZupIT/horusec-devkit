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

type (
	database struct {
		connectionWrite *gorm.DB
		connectionRead  *gorm.DB
		config          databaseConfig.IConfig
	}

	Connection struct {
		Read  IDatabaseRead
		Write IDatabaseWrite
	}
)

func NewDatabaseReadAndWrite(config databaseConfig.IConfig) (*Connection, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	database := &database{config: config}
	database.makeConnection()
	database.setLogMode()
	return database.setConnections(), nil
}

func (d *database) setConnections() *Connection {
	return &Connection{
		Read:  d,
		Write: d,
	}
}

func (d *database) makeConnection() {
	d.makeConnectionWrite()
	d.makeConnectionRead()
}

func (d *database) makeConnectionWrite() {
	connectionWrite, err := gorm.Open(postgres.Open(d.config.GetURI()), &gorm.Config{})
	if err != nil {
		logger.LogPanic(enums.MessageFailedToConnectToDatabase, err)
	}

	d.connectionWrite = connectionWrite
}

func (d *database) makeConnectionRead() {
	connectionRead, err := gorm.Open(postgres.Open(d.config.GetURI()), &gorm.Config{})
	if err != nil {
		logger.LogPanic(enums.MessageFailedToConnectToDatabase, err)
	}

	d.connectionRead = connectionRead
}

func (d *database) setLogMode() {
	if d.config.GetLogMode() {
		d.connectionWrite.Logger = d.connectionWrite.Logger.LogMode(gormLogger.Info)
		d.connectionRead.Logger = d.connectionRead.Logger.LogMode(gormLogger.Info)
		return
	}

	d.connectionWrite.Logger = d.connectionWrite.Logger.LogMode(gormLogger.Error)
	d.connectionRead.Logger = d.connectionRead.Logger.LogMode(gormLogger.Error)
}

func (d *database) StartTransaction() IDatabaseWrite {
	return &database{
		connectionWrite: d.connectionWrite.Begin(),
	}
}

func (d *database) RollbackTransaction() response.IResponse {
	result := d.connectionWrite.Rollback()
	return response.NewResponse(result.RowsAffected, result.Error, nil)
}

func (d *database) CommitTransaction() response.IResponse {
	result := d.connectionWrite.Commit()
	return response.NewResponse(result.RowsAffected, result.Error, nil)
}

func (d *database) IsAvailable() bool {
	if d.connectionWrite == nil || d.connectionRead == nil {
		return false
	}

	if !d.pingDatabase(d.connectionWrite.DB()) {
		return false
	}

	return d.pingDatabase(d.connectionRead.DB())
}

func (d *database) pingDatabase(db *sql.DB, err error) bool {
	if err != nil {
		logger.LogError(enums.MessageFailedToVerifyIsAvailable, err)
		return false
	}

	return db.Ping() == nil
}

func (d *database) Create(entityPointer interface{}, table string) response.IResponse {
	result := d.connectionWrite.Table(table).Create(entityPointer)

	return response.NewResponse(result.RowsAffected, result.Error, entityPointer)
}

func (d *database) CreateOrUpdate(entityPointer interface{}, where map[string]interface{},
	table string) response.IResponse {
	result := d.connectionWrite.Table(table).Where(where).Save(entityPointer)

	return response.NewResponse(result.RowsAffected, result.Error, entityPointer)
}

func (d *database) Find(entityPointer interface{}, where map[string]interface{}, table string) response.IResponse {
	result := d.connectionRead.Table(table).Where(where).Find(entityPointer)
	if err := d.verifyNotFoundError(result); err != nil {
		return response.NewResponse(result.RowsAffected, err, nil)
	}

	return response.NewResponse(result.RowsAffected, result.Error, entityPointer)
}

func (d *database) Update(entityPointer interface{}, where map[string]interface{}, table string) response.IResponse {
	result := d.connectionWrite.Table(table).Where(where).Updates(entityPointer)

	return response.NewResponse(result.RowsAffected, result.Error, entityPointer)
}

func (d *database) Delete(where map[string]interface{}, table string) response.IResponse {
	result := d.connectionWrite.Table(table).Where(where).Delete(nil)

	return response.NewResponse(result.RowsAffected, result.Error, nil)
}

func (d *database) First(entityPointer interface{}, where map[string]interface{}, table string) response.IResponse {
	result := d.connectionRead.Table(table).Where(where).First(entityPointer)
	if err := d.verifyNotFoundError(result); err != nil {
		return response.NewResponse(result.RowsAffected, err, nil)
	}

	return response.NewResponse(result.RowsAffected, result.Error, entityPointer)
}

func (d *database) Raw(rawSQL string, entityPointer interface{}) response.IResponse {
	result := d.connectionRead.Raw(rawSQL).Find(entityPointer)
	if err := d.verifyNotFoundError(result); err != nil {
		return response.NewResponse(result.RowsAffected, err, nil)
	}

	return response.NewResponse(result.RowsAffected, result.Error, entityPointer)
}

func (d *database) verifyNotFoundError(result *gorm.DB) error {
	if result.Error != nil {
		if strings.EqualFold(result.Error.Error(), "record not found") {
			return enums.ErrorNotFoundRecords
		}

		return result.Error
	}

	if result.RowsAffected == 0 {
		return enums.ErrorNotFoundRecords
	}

	return nil
}
