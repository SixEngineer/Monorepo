package entity

import "time"

type QuotaSnapshot struct {
	ID                uint      `gorm:"column:id;primaryKey" json:"id"`
	Provider          string    `gorm:"column:provider;size:100;not null;index" json:"provider"`
	ProviderAccountID uint      `gorm:"column:provider_account_id;not null;index" json:"provider_account_id"`
	Total             int64     `gorm:"column:total_quota" json:"total"`
	Used              int64     `gorm:"column:used_quota" json:"used"`
	Available         int64     `gorm:"column:available_quota" json:"available"`
	SyncedAt          time.Time `gorm:"column:synced_at;not null;index" json:"synced_at"`
	CreatedAt         time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
