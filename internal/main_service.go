package internal

import (
	"graphql_go/config"

	echo "github.com/labstack/echo/v4"
)

type Service struct {
	config     *config.Config
	httpServer *echo.Echo
}

func InitService(conf *config.Config) *Service {
	echoServer := echo.New()

	if _, e := conf.DBConn.DB(); e != nil {
		panic("[DB] unable to connect " + e.Error())
	}

	return &Service{
		config:     conf,
		httpServer: echoServer,
	}
}
