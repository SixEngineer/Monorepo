package repository

import (
	"openbridge/backend/internal/domain/entity"

	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

// 构造函数
func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

// 上传访问令牌和刷新令牌
func (repo *TokenRepository) InsertToken(token entity.Token) error {
    return repo.db.Create(&token).Error
}
