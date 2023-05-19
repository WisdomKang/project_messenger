package database

import (
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBService struct {
	DB gorm.DB
}

// func init() {
// 	dsn := "host=localhost user=msgclient password=test1234 dbname=my_messenger port=5432 sslmode=disable TimeZone=Asia/Seoul"
// 	db, err := gorm.Open(
// 		postgres.Open(dsn),
// 		&gorm.Config{})

// 	if err != nil {
// 		log.Panicf("database init error : %s\n", err)
// 	}

// 	log.Printf("database connected : %s\n", db.Migrator().CurrentDatabase())

// 	sqlDB, err := db.DB()

// 	if err != nil {
// 		log.Panicf("database init error : %s\n", err)
// 	}

// 	sqlDB.SetMaxIdleConns(10)

// 	sqlDB.SetMaxOpenConns(100)

// 	sqlDB.SetConnMaxLifetime(time.Minute * 30)

// }

var lock = &sync.Mutex{}
var dbInstance *DBService

func GetInstance() *DBService {
	if dbInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if dbInstance == nil {
			log.Println("Create DBService instance now.")
			dbInstance = &DBService{
				DB: *dbConnect(),
			}
		} else {
			log.Println("DBService intstance is already created.")
		}
	} else {
		log.Println("DBService intstance is already created.")
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
