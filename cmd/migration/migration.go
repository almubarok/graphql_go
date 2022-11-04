package main

import (
	"fmt"

	dotenv "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"graphql_go/config"
	"graphql_go/src/shared/log"

	userModel "graphql_go/src/modules/user/model"
)

func getMigrateTables() []interface{} {
	return []interface{}{
		userModel.User{},
	}
}

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		msg := ".env is not loaded properly"
		log.Logger().Error(fmt.Errorf(msg))
		panic(msg)
	}
	cfg := config.Init()
	conn, err := cfg.DBConn.DB()
	if err != nil {
		msg := "[DB] unable to connect " + err.Error()
		log.Logger().Error(fmt.Errorf(msg))
		panic(msg)
	}

	gormMigrate, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}))
	if err != nil {
		msg := "[DB] unable to connect " + err.Error()
		log.Logger().Error(fmt.Errorf(msg))
		panic(msg)
	}

	tx := gormMigrate.Begin()
	if err := gormMigrate.AutoMigrate(getMigrateTables()...); err != nil {
		tx.Rollback()
		msg := "[DB] unable to migrate " + err.Error()
		log.Logger().Error(fmt.Errorf(msg))
		panic(msg)
	}
	tx.Commit()
}
