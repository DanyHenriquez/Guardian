package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

func Connect() *gorm.DB {
	var conn string

	switch os.Getenv("DB_TYPE") {
	case "mysql":
		conn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)

	case "postgres":
		conn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_DATABASE"),
			os.Getenv("DB_PASSWORD"),
		)

	case "sqlite":
		conn = fmt.Sprintf("%s",
			os.Getenv("DB_DATABASE"),
		)

	case "mssql":
		conn = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DATABASE"),
		)
	}

	db, err := gorm.Open("mysql", conn)

	if err != nil {
		fmt.Println("Cannot connect to database.")
		os.Exit(2)
	}

	return db
}
