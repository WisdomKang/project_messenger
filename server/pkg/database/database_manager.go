package database

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
}
type DatabaseConfig struct {
	DNS          string `mapstructure:"DNS"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

var db *gorm.DB
var once sync.Once

func init() {
	exePath, err := os.Getwd()
	log.Println(exePath)
	if err != nil {
		log.Fatalf("Error finding executable path : %v", err)
		return
	}

	// 실행 파일이 있는 디렉토리 경로를 찾음
	exeDir := filepath.Dir(exePath)

	// 설정 파일의 경로를 조합
	configPath := filepath.Join(exeDir, "config.yaml")

	// read config.yaml
	viper.SetConfigFile(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file : %v", err)
	}
}

func InitDB() {
	// Unmarshal config file
	var databaseConfig DatabaseConfig
	err := viper.Unmarshal(&databaseConfig)

	if err != nil {
		log.Fatalf("Error inmarshal databaseConfig:%v", err)
	}

	// Set Database connection
	once.Do(func() {
		dsn := databaseConfig.DNS
		db, err := gorm.Open(
			postgres.Open(dsn),
			&gorm.Config{})

		if err != nil {
			log.Panicf("database init error : %s\n", err)
		}

		log.Printf("database connected : %s\n", db.Migrator().CurrentDatabase())

		// Set database schema
		err = db.Exec("SET search_path TO post_chest").Error

		if err != nil {
			log.Panicf("database schema set error : %s\n", err)
		}

		sqlDB, err := db.DB()

		if err != nil {
			log.Panicf("database init error : %s\n", err)
		}

		// Set connection option
		sqlDB.SetMaxIdleConns(databaseConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(databaseConfig.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Minute * 30)
	})
}

func GetDB() *gorm.DB { return db }
