package bundler

import (
	"github.com/nari-z/drunk-api/usecase"
)

// UseCaseBundler is usecase bundle.
type UseCaseBundler struct {
	LiquorUseCase usecase.LiquorUseCase
}

// NewUseCaseBundler return *UseCaseBundler.
func NewUseCaseBundler(r *RepositoryBundler) *UseCaseBundler {
	u := &UseCaseBundler{}
	u.LiquorUseCase = usecase.NewLiquorUseCase(r.LiquorRepository, r.ImageFileRepository)

	return u
}
