package dbpart

import (
	"ReadBook/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

/*
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
}*/

// ProvideDB provides PostgreSQL connection
func ProvideDB() *sqlx.DB {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Options().DBParam.Host, config.Options().DBParam.Port, config.Options().DBParam.Username,
		config.Options().DBParam.DBName, config.Options().DBParam.Password, config.Options().DBParam.SSLMode))
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping to PostgreSQL: %v", err)
	}
	return db
}

var Options = ProvideDB
