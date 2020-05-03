package host

import (
	"github.com/nari-z/drunk-api/bundler"

	"github.com/nari-z/drunk-api/generate/restapi/operations"
)

// NewRouter sets routing.
func NewRouter(api *operations.DrunkAPI, h *bundler.HahdleBundler) {
	// TODO: CORS
	// TODO: filter domain.
	// e.Use(middleware.CORS())

	// routing
	api.AddLiquorHandler = operations.AddLiquorHandlerFunc(h.LiquorHandler.RegistLiquor)
	api.GetLiquorsHandler = operations.GetLiquorsHandlerFunc(h.LiquorHandler.GetLiquorList)

	// TODO: add static file root(/LiquorImage)
	// // TODO: ディレクトリの設定を全体と共有したい。
	// e.Static("/LiquorImage", "LiquorImage")
}
