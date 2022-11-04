package config

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"graphql_go/src/shared"
	"graphql_go/src/shared/log"
)

type Config struct {
	DBConn *gorm.DB
}

type Env struct {
	DBConn     *gorm.DB
	HTTPPort   string `env:"HTTP_PORT;required"`
	DBHost     string `env:"DB_HOST;required"`
	DBPort     string `env:"DB_PORT;required"`
	DBUser     string `env:"DB_USER;required"`
	DBPassword string `env:"DB_PASSWORD;required"`
	DBName     string `env:"DB_NAME;required"`
}

var GlobalEnv Env

func Init() *Config {
	var err error
	cfg := Config{}

	err = shared.ParseEnv(&GlobalEnv)
	if err != nil {
		panic(err)
	}

	if cfg.DBConn == nil {
		connString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
		dsn := fmt.Sprintf(connString, GlobalEnv.DBHost, GlobalEnv.DBUser, GlobalEnv.DBPassword, GlobalEnv.DBName, GlobalEnv.DBPort)
		cfg.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Logger().Error(err)
			panic(err)
		}
		sqlDB, _ := cfg.DBConn.DB()
		sqlDB.SetConnMaxIdleTime(10)
		sqlDB.SetConnMaxLifetime(5 * time.Minute)
		sqlDB.SetMaxOpenConns(512)

		err = sqlDB.Ping()
		if err != nil {
			log.Logger().Error(err)
			panic(err)
		}

		log.Logger().Info(fmt.Sprintf(connString, GlobalEnv.DBHost, GlobalEnv.DBUser, "********", GlobalEnv.DBName, GlobalEnv.DBPort))

		GlobalEnv.DBConn = cfg.DBConn
	}
	return &cfg
}
