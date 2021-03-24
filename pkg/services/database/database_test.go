package database

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ZupIT/horusec-devkit/pkg/services/database/config"
	"github.com/ZupIT/horusec-devkit/pkg/services/database/enums"
)

type testEntity struct {
	text string
}

func newTestEntity() *testEntity {
	return &testEntity{text: "test"}
}

func getMockedConnection(db *sql.DB) *gorm.DB {
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	connection, _ := gorm.Open(dialector, &gorm.Config{})
	return connection
}

func TestNewDatabaseReadAndWrite(t *testing.T) {
	t.Run("should panic when failed to connect to database", func(t *testing.T) {
		assert.Panics(t, func() {
			databaseConfig := &config.Config{}
			databaseConfig.SetURI("test")

			_, _ = NewDatabaseReadAndWrite(databaseConfig)
		})
	})

	t.Run("should return error when invalid config", func(t *testing.T) {
		_, err := NewDatabaseReadAndWrite(&config.Config{})

		assert.Error(t, err)
	})
}

func TestSetConnections(t *testing.T) {
	t.Run("should success set connections", func(t *testing.T) {
		database := &database{}

		assert.NotNil(t, database.setConnections())
	})
}

func TestMakeConnectionWrite(t *testing.T) {
	t.Run("should panic when failed to connect to database", func(t *testing.T) {
		database := &database{config: &config.Config{}}

		assert.Panics(t, func() {
			database.makeConnectionWrite()
		})
	})
}

func TestMakeConnectionRead(t *testing.T) {
	t.Run("should panic when failed to connect to database", func(t *testing.T) {
		database := &database{config: &config.Config{}}

		assert.Panics(t, func() {
			database.makeConnectionRead()
		})
	})
}

func TestSetLogMode(t *testing.T) {
	t.Run("should not panic when setting log mode false", func(t *testing.T) {
		db, _, err := sqlmock.New()
		assert.NoError(t, err)

		databaseConfig := config.NewDatabaseConfig()
		databaseConfig.SetLogMode(false)

		database := &database{
			config:          databaseConfig,
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		assert.NotPanics(t, func() {
			database.setLogMode()
		})
	})

	t.Run("should not panic when setting log mode true", func(t *testing.T) {
		db, _, err := sqlmock.New()
		assert.NoError(t, err)

		databaseConfig := config.NewDatabaseConfig()
		databaseConfig.SetLogMode(true)

		database := &database{
			config:          databaseConfig,
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		assert.NotPanics(t, func() {
			database.setLogMode()
		})
	})
}

func TestStartTransaction(t *testing.T) {
	t.Run("should success start transaction and not panic", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectBegin()

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		assert.NotPanics(t, func() {
			assert.NotEmpty(t, database.StartTransaction())
		})
	})
}

func TestRollbackTransaction(t *testing.T) {
	t.Run("should success rollback transaction with no errors", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectBegin()
		mock.ExpectRollback()

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		assert.NotPanics(t, func() {
			transaction := database.StartTransaction()
			assert.NoError(t, transaction.RollbackTransaction().GetError())
		})
	})
}

func TestCommitTransaction(t *testing.T) {
	t.Run("should success commit transaction with no errors", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectBegin()
		mock.ExpectCommit()

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		assert.NotPanics(t, func() {
			transaction := database.StartTransaction()
			assert.NoError(t, transaction.CommitTransaction().GetError())
		})
	})
}

func TestIsAvailable(t *testing.T) {
	t.Run("should return true when database connections are ok", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectPing()

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		assert.True(t, database.IsAvailable())
	})

	t.Run("should return false when ping returns error", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
		assert.NoError(t, err)

		mock.ExpectPing().WillReturnError(errors.New("test"))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		assert.False(t, database.IsAvailable())
	})

	t.Run("should return false when any connection is nil", func(t *testing.T) {
		db, _, err := sqlmock.New()
		assert.NoError(t, err)

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: nil,
		}

		assert.False(t, database.IsAvailable())

		database.connectionWrite = getMockedConnection(db)
		database.connectionRead = nil

		assert.False(t, database.IsAvailable())
	})
}

func TestPingDatabase(t *testing.T) {
	t.Run("should return false when err is not nil", func(t *testing.T) {
		database := &database{}

		assert.False(t, database.pingDatabase(nil, errors.New("test")))
	})
}

func TestCreate(t *testing.T) {
	t.Run("should success create a new entity", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectExec("INSERT").
			WillReturnResult(sqlmock.NewResult(1, 1))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.Create(newTestEntity(), "test")

		assert.NoError(t, response.GetError())
		assert.Equal(t, 1, response.GetRowsAffected())
		assert.Equal(t, newTestEntity(), response.GetData())
	})
}

func TestCreateOrUpdate(t *testing.T) {
	t.Run("should success create or update a new entity", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"text"}).
				AddRow("test"))

		mock.ExpectExec("INSERT").
			WithArgs(sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.CreateOrUpdate(newTestEntity(), map[string]interface{}{"text": "test"}, "test")

		assert.NoError(t, response.GetError())
		assert.Equal(t, 0, response.GetRowsAffected())
		assert.Equal(t, newTestEntity(), response.GetData())
	})
}

