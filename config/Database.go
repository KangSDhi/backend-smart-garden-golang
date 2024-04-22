package config

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
	"time"
)

var DB *sql.DB

func InitDB() *sql.DB {
	DB = connectDB()
	return DB
}

func connectDB() *sql.DB {
	var err error

	cfg := mysql.Config{
		User:                 os.Getenv("DB_USERNAME"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		DBName:               os.Getenv("DB_DATABASE"),
		AllowNativePasswords: true,
	}

	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Error connecting database :", err)
		return nil
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal("Error Ping Database :", pingErr)
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(5 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)

	log.Println("Successfully connected to the database")

	return DB
}
