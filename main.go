package main

import (
	"fmt"
	"sync"

	dotenv "github.com/joho/godotenv"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"graphql_go/config"
	"graphql_go/src/modules"
)

func main() {
	err := dotenv.Load(".env")
	if err != nil {
		panic(".env is not loaded properly")
	}
	cfg := config.Init()

	e := echo.New()
	e.Use(middleware.CORS())
	group := e.Group("")

	service := modules.NewService(cfg)
	for _, mod := range service.GetModules() {
		if h := mod.RESTHandler(); h != nil {
			h.Mount(group)
		}
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		e.Start(fmt.Sprintf(":%v", config.GlobalEnv.HTTPPort))
	}()

	wg.Wait()
}
