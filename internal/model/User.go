package model

import (
	"LibSystem/common"
	"gorm.io/gorm"
	"time"
)

// User model for 用户管理
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"size:255;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	Name      string `gorm:"size:255;"`
	Role      string `gorm:"size:50;not null"` // e.g., student, professor
	Email     string `gorm:"size:100"`
	Phone     string `gorm:"size:20"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	if u.Role == "" {
		u.Role = common.RoleUser
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	// 在更新记录千自动填充更新时间
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) AfterFind(tx *gorm.DB) error {
	return nil
}
