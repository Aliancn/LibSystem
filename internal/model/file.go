package model

import "gorm.io/gorm"

type File struct {
	ID       uint   `gorm:"primaryKey"`
	FilePath string `gorm:"size:255;not null"`
}

func (f *File) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}
