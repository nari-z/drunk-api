package host

import (
	"os"

	"github.com/go-openapi/runtime/middleware"

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

	// add static file root(/LiquorImage)
	// // TODO: ディレクトリの設定を全体と共有したい。
	api.GetLiquorImageFileNameHandler = operations.GetLiquorImageFileNameHandlerFunc(func(params operations.GetLiquorImageFileNameParams) middleware.Responder {
		// return middleware.NotImplemented("operation GetLiquorImageFileName has not yet been implemented")
		// TODO: 暫定。ディレクトリ追加するだけで可能としたい。
		//       like a `e.Static("/LiquorImage", "LiquorImage")`

		file, err := os.Open("./LiquorImage/" + params.FileName)
		if err != nil {
			return operations.NewGetLiquorImageFileNameBadRequest()
		}

		return operations.NewGetLiquorImageFileNameOK().WithPayload(file)
	})
}
