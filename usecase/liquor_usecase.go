package usecase

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/nari-z/drunk-api/domain/model"
	"github.com/nari-z/drunk-api/domain/repository"
)

// LiquorUseCase is usecase for liquor model.
type LiquorUseCase interface {
	GetLiquorList(ctx context.Context) ([]*model.Liquor, error)
	RegistLiquor(ctx context.Context, liquorName string, liquorImage *multipart.FileHeader) (*model.Liquor, error)
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

func (l *liquorUseCase) RegistLiquor(ctx context.Context, liquorName string, liquorImage *multipart.FileHeader) (*model.Liquor, error) {
	src, err := liquorImage.Open()
	if err != nil {
		fmt.Println("upload_file.Open Error.")
		return nil, err
	}
	defer src.Close()

	// test code.
	var imageFile *model.ImageFile
	imageFile, err = l.ImageFileRepository.Create(ctx, liquorImage.Filename, src)
	if err != nil {
		fmt.Println("save Error.")
		fmt.Println(err.Error())
		return nil, err
	}

	// TODO: ファイルパスやら何やらの情報を元にentityを作成
	var newLiquor *model.Liquor = &model.Liquor{}
	newLiquor.Name = liquorName
	newLiquor.ImageFilePath = imageFile.FilePath
	// newLiquor.ImageFilePath, err = l.ImageFileRepository.Save(ctx, imageFile);
	// if err != nil {
	//     fmt.Println("save Error.");
	//     fmt.Println(err.Error());
	//     return nil, err;
	// }
	// newLiquor.ImageFilePath = l.createImageName(newLiquor);

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
