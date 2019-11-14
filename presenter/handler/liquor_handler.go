package handler

import (
	"fmt"
	"net/http"
	"context"
	"mime/multipart"
	"github.com/labstack/echo"

	"github.com/nari-z/drunk-api/usecase"
)

type LiquorHandler interface {
	GetLiquorList(c echo.Context) error
	RegistLiquor(c echo.Context) error
}

type liquorHandler struct {
	LiquorUseCase usecase.LiquorUseCase
}

func NewLiquorHandler(u usecase.LiquorUseCase) LiquorHandler {
	return &liquorHandler{u};
}

func (h *liquorHandler) GetLiquorList(c echo.Context) error {
	ctx := c.Request().Context();
	if ctx == nil {
		ctx = context.Background();
	}

	fmt.Println("LiquorHandler.GetLiquorList.");

	liquorList, err := h.LiquorUseCase.GetLiquorList(ctx);
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, liquorList);
}

func (h *liquorHandler) RegistLiquor(c echo.Context) error {
	fmt.Println("In RegistLiquor.");
	ctx := c.Request().Context();
	if ctx == nil {
		ctx = context.Background();
	}

	var upload_file *multipart.FileHeader;
	var err error;
    upload_file, err = c.FormFile("image")
    if err != nil {
		fmt.Println("FormFile Error.");
        return err;
	}

	liquorName := c.FormValue("name");
	fmt.Println(liquorName);

	liquor, err := h.LiquorUseCase.RegistLiquor(ctx, liquorName, upload_file);
	if err != nil {
		return err;
	}

	return c.JSON(http.StatusOK, liquor);
}
