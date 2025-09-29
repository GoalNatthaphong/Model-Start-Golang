package db

import (
	"Goal/configs/logs"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func InitDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", host, user, password, dbname, port)
	if host == "" {
		dsn = os.Getenv("DATABASE_URL")
	}
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// AutoMigrate struct เป็น table
	// err = db.AutoMigrate(&models{})
	// if err != nil {
	// 	return nil, err
	// }
	logs.Info("Database migrated successfully!")
	return db, nil
}

func InitReadDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST_Read")
	user := os.Getenv("DB_USER_Read")
	password := os.Getenv("DB_PASSWORD_Read")
	dbname := os.Getenv("DB_NAME_Read")
	port := os.Getenv("DB_PORT_Read")
	if port == "" {
		port = "5432"
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	logs.Info("Database Read Connected successfully!")
	return db, nil
}