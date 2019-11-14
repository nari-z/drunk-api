package datastore

// TODO: 試行錯誤して汚いから後で整理。

import (
	"fmt"
	"context"
	"os"
    "io"
    "io/ioutil"
    "path/filepath"
    "time"
	"strings"
	"bytes"
	"encoding/base64"

	"errors"
	
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"

	"github.com/nari-z/drunk-api/domain/repository"
	"github.com/nari-z/drunk-api/domain/model"
)

type imageFileRepository struct {
    imageSaveDirectoryName string
}

func NewImageFileReposiotry() repository.ImageFileRepository {
	return &imageFileRepository{"LiquorImage"};
}

func (r *imageFileRepository) Create(ctx context.Context, fileName string, reader io.Reader) (*model.ImageFile, error) {
	var newImageFile *model.ImageFile = &model.ImageFile{};
	newImageFile.FilePath = r.createImageName(fileName);

	var err error;
	var f *os.File;
	f, err = os.Create(newImageFile.FilePath);
	if err != nil {
		return nil, err;
	}
	defer f.Close();

	// io.Readerのままだと複数回Readできないので、[]byteを作成し使い回す。
	newImageFile.Data, err = ioutil.ReadAll(reader);
	if err != nil {
		return nil, err;
	}

	_, err = io.Copy(f, bytes.NewBuffer(newImageFile.Data));
	if err != nil {
		return nil, err;
	}
	
	_, newImageFile.ImageFormat, err = image.Decode(bytes.NewBuffer(newImageFile.Data));
	if err != nil {
		r.Delete(ctx, newImageFile.FilePath);
		return nil, err;
	}
	fmt.Println(newImageFile.ImageFormat);
	fmt.Println(newImageFile.FilePath);

	// TODO: 保存処理を分けておくと、Readで一部処理を使いまわせるかも。
	return newImageFile, nil;
}

func (r *imageFileRepository) Delete(ctx context.Context, filePath string) error {
	return os.Remove(filePath);
}

func (r *imageFileRepository) Read(ctx context.Context, filePath string) (*model.ImageFile, error) {
	var err error;
	var f *os.File;
	// file read.
    f, err = os.Open(filePath);
    if err != nil {
		return nil, err;
    }
	defer f.Close();

	var imageFile *model.ImageFile = &model.ImageFile{};
	imageFile.FilePath = filePath;
	fmt.Println(filePath);

	// io.Readerのままだと複数回Readできないので、[]byteを作成し使い回す。
	imageFile.Data, err = ioutil.ReadAll(f);
	if err != nil {
		return nil, err;
	}

	_, imageFile.ImageFormat, err = image.Decode(bytes.NewBuffer(imageFile.Data));
	if err != nil {
		return nil, err;
	}

	return imageFile, nil;
}

func (r *imageFileRepository) Exists(ctx context.Context, filePath string) bool {
	// TODO: 
	return false;
}

func (l *imageFileRepository) ToBase64(imageFile *model.ImageFile) (string, error) {

    img, imgFmt, err := image.Decode(bytes.NewBuffer(imageFile.Data));
    if err != nil {
		return "", err;
    }

	buffer := new(bytes.Buffer);
    switch imgFmt {
    case "jpeg":
		err = jpeg.Encode(buffer, img, nil);
    case "gif":
        err = gif.Encode(buffer, img, nil);
    case "png":
        err = png.Encode(buffer, img);
	default:
		return "", errors.New("unknown image format.");
	}

	if err != nil {
		return "", err;
	}

	return base64.StdEncoding.EncodeToString(imageFile.Data), nil;
}

func (l *imageFileRepository) createImageName(fileName string) string {
	var saveFilename string = strings.Join([]string{time.Now().Format("2006-01-02T15:04:05Z07:00"), fileName}, "_");
	return filepath.Join(l.imageSaveDirectoryName, saveFilename);
}

