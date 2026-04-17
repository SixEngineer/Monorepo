package repository

import (
	"openbridge/backend/internal/domain/entity"

	"gorm.io/gorm"
)

type QuotaRepository struct {
	db *gorm.DB
}

// 构造函数
func NewQuotaRepository(db *gorm.DB) *QuotaRepository {
	return &QuotaRepository{db: db}
}

// 写入配额快照
func (repo *QuotaRepository) InsertQuotaSnapshot(snapshot *entity.QuotaSnapshot) error {
	return repo.db.Create(snapshot).Error
}
