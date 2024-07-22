package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"

	_ "ego/docs"
)

type Say struct {
}

func NewSay() Say {
	return Say{}
}

// Hello godoc
//
//	@Summary		Say Hello
//	@Description	says hello to the user
//	@Tags			say
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SingleResp
//	@Failure		500	{object}	SingleResp
//	@Router			/hello [get]
func (h Say) Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, SingleResp{Data: "Hello!"})
}

// Bye godoc
//
//	@Summary		Say Bye
//	@Description	says bye to the user
//	@Tags			say
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SingleResp
//	@Failure		500	{object}	SingleResp
//	@Router			/bye [get]
func (h Say) Bye(c echo.Context) error {
	return c.JSON(http.StatusOK, SingleResp{Data: "Bye!"})
}

// Damn godoc
//
//	@Summary		Say Damn
//	@Description	says damn to the user
//	@Tags			say
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	SingleResp
//	@Failure		500	{object}	SingleResp
//	@Router			/damn [get]
func (h Say) Damn(c echo.Context) error {
	return echo.
		NewHTTPError(http.StatusInternalServerError, SingleResp{Error: []string{"I`m angry"}}).
		WithInternal(errors.New("i'm angry internally"))
}
