package handler

import (
	"net/http"
	"openbridge/backend/internal/pkg/myerror"
	"openbridge/backend/internal/tool"
	"openbridge/backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type StorageHandler struct {
	storageUseCase *usecase.StorageUseCase
}

func NewStorageHandler(storageUseCase *usecase.StorageUseCase) *StorageHandler {
    return &StorageHandler{
        storageUseCase: storageUseCase,
    }
}

// 获取驱动列表
func (h *StorageHandler) GetDrivers(c *gin.Context) {

	// 调用 usecase 获取驱动列表
	drivers, err := h.storageUseCase.GetDrivers()
	if err != nil {
	    c.JSON(http.StatusInternalServerError, tool.HttpResult{Code: myerror.ErrorCodeGetDriversFailed, Message: err.Error()})
		return
	}

	// 返回获取驱动列表成功的结果
	c.JSON(http.StatusOK, tool.HttpResult{Code: myerror.ErrorCodeOK, Message: myerror.SuccessMessage, Data: drivers})
}

// 获取驱动信息
func (h *StorageHandler) GetDriverInfo(c *gin.Context) {
    
	// 从 URL 参数中获取驱动名称
	driverName := c.Query("name")

	// 调用 usecase 获取驱动信息
	driverInfo, err := h.storageUseCase.GetDriverInfo(driverName)
	if err != nil {
	    c.JSON(http.StatusInternalServerError, tool.HttpResult{Code: myerror.ErrorCodeGetDriverInfoFailed, Message: err.Error()})
		return
	}

	// 返回获取驱动信息成功的结果
	c.JSON(http.StatusOK, tool.HttpResult{Code: myerror.ErrorCodeOK, Message: myerror.SuccessMessage, Data: driverInfo})
}