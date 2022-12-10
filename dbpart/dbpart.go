package dbpart

import (
	"ReadBook/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// DBHandler struct for HTTP requests
type DBHandler struct {
	logger *zap.SugaredLogger
	dbcon  *sqlx.DB
}

// New creates a DBHandler struct
func New(logger *zap.SugaredLogger) *DBHandler {
	dbcon, _ := ConnectDB()
	h := DBHandler{logger, dbcon}

	return &h
}

// ProvideDB provides PostgreSQL connection
func ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Options().DBParam.Host, config.Options().DBParam.Port, config.Options().DBParam.Username,
		config.Options().DBParam.DBName, config.Options().DBParam.Password, config.Options().DBParam.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
