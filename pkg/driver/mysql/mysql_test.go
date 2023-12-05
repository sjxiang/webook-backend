package mysql

import (
	"testing"

	"go.uber.org/zap"
	"github.com/stretchr/testify/require"
)

func TestNewMySQLConnection(t *testing.T) {
	mysqlConfig := &MySQLConfig{
		Addr:       "localhost",
		Port:       "3306",
		User:       "root",
		Password:   "123456",
		Database:   "webook_backend",
		Parameters: "charset=utf8mb4&parseTime=true&loc=Local",
	}

	db, err := NewMySQLConnection(mysqlConfig, zap.S())
	require.NoError(t, err)
	require.NotEmpty(t, db)
}