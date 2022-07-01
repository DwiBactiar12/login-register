package routes

import (
	"net/http"
	user "test/delivery/controllers/userController"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo, connUser user.Usercontroller) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time:${time_rfc3339}, method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Selamat Datang")
	})

	e.POST("/loginadmin", connUser.LoginAdmin())
	e.POST("/login", connUser.Login())
	e.POST("/registeradmin", connUser.RegisterAdmin())

	e.GET("/search",connUser.Search())

	user := e.Group("/users", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	user.POST("", connUser.RegisterUser())
	user.GET("", connUser.GetAllUser())
	user.GET("/me", connUser.GetUserPrivat())
	user.GET("/:id", connUser.GetUser())
	user.PUT("/:id", connUser.UpdateUser())
	user.DELETE("/:id", connUser.DeleteUser())


	admin := e.Group("/admin", middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("$p4ssw0rd")}))
	admin.GET("/me", connUser.GetadminPrivat())
}
