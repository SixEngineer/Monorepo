package handler

import (
	"net/http"
	"openbridge/backend/internal/domain/entity"
	"openbridge/backend/internal/pkg/myerror"
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

// 上传访问令牌和刷新令牌
func (h *TokenHandler) UploadToken(c *gin.Context) {

	// 绑定JSON数据到Token结构体
	var token entity.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, tool.HttpResult{Code: myerror.ErrorCodeJsonFormatInvalid, Message: err.Error()})
		return
	}

	// 调用UseCase层上传令牌
	if err := h.tokenUsecase.UploadToken(token); err != nil {
		c.JSON(http.StatusInternalServerError, tool.HttpResult{Code: myerror.ErrorCodeTokenUploadFailed, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tool.HttpResult{}.Success(nil))

}
