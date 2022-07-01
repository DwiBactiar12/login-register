package usercontroller

import "github.com/labstack/echo/v4"

type Usercontroller interface {
	RegisterAdmin() echo.HandlerFunc
	RegisterUser() echo.HandlerFunc
	Login() echo.HandlerFunc
	LoginAdmin() echo.HandlerFunc
	GetUser() echo.HandlerFunc
	GetadminPrivat() echo.HandlerFunc
	GetUserPrivat() echo.HandlerFunc
	GetAllUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	Search() echo.HandlerFunc
}
