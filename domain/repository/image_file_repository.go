package repository

import (
	"context"
	"io"

	"github.com/nari-z/drunk-api/domain/model"
)

type ImageFileRepository interface {
	Create(ctx context.Context, fileName string, reader io.Reader) (*model.ImageFile, error)
	Delete(ctx context.Context, filePath string) error
	Read(ctx context.Context, filePath string) (*model.ImageFile, error)
	Exists(ctx context.Context, filePath string) bool
	ToBase64(imageFile *model.ImageFile) (string, error)
}