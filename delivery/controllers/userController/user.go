package usercontroller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"test/delivery/middleware"
	request "test/delivery/view/request"
	"test/delivery/view/response"
	"test/entity"
	user "test/repository/userRepo"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserControll struct {
	userRepo user.UserRepo
	Validate *validator.Validate
}

func NewUserControllers(userRepo user.UserRepo) *UserControll {
	return &UserControll{
		userRepo: userRepo,
		Validate: validator.New(),
	}
}

func (uc *UserControll) RegisterAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request request.InsertUser

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := uc.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
		user := entity.Admin{
			Nama:     request.Nama,
			Email:    request.Email,
			Password: request.Password,
			HP:       request.HP,
			Role:     "admin",
			Umur:     request.Umur,
			Gender:   request.Gender,
		}

		result, err := uc.userRepo.InsertAdmin(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("success register admin!", result))
	}
}

func (uc *UserControll) RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, Role := middleware.ExtractToken(c)
		if Role != "admin" {
			return c.JSON(http.StatusForbidden, response.StatusUnauthorized(errors.New("anda tidak di izinkan mengakses halama ini")))
		}
		var request request.InsertUser

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := uc.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
		user := entity.User{
			Nama:     request.Nama,
			Email:    request.Email,
			Password: request.Password,
			HP:       request.HP,
			Role:     "user",
			Umur:     request.Umur,
			Gender:   request.Gender,
			AdminID:  uint(id),
		}

		result, err := uc.userRepo.InsertUser(user)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusCreated, response.StatusCreated("success register User!", result))
	}
}

func (uc *UserControll) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request request.Login

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := uc.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
		user, err := uc.userRepo.Login(request)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.StatusUnauthorized(err))
		}

		result := response.LoginDetail{User: response.Login{
			ID:    user.ID,
			Nama:  user.Nama,
			Email: user.Email,
		}}
		if result.Token == "" {
			token, _ := middleware.CreateToken(user.ID, user.Role)
			result.Token = token
		}

		return c.JSON(http.StatusOK, response.StatusOK("success login!", result))
	}
}

func (uc *UserControll) LoginAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request request.Login

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusInvalidRequest())
		}

		if err := uc.Validate.Struct(request); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
		user, err := uc.userRepo.LoginAdmin(request)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, response.StatusUnauthorized(err))
		}

		result := response.LoginDetail{User: response.Login{
			ID:    user.ID,
			Nama:  user.Nama,
			Email: user.Email,
		}}
		fmt.Println(user.Role)
		if result.Token == "" {
			token, _ := middleware.CreateToken(user.ID, user.Role)
			result.Token = token
		}

		return c.JSON(http.StatusOK, response.StatusOK("success login!", result))
	}
}

func (uc *UserControll) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		result, err := uc.userRepo.GetUser(uint(id))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusOK, response.StatusOK("success get User!", result))
	}
}

func (uc *UserControll) GetadminPrivat() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := middleware.ExtractToken(c)
		result, err := uc.userRepo.GetAdmin(uint(id))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusOK, response.StatusOK("success get User!", result))
	}
}

func (uc *UserControll) GetUserPrivat() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := middleware.ExtractToken(c)

		result, err := uc.userRepo.GetUser(uint(id))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusOK, response.StatusOK("success get User!", result))
	}
}

// get All
func (uc *UserControll) GetAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		offsite := 0
		if page >= 1 {
			offsite = (4 * (page - 1)) - 1

		}

		result, err := uc.userRepo.GetAllUser(offsite)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, response.StatusNotFound("Room By ID Not Found"))
		}
		return c.JSON(http.StatusOK, response.StatusOK("Success Get All user", result))
	}
}
func (uc *UserControll) Search() echo.HandlerFunc {
	return func(c echo.Context) error {
		title := c.QueryParam("title")
		page, _ := strconv.Atoi(c.QueryParam("page"))
		offsite := 0
		if page >= 1 {
			offsite = (4 * (page - 1)) - 1

		}

		result, err := uc.userRepo.Search(title,offsite)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, response.StatusNotFound("Room By ID Not Found"))
		}
		return c.JSON(http.StatusOK, response.StatusOK("Success Get All user", result))
	}
}

func (uc *UserControll) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, Role := middleware.ExtractToken(c)
		if Role != "admin" {
			return c.JSON(http.StatusForbidden, response.StatusUnauthorized(errors.New("anda tidak di izinkan mengakses halama ini")))
		}

		id, _ := strconv.Atoi(c.Param("id"))

		var update request.UpdateUser
		err := c.Bind(&update)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
		if err := uc.Validate.Struct(update); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
		result, err := uc.userRepo.GetUser(uint(id))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(errors.New("user tidak ditemukan")))
		}

		result, err = uc.userRepo.UpdateUser(uint(id), update)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}

		return c.JSON(http.StatusOK, response.StatusOK("Success Update user", result))
	}
}

func (uc *UserControll) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, Role := middleware.ExtractToken(c)
		if Role != "admin" {
			return c.JSON(http.StatusForbidden, response.StatusUnauthorized(errors.New("anda tidak di izinkan mengakses halama ini")))
		}

		id, _ := strconv.Atoi(c.Param("id"))

		var update request.UpdateUser
		err := c.Bind(&update)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
		if err := uc.Validate.Struct(update); err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequest(err))
		}
		result, err := uc.userRepo.GetUser(uint(id))

		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(errors.New("user tidak ditemukan")))
		}

		err = uc.userRepo.DeleteUser(uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.StatusBadRequestDuplicate(err))
		}
		result = response.User{}

		return c.JSON(http.StatusOK, response.StatusOK("Success delete User", result))
	}
}