func TestFind(t *testing.T) {
	t.Run("should success find a database record", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"text"}).
				AddRow("test"))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.Find(newTestEntity(), map[string]interface{}{"text": "test"}, "test")

		assert.NoError(t, response.GetError())
		assert.Equal(t, 1, response.GetRowsAffected())
		assert.Equal(t, newTestEntity(), response.GetData())
	})

	t.Run("should return error not found records when database return it", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WithArgs(sqlmock.AnyArg()).
			WillReturnError(errors.New("record not found"))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.Find(newTestEntity(), map[string]interface{}{"text": "test"}, "test")

		assert.Error(t, response.GetError())
		assert.Equal(t, enums.ErrorNotFoundRecords, response.GetError())
		assert.Equal(t, 0, response.GetRowsAffected())
		assert.Equal(t, nil, response.GetData())
	})

	t.Run("should return error not found records when no rows affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"text"}))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.Find(newTestEntity(), map[string]interface{}{"text": "test"}, "test")

		assert.Error(t, response.GetError())
		assert.Equal(t, enums.ErrorNotFoundRecords, response.GetError())
		assert.Equal(t, 0, response.GetRowsAffected())
		assert.Equal(t, nil, response.GetData())
	})

	t.Run("should return a error different than not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WillReturnError(errors.New("test"))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.Find(newTestEntity(), map[string]interface{}{"text": "test"}, "test")

		assert.Error(t, response.GetError())
		assert.NotEqual(t, enums.ErrorNotFoundRecords, response.GetError())
		assert.Equal(t, 0, response.GetRowsAffected())
		assert.Equal(t, nil, response.GetData())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should success update entity", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectExec("UPDATE").
			WithArgs(sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.Update(newTestEntity(), map[string]interface{}{"text": "test"}, "test")

		assert.NoError(t, response.GetError())
		assert.Equal(t, 0, response.GetRowsAffected())
		assert.Equal(t, newTestEntity(), response.GetData())
	})
}

func TestDelete(t *testing.T) {
	t.Run("should success delete entity", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectExec("DELETE").
			WithArgs(sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.Delete(map[string]interface{}{"text": "test"}, "test")

		assert.NoError(t, response.GetError())
		assert.Equal(t, 1, response.GetRowsAffected())
		assert.Nil(t, response.GetData())
	})
}

func TestFirst(t *testing.T) {
	t.Run("should success get first entity", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WithArgs(sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"text"}).
				AddRow("test"))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.First(newTestEntity(), map[string]interface{}{"text": "test"}, "test")

		assert.NoError(t, response.GetError())
		assert.Equal(t, 1, response.GetRowsAffected())
		assert.Equal(t, newTestEntity(), response.GetData())
	})

	t.Run("should return error not found records when database return it", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WithArgs(sqlmock.AnyArg()).
			WillReturnError(errors.New("record not found"))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.First(newTestEntity(), map[string]interface{}{"text": "test"}, "test")

		assert.Error(t, response.GetError())
		assert.Equal(t, enums.ErrorNotFoundRecords, response.GetError())
		assert.Equal(t, 0, response.GetRowsAffected())
		assert.Equal(t, nil, response.GetData())
	})

	t.Run("should return error not found records when no rows affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"text"}))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.First(newTestEntity(), map[string]interface{}{"text": "test"}, "test")

		assert.Error(t, response.GetError())
		assert.Equal(t, enums.ErrorNotFoundRecords, response.GetError())
		assert.Equal(t, 0, response.GetRowsAffected())
		assert.Equal(t, nil, response.GetData())
	})

	t.Run("should return a error different than not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WillReturnError(errors.New("test"))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.First(newTestEntity(), map[string]interface{}{"text": "test"}, "test")

		assert.Error(t, response.GetError())
		assert.NotEqual(t, enums.ErrorNotFoundRecords, response.GetError())
		assert.Equal(t, 0, response.GetRowsAffected())
		assert.Equal(t, nil, response.GetData())
	})
}

func TestRaw(t *testing.T) {
	t.Run("should success get by raw sql", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"text"}).
				AddRow("test"))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.Raw("SELECT * FROM \"test\"", newTestEntity())

		assert.NoError(t, response.GetError())
		assert.Equal(t, 1, response.GetRowsAffected())
		assert.Equal(t, newTestEntity(), response.GetData())
	})

	t.Run("should return error not found records when database return it", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WillReturnError(errors.New("record not found"))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.Raw("SELECT * FROM \"test\"", newTestEntity())

		assert.Error(t, response.GetError())
		assert.Equal(t, enums.ErrorNotFoundRecords, response.GetError())
		assert.Equal(t, 0, response.GetRowsAffected())
		assert.Equal(t, nil, response.GetData())
	})

	t.Run("should return error not found records when no rows affected", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WillReturnRows(sqlmock.NewRows([]string{"text"}))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.Raw("SELECT * FROM \"test\"", newTestEntity())

		assert.Error(t, response.GetError())
		assert.Equal(t, enums.ErrorNotFoundRecords, response.GetError())
		assert.Equal(t, 0, response.GetRowsAffected())
		assert.Equal(t, nil, response.GetData())
	})

	t.Run("should return a error different than not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery("SELECT").
			WillReturnError(errors.New("test"))

		database := &database{
			config:          config.NewDatabaseConfig(),
			connectionRead:  getMockedConnection(db),
			connectionWrite: getMockedConnection(db),
		}

		response := database.Raw("SELECT * FROM \"test\"", newTestEntity())

		assert.Error(t, response.GetError())
		assert.NotEqual(t, enums.ErrorNotFoundRecords, response.GetError())
		assert.Equal(t, 0, response.GetRowsAffected())
		assert.Equal(t, nil, response.GetData())
	})
}
