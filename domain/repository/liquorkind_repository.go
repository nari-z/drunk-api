package repository

import (
	"context"

	"github.com/nari-z/drunk-api/domain/model"
)

type LiquorKindRepository interface {
	Fetch(ctx context.Context) ([]*model.LiquorKind, error)
	Create(ctx context.Context, liquorKind *model.LiquorKind) (*model.LiquorKind, error)
	Update(ctx context.Context, liquorKind *model.LiquorKind) (*model.LiquorKind, error)
	Delete(ctx context.Context, id int) error
}