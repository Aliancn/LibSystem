package model

import (
	"LibSystem/common"
	"gorm.io/gorm"
	"time"
)

// Borrowing model for 借阅管理
type Borrow struct {
	ID         uint      `gorm:"primaryKey"`
	BorrowDate time.Time `gorm:"not null"`
	ReturnDate time.Time
	Status     string `gorm:"size:50;not null"` // e.g., borrowed, returned
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// Relationships
	BookID int
	UserID int
}

func (b *Borrow) BeforeCreate(tx *gorm.DB) (err error) {
	b.BorrowDate = time.Now()
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	b.Status = common.StatusBorrowed
	return nil
}
