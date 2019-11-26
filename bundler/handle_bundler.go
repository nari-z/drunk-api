package bundler

import (
	"github.com/nari-z/drunk-api/presenter/handler"
)

// HahdleBundler is handler bundle.
type HahdleBundler struct {
	UserHandler   handler.UserHandler
	LiquorHandler handler.LiquorHandler
}

// NewHandleBundler return *HahdleBundler.
func NewHandleBundler(u *UseCaseBundler) *HahdleBundler {
	h := &HahdleBundler{}

	h.UserHandler = handler.NewUserHandler()
	h.LiquorHandler = handler.NewLiquorHandler(u.LiquorUseCase)

	return h
}
