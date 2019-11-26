package bundler

import (
	"github.com/nari-z/drunk-api/usecase"
)

type UseCaseBundler struct {
	LiquorUseCase usecase.LiquorUseCase
}

func NewUseCaseBundler(r *RepositoryBundler) *UseCaseBundler {
	var u *UseCaseBundler = &UseCaseBundler{}
	u.LiquorUseCase = usecase.NewLiquorUseCase(r.LiquorRepository, r.ImageFileRepository)

	return u
}
