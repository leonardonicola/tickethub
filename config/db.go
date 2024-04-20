package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var (
	db     *sql.DB
	logger *Logger
)

func InitDB() {
	logger = NewLogger()
	cstr, ok := os.LookupEnv("DB_URL")
	if !ok {
		logger.Fatal("couldn't find DB_URL variable")
	}
	var err error
	db, err = sql.Open("postgres", cstr)
	if err != nil {
		logger.Fatalf("couldn't open a connection to DB: %v", err)
	}
	err = db.Ping()
	if err != nil {
		logger.Errorf("couldn't ping to DB: %v", err)
	}
}

func GetDB() *sql.DB {
	return db
}
