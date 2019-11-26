package bundler

import (
	"github.com/jinzhu/gorm"
)

type Bundler struct {
	DBConn     *gorm.DB
	Handle     *HahdleBundler
	UseCase    *UseCaseBundler
	Repository *RepositoryBundler
}

func NewBundler(conn *gorm.DB) *Bundler {
	var b *Bundler = &Bundler{}

	b.DBConn = conn
	b.Repository = NewRepositoryBundler(b.DBConn)
	b.UseCase = NewUseCaseBundler(b.Repository)
	b.Handle = NewHandleBundler(b.UseCase)

	return b
}
