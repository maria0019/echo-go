package handler

import (
	"ego/internal/entity"
	"github.com/labstack/echo/v4"
	slogecho "github.com/samber/slog-echo"
	"log/slog"
	"net/http"
	"strconv"

	_ "ego/docs"
)

type User struct {
}

func NewUser() User {
	return User{}
}

// Get User godoc
//
//	@Summary		Get user
//	@Description	get user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param        	id   path      int  true  "User ID"
//	@Success		200	{object}	SingleResp
//	@Failure		500	{object}	SingleResp
//	@Router			/user/{id} [get]
func (h User) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.
			NewHTTPError(http.StatusInternalServerError, SingleResp{Error: []string{err.Error()}}).
			WithInternal(err)
	}

	slogecho.AddCustomAttributes(c, slog.Int("userId", id))

	return c.JSON(http.StatusOK, SingleResp{Data: entity.User{
		Id:        id,
		FirstName: "John",
		LastName:  "Smith",
	}})
}

// Create User godoc
//
//	@Summary		Create user
//	@Description	create user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param 			request body 	UserRequest true "user entity"
//	@Success 		200 {object} 	entity.User
//	@Success		200	{object}	SingleResp
//	@Failure		500	{object}	SingleResp
//	@Router			/user [post]
func (h User) Create(c echo.Context) error {
	u := new(UserRequest)
	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, SingleResp{Data: entity.User{
		Id:        101,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}})
}

// Update User godoc
//
//	@Summary		Update user
//	@Description	update user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param   		firstName     query     string     true  "string valid"  minlength(1)  maxlength(100)
//	@Param   		lastName      query     string     true  "string valid"  minlength(1)  maxlength(100)
//	@Success		200	{object}	SingleResp
//	@Failure		500	{object}	SingleResp
//	@Router			/user [put]
func (h User) Update(c echo.Context) error {
	name := c.FormValue("firstName")
	lastName := c.FormValue("lastName")

	return c.JSON(http.StatusOK, SingleResp{Data: entity.User{
		Id:        101,
		FirstName: name,
		LastName:  lastName,
	}})
}
