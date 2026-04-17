package entity

import "time"

type DownloadTask struct {
	ID                uint   `gorm:"primaryKey"`
	TaskID            string `gorm:"size:64;uniqueIndex;not null"`
	SourceURL         string `gorm:"type:text;not null"`
	SourceType        string `gorm:"size:50"`
	FileName          string `gorm:"size:255"`
	FileSize          int64
	DirectLink        string  `gorm:"type:text"`
	ProviderAccountID *uint   `gorm:"index"`
	ProviderType      string  `gorm:"size:50;index"`
	Aria2GID          string  `gorm:"size:64;index"`
	Status            string  `gorm:"size:30;not null;index"`
	Progress          float64 `gorm:"default:0"`
	ErrorMessage      string  `gorm:"type:text"`
	RetryCount        int     `gorm:"default:0"`
	StartedAt         *time.Time
	FinishedAt        *time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}