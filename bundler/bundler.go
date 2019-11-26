package bundler

import (
	"github.com/jinzhu/gorm"
)

// Bundler is all bundle.
type Bundler struct {
	DBConn     *gorm.DB
	Handle     *HahdleBundler
	UseCase    *UseCaseBundler
	Repository *RepositoryBundler
}

// NewBundler return *Bundler.
func NewBundler(conn *gorm.DB) *Bundler {
	b := &Bundler{}

	b.DBConn = conn
	b.Repository = NewRepositoryBundler(b.DBConn)
	b.UseCase = NewUseCaseBundler(b.Repository)
	b.Handle = NewHandleBundler(b.UseCase)

	return b
}
