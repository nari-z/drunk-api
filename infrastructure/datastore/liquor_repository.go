package datastore

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/nari-z/drunk-api/domain/model"
	"github.com/nari-z/drunk-api/domain/repository"
)

type liquorRepository struct {
	DBConn *gorm.DB
}

// NewLiquorRepository return repository.LiquorRepository.
func NewLiquorRepository(conn *gorm.DB) repository.LiquorRepository {
	return &liquorRepository{conn}
}

func (r *liquorRepository) Fetch(ctx context.Context) ([]*model.Liquor, error) {
	var liquorList []*model.Liquor
	var err error

	err = r.DBConn.Order("id desc").Find(&liquorList).Error

	if err != nil {
		fmt.Println(err.Error())
	}

	return liquorList, err
}

func (r *liquorRepository) FindByID(ctx context.Context, id uint64) (*model.Liquor, error) {
	var liquor model.Liquor
	var err error

	err = r.DBConn.Where("id = ?", id).First(&liquor).Error

	return &liquor, err
}

func (r *liquorRepository) Create(ctx context.Context, l *model.Liquor) (*model.Liquor, error) {
	err := r.DBConn.Create(l).Error
	return l, err
}

func (r *liquorRepository) Update(ctx context.Context, l *model.Liquor) (*model.Liquor, error) {
	err := r.DBConn.Model(l).Update(l).Error
	return l, err
}

func (r *liquorRepository) Delete(ctx context.Context, id int) error {
	// TODO:
	// l := &model.Liquor{ID: id};
	// err := r.DBConn.Delete(l).Error;
	// return err;

	return nil
}
