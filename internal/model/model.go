package model

import "time"

// ShelfBook join table model for the many-to-many relationship
type ShelfBook struct {
	ShelfID uint `gorm:"primaryKey"`
	BookID  uint `gorm:"primaryKey"`
	Count   int  `gorm:"not null"` // Number of copies of a book on this shelf
}

// Statistic model for 信息统计
type Statistic struct {
	ID        uint      `gorm:"primaryKey"`
	Type      string    `gorm:"size:50;not null"` // e.g., borrowing, downloads
	ItemID    uint      `gorm:"not null"`         // Foreign key to either Book or Thesis
	Count     int       `gorm:"not null"`
	Date      time.Time `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Shelf model for 书架管理
type Shelf struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"size:100;not null"` // Name or label for the shelf
	QuantityLimit int    `gorm:"not null"`          // Maximum allowed quantity of books on the shelf
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
