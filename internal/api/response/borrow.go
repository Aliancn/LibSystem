package response

import "time"

type BorrowVO struct {
	ID         uint      `json:"borrow_id"`
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date"`
	Status     string    `json:"status"`
	Book       int       `json:"book_id"`
	User       int       `json:"user_id"`
}

//type Borrow struct {
//	ID         uint      `gorm:"primaryKey"`
//	BorrowDate time.Time `gorm:"not null"`
//	ReturnDate time.Time
//	Status     string `gorm:"size:50;not null"` // e.g., borrowed, returned
//	CreatedAt  time.Time
//	UpdatedAt  time.Time
//
//	// Relationships
//	Book Book `gorm:"foreignKey:BookID"`
//	User User `gorm:"foreignKey:UserID"`
//}
