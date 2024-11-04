package request

type PaperDTO struct {
	PaperID    uint   `json:"paper_id"`
	Title      string `json:"title" `
	Author     string `json:"author" `
	Department string `json:"department" `
	Year       int    `json:"year" `
	Status     string `json:"status"`
	FilePath   string `json:"file_path"`
}

// Paper model for 论文管理
//type Paper struct {
//	ID            uint   `gorm:"primaryKey"`
//	Title         string `gorm:"size:255;not null"`
//	Author        string `gorm:"size:255;not null"`
//	Department    string `gorm:"size:100"`
//	Year          int    `gorm:"not null"`
//	Status        string `gorm:"size:50;not null"` // e.g., available, archived
//	DownloadTimes int    `gorm:"not null"`
//	FilePath      string `gorm:"size:255;not null"`
//	CreatedAt     time.Time
//	UpdatedAt     time.Time
//}
