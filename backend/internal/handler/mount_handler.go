package handler

import (
	"context"
	"errors"
	"net/http"
	"openbridge/backend/internal/domain/entity"
	"openbridge/backend/internal/pkg/logger"
	"openbridge/backend/internal/pkg/myerror"
	"openbridge/backend/internal/tool"
	"openbridge/backend/internal/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MountHandler struct {
	mountUseCase *usecase.MountUseCase
}

func NewMountHandler(mountUseCase *usecase.MountUseCase) *MountHandler {
	return &MountHandler{mountUseCase: mountUseCase}
}

func (h *MountHandler) CreateMount(c *gin.Context) {
	var req entity.MountPoint
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, tool.HttpResult{Code: myerror.ErrorCodeJsonFormatInvalid, Message: err.Error()})
		c.Set(logger.LoggerErrorCodeKey, myerror.ErrorCodeJsonFormatInvalid)
		c.Set(logger.LoggerMessageKey, err.Error())
		return
	}

	mount, err := h.mountUseCase.CreateMount(context.Background(), req)
	if err != nil {
		status, code := mapMountError(err)
		c.JSON(status, tool.HttpResult{Code: code, Message: err.Error()})
		c.Set(logger.LoggerErrorCodeKey, code)
		c.Set(logger.LoggerMessageKey, err.Error())
		return
	}

	c.JSON(http.StatusOK, tool.HttpResult{Code: myerror.ErrorCodeOK, Message: myerror.SuccessMessage, Data: mount})
}

func (h *MountHandler) GetMountQuota(c *gin.Context) {
	mountID, err := parseMountID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, tool.HttpResult{Code: myerror.ErrorCodeParameterInvalid, Message: err.Error()})
		c.Set(logger.LoggerErrorCodeKey, myerror.ErrorCodeParameterInvalid)
		c.Set(logger.LoggerMessageKey, err.Error())
		return
	}

	result, err := h.mountUseCase.GetMountQuota(context.Background(), mountID)
	if err != nil {
		status, code := mapMountError(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
			code = myerror.ErrorCodeMountGetFailed
		}
		c.JSON(status, tool.HttpResult{Code: code, Message: err.Error()})
		c.Set(logger.LoggerErrorCodeKey, code)
		c.Set(logger.LoggerMessageKey, err.Error())
		return
	}

	c.JSON(http.StatusOK, tool.HttpResult{Code: myerror.ErrorCodeOK, Message: myerror.SuccessMessage, Data: result})
}

// SyncMountQuota 处理挂载配额同步请求的函数
// 该函数接收一个gin.Context对象，用于处理HTTP请求和响应
// 它会解析挂载ID，调用业务逻辑层进行配额同步，并根据结果返回相应的响应
func (h *MountHandler) SyncMountQuota(c *gin.Context) {
    // 解析请求中的挂载ID，如果解析失败则返回参数错误响应
	mountID, err := parseMountID(c)
	if err != nil {
        // 返回HTTP 400状态码，表示请求参数无效
		c.JSON(http.StatusBadRequest, tool.HttpResult{Code: myerror.ErrorCodeParameterInvalid, Message: err.Error()})
        // 设置日志记录的错误代码和消息
		c.Set(logger.LoggerErrorCodeKey, myerror.ErrorCodeParameterInvalid)
		c.Set(logger.LoggerMessageKey, err.Error())
		return
	}

    // 调用业务逻辑层的SyncMountQuota方法进行配额同步
	result, err := h.mountUseCase.SyncMountQuota(context.Background(), mountID)
	if err != nil {
        // 映射业务错误到HTTP状态码和错误代码
		status, code := mapMountError(err)
        // 如果是记录未找到错误，设置HTTP 404状态码和特定的错误代码
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
			code = myerror.ErrorCodeMountGetFailed
		}
        // 如果是配额解析失败错误，转换为配额同步失败错误
		if code == myerror.ErrorCodeQuotaResolveFailed {
			code = myerror.ErrorCodeMountQuotaSyncFailed
		}
        // 返回相应的错误响应
		c.JSON(status, tool.HttpResult{Code: code, Message: err.Error()})
        // 设置日志记录的错误代码和消息
		c.Set(logger.LoggerErrorCodeKey, code)
		c.Set(logger.LoggerMessageKey, err.Error())
		return
	}

    // 成功响应，返回HTTP 200状态码，成功消息和同步结果
	c.JSON(http.StatusOK, tool.HttpResult{Code: myerror.ErrorCodeOK, Message: myerror.SuccessMessage, Data: result})
}

func parseMountID(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	parsed, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(parsed), nil
}

func mapMountError(err error) (int, int) {
	switch {
	case errors.Is(err, usecase.ErrMountInvalidMode),
		errors.Is(err, usecase.ErrMountProviderRequired),
		errors.Is(err, usecase.ErrMountParentRequired),
		errors.Is(err, usecase.ErrMountParentNotReal),
		errors.Is(err, usecase.ErrMountCircularInherit),
		errors.Is(err, usecase.ErrMountVirtualExceedsAllowed),
		errors.Is(err, usecase.ErrMountVirtualUsedInvalid),
		errors.Is(err, usecase.ErrMountDisabled):
		return http.StatusBadRequest, myerror.ErrorCodeMountValidationFailed
	default:
		return http.StatusInternalServerError, myerror.ErrorCodeQuotaResolveFailed
	}
}
