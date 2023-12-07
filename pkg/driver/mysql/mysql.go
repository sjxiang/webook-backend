package mysql

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/sjxiang/webook-backend/internal/conf"
	"github.com/sjxiang/webook-backend/internal/data"
)

const RETRY_TIMES = 6

type MySQLConfig struct {
	Addr       string
	Port       string
	User       string
	Password   string
	Database   string
	Parameters string
}

func (m *MySQLConfig) Datasource() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		m.User, m.Password, m.Addr, m.Port, m.Database, m.Parameters)
}


func NewMySQLConnectionByGlobalConfig(config *conf.Config, logger *zap.SugaredLogger) (*gorm.DB, error) {
	mysqlConfig := &MySQLConfig{
		Addr:       config.GetMySQLAddr(),
		Port:       config.GetMySQLPort(),
		User:       config.GetMySQLUser(),
		Password:   config.GetMySQLPassword(),
		Database:   config.GetMySQLDatabase(),
		Parameters: config.GetMySQLParameters(),
	}
	return NewMySQLConnection(mysqlConfig, logger)
}

func NewMySQLConnection(config *MySQLConfig, logger *zap.SugaredLogger) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	retries := RETRY_TIMES

	db, err = gorm.Open(mysql.Open(config.Datasource()), &gorm.Config{
		// 不允许外键
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}

	for err != nil {
		if logger != nil {
			logger.Errorw("Failed to connect to database, %d", retries)
		}
		if retries > 1 {
			retries--
			time.Sleep(10 * time.Second)
			
			db, err = gorm.Open(mysql.Open(config.Datasource()), &gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
			})
			continue
		}
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Errorw("error in connecting db ", "db", config, "err", err)
		return nil, err
	}

	// check db connection
	err = sqlDB.Ping()
	if err != nil {
		logger.Errorw("error in connecting db ", "db", config, "err", err)
		return nil, err
	}

	db.AutoMigrate(&data.UserM{})

	logger.Infow("connected with db", "db", config)

	return db, err
}
