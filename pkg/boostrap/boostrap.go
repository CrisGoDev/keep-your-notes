package boostrap

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/CrisGoDev/keep-your-notes/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitLogger() *log.Logger {
	l := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	return l
}

func GetPort() (int, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return 0, err
	}

	return port, nil
}

func DBConnection() (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if os.Getenv("DATABASE_DEBUG") == "true" {
		db = db.Debug()
	}

	if os.Getenv("DATABASE_MIGRATE") == "true" {
		if err := db.AutoMigrate(&domain.Note{}); err != nil {
			return nil, err
		}
	}

	return db, err

}
