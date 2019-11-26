package bundler

import (
	"github.com/jinzhu/gorm"

	"github.com/nari-z/drunk-api/domain/repository"
	"github.com/nari-z/drunk-api/infrastructure/datastore"
)

// RepositoryBundler is repository bundle.
type RepositoryBundler struct {
	LiquorRepository    repository.LiquorRepository
	ImageFileRepository repository.ImageFileRepository
}

// NewRepositoryBundler return *RepositoryBundler.
func NewRepositoryBundler(conn *gorm.DB) *RepositoryBundler {
	r := &RepositoryBundler{}

	r.LiquorRepository = datastore.NewLiquorRepository(conn)
	r.ImageFileRepository = datastore.NewImageFileReposiotry()

	return r
}
