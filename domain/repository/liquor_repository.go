package repository

import (
	"context"

	"github.com/nari-z/drunk-api/domain/model"
)

type LiquorRepository interface {
	Fetch(ctx context.Context) ([]*model.Liquor, error)
	Create(ctx context.Context, liquor *model.Liquor) (*model.Liquor, error)
	Update(ctx context.Context, liquor *model.Liquor) (*model.Liquor, error)
	Delete(ctx context.Context, id int) error
	FindByID(ctx context.Context, id uint64) (*model.Liquor, error)
}