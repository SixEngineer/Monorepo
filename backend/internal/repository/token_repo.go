package repository

import (
	"openbridge/backend/internal/domain/entity"
	"openbridge/backend/internal/pkg/logger"

	"go.uber.org/zap"
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
	err := repo.db.Create(&token).Error
	if err != nil {
		logger.L().Error("insert token failed", zap.Error(err), zap.String("net_disk", token.NetDisk))
		return err
	}
	return nil
}
