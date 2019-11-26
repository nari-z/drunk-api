package bundler

import (
	"github.com/nari-z/drunk-api/presenter/handler"
)

type HahdleBundler struct {
	UserHandler   handler.UserHandler
	LiquorHandler handler.LiquorHandler
}

func NewHandleBundler(u *UseCaseBundler) *HahdleBundler {
	var h *HahdleBundler = &HahdleBundler{}

	h.UserHandler = handler.NewUserHandler()
	h.LiquorHandler = handler.NewLiquorHandler(u.LiquorUseCase)

	return h
}
