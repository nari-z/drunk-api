package host

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/nari-z/drunk-api/bundler"
)

func NewRouter(e *echo.Echo, h *bundler.HahdleBundler) {
	// CORS
	// TODO: filter domain.
	e.Use(middleware.CORS())

	// routing
	e.GET("/liquor", h.LiquorHandler.GetLiquorList)
	e.POST("/liquor", h.LiquorHandler.RegistLiquor)
}
