package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type UserHandler interface {
	GetUsers(c echo.Context) error
}

type userHandler struct {
}

func NewUserHandler() UserHandler {
	return &userHandler{}
}

func (h *userHandler) GetUsers(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	fmt.Println("UserHandler.GetUsers.")
	var users string
	users = "{user: UserA}"
	return c.JSON(http.StatusOK, users)
}
