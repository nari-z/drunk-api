package bundler

import (
	"github.com/jinzhu/gorm"

	"github.com/nari-z/drunk-api/domain/repository"
	"github.com/nari-z/drunk-api/infrastructure/datastore"
)

type RepositoryBundler struct {
	LiquorRepository repository.LiquorRepository
	ImageFileRepository repository.ImageFileRepository
}

func NewRepositoryBundler(conn *gorm.DB) *RepositoryBundler {
	var r *RepositoryBundler = &RepositoryBundler{};

	r.LiquorRepository = datastore.NewLiquorRepository(conn);
	r.ImageFileRepository = datastore.NewImageFileReposiotry();

	return r;
}
