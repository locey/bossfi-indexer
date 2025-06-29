package model

import "gorm.io/gorm"

type Model[T any] interface {
	GetById(id int64) (*T, error)
	Create(v *T) error
	UpdateById(v *T) error
	DeleteById(id int64) error
}

func NotDeleted(db *gorm.DB) *gorm.DB {
	return db.Where("deleted = false")
}
