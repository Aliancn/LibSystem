package model

import "time"

type Info struct {
	ID           int       `json:"id" gorm:"primary_key"`
	PaperID      int       `json:"download"`
	DownloadTime time.Time `json:"download_time"`
}
