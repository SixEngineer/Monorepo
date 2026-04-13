package entity

import "time"

type ProviderAccount struct {
	ID              uint          `gorm:"column:id;primaryKey"`
	Name            string        `gorm:"column:name;size:100;not null"`
	ProviderType    string        `gorm:"column:provider_type;size:50;not null;index"`
	NetDisk         string        `gorm:"column:net_disk;size:50;not null"`
	AccountID       string        `gorm:"column:account_id;size:100"`
	Status          string        `gorm:"column:status;size:20;not null;default:'active'"`
	AccessToken     string        `gorm:"column:access_token;type:text"`
	RefreshToken    string        `gorm:"column:refresh_token;type:text"`
	TokenExpiresAt  *time.Time    `gorm:"column:token_expires_at"`
	TotalQuota      int64         `gorm:"column:total_quota"`
	UsedQuota       int64         `gorm:"column:used_quota"`
	AvailableQuota  int64         `gorm:"column:available_quota"`
	LastQuotaSyncAt *time.Time    `gorm:"column:last_quota_sync_at"`
	LastError       string        `gorm:"column:last_error;type:text"`
	CreatedAt       time.Time     `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time     `gorm:"column:updated_at;autoUpdateTime"`
}