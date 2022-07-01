package main

import (
	"fmt"
	"test/config"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	config.AutoMigrate(db)
	e := echo.New()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
