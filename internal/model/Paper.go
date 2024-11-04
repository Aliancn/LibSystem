package model

import (
	"LibSystem/common"
	"gorm.io/gorm"
	"time"
)

// Paper model for 论文管理
type Paper struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"size:255;not null"`
	Author        string `gorm:"size:255;not null"`
	Department    string `gorm:"size:100"`
	Year          int    `gorm:"not null"`
	Status        string `gorm:"size:50;not null"` // e.g., available, archived
	DownloadTimes int    `gorm:"not null"`
	FilePath      string `gorm:"size:255;not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (p *Paper) BeforeCreate(tx *gorm.DB) (err error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.Status = common.StatusAvailable
	p.DownloadTimes = 0
	return nil
}

func (p *Paper) BeforeUpdate(tx *gorm.DB) error {
	p.UpdatedAt = time.Now()
	return nil
}
