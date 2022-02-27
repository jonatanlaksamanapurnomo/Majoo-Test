package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var (
	host     = GetValueEnv("PG_HOST")
	port     = 2345
	user     = GetValueEnv("PG_USER")
	password = GetValueEnv("PG_PASSWORD")
	dbname   = GetValueEnv("PG_DB_NAME")
)

func GetConnection() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("DB Connection established...")
	return db
}
