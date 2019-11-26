package model

import (
	"time"
)

// DefaultModel is DB default property.// TODO: ここにDBの知識を入れたくない
type DefaultModel struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}
