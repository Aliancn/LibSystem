package model

import (
	"LibSystem/common"
	"gorm.io/gorm"
	"time"
)

// Book model for 图书管理
type Book struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Author      string `gorm:"size:255;not null"`
	Publisher   string `gorm:"size:255"`
	Year        int    `gorm:"not null"`
	Genre       string `gorm:"size:100"`
	Status      string `gorm:"size:50;not null"` // e.g., available, borrowed
	Location    string `gorm:"size:100"`
	BorrowTimes int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	if b.Status == "" {
		b.Status = common.StatusAvailable
	}
	b.BorrowTimes = 0
	return nil
}

func (b *Book) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}

func (b *Book) AfterFind(tx *gorm.DB) error {
	return nil
}
