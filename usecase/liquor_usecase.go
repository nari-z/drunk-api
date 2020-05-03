package usecase

import (
	"context"
	"fmt"
	"io"

	"github.com/nari-z/drunk-api/domain/model"
	"github.com/nari-z/drunk-api/domain/repository"
)

// LiquorUseCase is usecase for liquor model.
type LiquorUseCase interface {
	GetLiquorList(ctx context.Context) ([]*model.Liquor, error)
	RegistLiquor(ctx context.Context, liquorName string, fileName string, liquorImage io.Reader) (*model.Liquor, error)
	GetLiquor(ctx context.Context, liquorID uint64) (*model.Liquor, error)
	GetImageFormatAndBase64Data(ctx context.Context, liquorID uint64) (string, string, error)
	GetNameAndImageFilePath(ctx context.Context, liquorID uint64) (string, string, error)
}

type liquorUseCase struct {
	repository.LiquorRepository
	repository.ImageFileRepository
}

// NewLiquorUseCase return LiquorUseCase.
func NewLiquorUseCase(l repository.LiquorRepository, i repository.ImageFileRepository) LiquorUseCase {
	return &liquorUseCase{l, i}
}

func (l *liquorUseCase) GetLiquorList(ctx context.Context) ([]*model.Liquor, error) {
	fmt.Println("LiquorUseCase.GetLiquorList().")

	return l.LiquorRepository.Fetch(ctx)
}

func (l *liquorUseCase) GetLiquor(ctx context.Context, liquorID uint64) (*model.Liquor, error) {
	fmt.Println("LiquorUseCase.GetLiquor().")

	return l.LiquorRepository.FindByID(ctx, liquorID)
}

func (l *liquorUseCase) RegistLiquor(ctx context.Context, liquorName string, fileName string, liquorImage io.Reader) (*model.Liquor, error) {
	// create image file
	var imageFile *model.ImageFile
	imageFile, err := l.ImageFileRepository.Create(ctx, fileName, liquorImage)
	if err != nil {
		fmt.Println("save Error.")
		fmt.Println(err.Error())
		return nil, err
	}

	// create model
	var newLiquor *model.Liquor = &model.Liquor{ Name: liquorName, ImageFilePath: imageFile.FilePath}

	return l.LiquorRepository.Create(ctx, newLiquor)
}

func (l *liquorUseCase) GetImageFormatAndBase64Data(ctx context.Context, liquorID uint64) (string, string, error) {
	fmt.Println("LiquorUseCase.GetImageFormatAndBase64Data().")

	var liquor *model.Liquor
	var err error
	liquor, err = l.GetLiquor(ctx, liquorID)
	if err != nil {
		return "", "", err
	}

	var imageFile *model.ImageFile
	imageFile, err = l.ImageFileRepository.Read(ctx, liquor.ImageFilePath)
	if err != nil {
		return "", "", err
	}

	var imageData string
	imageData, err = l.ImageFileRepository.ToBase64(imageFile)
	if err != nil {
		return "", "", err
	}

	return imageFile.ImageFormat, imageData, nil
}

func (l *liquorUseCase) GetNameAndImageFilePath(ctx context.Context, liquorID uint64) (string, string, error) {
	var liquor *model.Liquor
	var err error
	liquor, err = l.GetLiquor(ctx, liquorID)
	if err != nil {
		return "", "", err
	}

	// TODO: Read必要？
	var imageFile *model.ImageFile
	imageFile, err = l.ImageFileRepository.Read(ctx, liquor.ImageFilePath)
	if err != nil {
		return "", "", err
	}

	return liquor.Name, imageFile.FilePath, nil
}
