package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo"
	"mime/multipart"
	"net/http"

	"github.com/nari-z/drunk-api/usecase"
)

// LiquorHandler is request handler for liquor.
type LiquorHandler interface {
	GetLiquorList(c echo.Context) error
	RegistLiquor(c echo.Context) error
}

type liquorHandler struct {
	LiquorUseCase usecase.LiquorUseCase
}

// NewLiquorHandler return LiquorHandler。
func NewLiquorHandler(u usecase.LiquorUseCase) LiquorHandler {
	return &liquorHandler{u}
}

func (h *liquorHandler) GetLiquorList(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	fmt.Println("LiquorHandler.GetLiquorList.")

	liquorList, err := h.LiquorUseCase.GetLiquorList(ctx)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, liquorList)
}

func (h *liquorHandler) RegistLiquor(c echo.Context) error {
	fmt.Println("In RegistLiquor.")
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var uploadFile *multipart.FileHeader
	var err error
	uploadFile, err = c.FormFile("image")
	if err != nil {
		fmt.Println("FormFile Error.")
		return err
	}

	liquorName := c.FormValue("name")
	fmt.Println(liquorName)

	liquor, err := h.LiquorUseCase.RegistLiquor(ctx, liquorName, uploadFile)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, liquor)
}
