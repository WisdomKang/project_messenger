package database

import (
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBManager struct {
	DB gorm.DB
}

var lock = &sync.Mutex{}
var dbInstance *DBManager

func GetInstance() *DBManager {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if dbInstance == nil {
			log.Println("Create DBManager instance now.")
			dbInstance = &DBManager{
				DB: *dbConnect(),
			}
		} else {
			log.Println("DBManager intstance is already created.")
		}
	} else {
		log.Println("DBManager intstance is already created.")
	}

	return dbInstance
}

func dbConnect() *gorm.DB {
	dsn := "host=localhost user=msgclient password=test1234 dbname=my_messenger port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{})

	if err != nil {
		log.Panicf("database init error : %s\n", err)
	}

	log.Printf("database connected : %s\n", db.Migrator().CurrentDatabase())

	sqlDB, err := db.DB()

	if err != nil {
		log.Panicf("database init error : %s\n", err)
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	return db
}
