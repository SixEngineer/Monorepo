package handler

import (
	"context"
	"errors"
	"net/http"
	"openbridge/backend/internal/pkg/logger"
	"openbridge/backend/internal/pkg/myerror"
	"openbridge/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type quotaRequest struct {
	Provider string `json:"name" binding:"required"`
}

type quotaResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type QuotaHandler struct {
	quotaUseCase *usecase.QuotaUseCase
}

// 构造函数
func NewQuotaHandler(quotaUseCase *usecase.QuotaUseCase) *QuotaHandler {
	return &QuotaHandler{quotaUseCase: quotaUseCase}
}

// QueryQuota 查询数据库中的当前 quota，不触发远端调用
func (h *QuotaHandler) QueryQuota(c *gin.Context) {
	var req quotaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, quotaResponse{Code: myerror.ErrorCodeJsonFormatInvalid, Message: err.Error(), Data: nil})
		c.Set(logger.LoggerErrorCodeKey, myerror.ErrorCodeJsonFormatInvalid)
		c.Set(logger.LoggerMessageKey, err.Error())
		return
	}

	quota, err := h.quotaUseCase.QueryQuota(req.Provider)
	if err != nil {
		status := http.StatusInternalServerError
		code := myerror.ErrorCodeQuotaQueryFailed
		if usecase.IsRecordNotFound(err) {
			status = http.StatusNotFound
			code = myerror.ErrorCodeProviderGetFailed
		}
		c.JSON(status, quotaResponse{Code: code, Message: err.Error(), Data: nil})
		c.Set(logger.LoggerErrorCodeKey, code)
		c.Set(logger.LoggerMessageKey, err.Error())
		return
	}

	c.JSON(http.StatusOK, quotaResponse{Code: 0, Message: "ok", Data: quota})
}

// SyncQuota 调用远端 GetQuota 同步并落库
func (h *QuotaHandler) SyncQuota(c *gin.Context) {
	var req quotaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, quotaResponse{Code: myerror.ErrorCodeJsonFormatInvalid, Message: err.Error(), Data: nil})
		c.Set(logger.LoggerErrorCodeKey, myerror.ErrorCodeJsonFormatInvalid)
		c.Set(logger.LoggerMessageKey, err.Error())
		return
	}

	quota, err := h.quotaUseCase.SyncQuota(context.Background(), req.Provider)
	if err != nil {
		status := http.StatusInternalServerError
		code := myerror.ErrorCodeQuotaSyncFailed
		if usecase.IsRecordNotFound(err) {
			status = http.StatusNotFound
			code = myerror.ErrorCodeProviderGetFailed
		}
		if errors.Is(err, usecase.ErrProviderNotFound) {
			status = http.StatusNotFound
			code = myerror.ErrorCodeProviderGetFailed
		}
		c.JSON(status, quotaResponse{Code: code, Message: err.Error(), Data: nil})
		c.Set(logger.LoggerErrorCodeKey, code)
		c.Set(logger.LoggerMessageKey, err.Error())
		return
	}

	c.JSON(http.StatusOK, quotaResponse{Code: 0, Message: "ok", Data: quota})
}
