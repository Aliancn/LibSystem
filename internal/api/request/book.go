package request

//
//type Book struct {
//	ID          uint   `gorm:"primaryKey"`
//	Title       string `gorm:"size:255;not null"`
//	Author      string `gorm:"size:255;not null"`
//	Publisher   string `gorm:"size:255"`
//	Year        int    `gorm:"not null"`
//	Genre       string `gorm:"size:100"`
//	Status      string `gorm:"size:50;not null"` // e.g., available, borrowed
//	Location    string `gorm:"size:100"`
//	BorrowTimes int    `gorm:"not null"`
//	CreatedAt   time.Time
//	UpdatedAt   time.Time
//}

type BookDTO struct {
	BookID    uint   `json:"book_id"`
	Title     string `json:"title" binding:"required"`
	Author    string `json:"author" binding:"required"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year" binding:"required"`
	Genre     string `json:"genre"`
	Status    string `json:"status"`
	Location  string `json:"location"`
}
