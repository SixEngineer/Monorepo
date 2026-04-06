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

// 上传访问令牌
func (t *TokenUseCase) UploadBaiduAccessToken(accessToken entity.AccessToken) error {
	return t.TokenRepo.InsertBaiduAccessToken(accessToken)
}

// 上传刷新令牌
func (t *TokenUseCase) UploadBaiduRefreshToken(refreshToken entity.RefreshToken) error {
	return t.TokenRepo.InsertBaiduRefreshToken(refreshToken)
}