package main

import (
	"fmt"
	"test/config"
	"test/delivery/controllers/userController"
	"test/delivery/routes"
	"test/repository/userRepo"

	"github.com/labstack/echo/v4"
)

func main() {
	conf := config.InitConfig()
	db := config.InitDB(*conf)
	// aktifkan jika ingin membuat table
	config.AutoMigrate(db)
	e := echo.New()

	userRepo:=userrepo.NewUserRepository(db)
	userCont:=usercontroller.NewUserControllers(userRepo)

	routes.Route(e,userCont)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
