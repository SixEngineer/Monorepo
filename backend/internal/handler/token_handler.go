package handler

import (
	"net/http"
	"openbridge/backend/internal/domain/entity"
	"openbridge/backend/internal/tool"
	"openbridge/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

// 令牌处理接口
type TokenHandler struct {
	tokenUsecase *usecase.TokenUseCase
}

// 构造函数
func NewTokenHandler(tokenUsecase *usecase.TokenUseCase) *TokenHandler {
	return &TokenHandler{
		tokenUsecase: tokenUsecase,
	}
}

// 上传访问令牌
func (b *TokenHandler) UploadAccessToken(c *gin.Context) {
	
	// 绑定请求体
	var accessTokenReq entity.AccessToken
	if err := c.ShouldBindJSON(&accessTokenReq); err != nil {
		c.JSON(http.StatusBadRequest, tool.HttpResult{Code: 1001, Message: err.Error()})
		return
	}

	// 保存令牌
	if err := b.tokenUsecase.UploadBaiduAccessToken(accessTokenReq); err != nil {
		c.JSON(http.StatusInternalServerError, tool.HttpResult{Code: 1001, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tool.HttpResult{Code: 0, Message: "ok"})

}

// 上传刷新令牌
func (b *TokenHandler) UploadRefreshToken(c *gin.Context) {
    
	// 绑定请求体
	var refreshTokenReq entity.RefreshToken
	if err := c.ShouldBindJSON(&refreshTokenReq); err != nil {
		c.JSON(http.StatusBadRequest, tool.HttpResult{Code: 1001, Message: err.Error()})
		return
	}

	// 保存令牌
	if err := b.tokenUsecase.UploadBaiduRefreshToken(refreshTokenReq); err != nil {
	    c.JSON(http.StatusInternalServerError, tool.HttpResult{Code: 1001, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tool.HttpResult{Code: 0, Message: "ok"})
}
