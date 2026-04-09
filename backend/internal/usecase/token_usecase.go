package usecase

import (
	"openbridge/backend/internal/domain/entity"
	"openbridge/backend/internal/repository"
)

type TokenUseCase struct {
	TokenRepo *repository.TokenRepository
}

// 构造函数
func NewTokenUseCase(tokenRepo *repository.TokenRepository) *TokenUseCase {
	return &TokenUseCase{
		TokenRepo: tokenRepo,
	}
}

// 上传访问令牌和刷新令牌
func (t *TokenUseCase) UploadToken(token entity.Token) error {
	return t.TokenRepo.InsertToken(token)
}
