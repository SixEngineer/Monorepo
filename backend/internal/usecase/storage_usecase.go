package usecase

import (
	"encoding/json"
	"io"
	"net/http"
	"openbridge/backend/internal/config"
	"time"
)

type StorageUseCase struct {
	config *config.Config
}

// 这个结构体用于 Get driver names (Admin) 接口的响应解析

// DriverResponse 定义用于解析驱动列表响应的结构体
type DriverResponse struct {
    Code    int      `json:"code"`
    Message string   `json:"message"`
    Data    []string `json:"data"`
}

// 以下四个结构体用于 Get driver info (Admin) 接口的响应解析

type InfoResponse struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    InfoResponseData `json:"data"`
}

// data 字段的结构
type InfoResponseData struct {
    Common     []ConfigField   `json:"common"`
    Additional []ConfigField   `json:"additional"`
    Config     StorageConfig   `json:"config"`
}

// 配置字段（common 和 additional 共用）
type ConfigField struct {
    Name     string `json:"name"`
    Type     string `json:"type"`
    Default  string `json:"default"`
    Options  string `json:"options"`
    Required bool   `json:"required"`
    Help     string `json:"help"`
}

// config 字段的结构
type StorageConfig struct {
    Name          string `json:"name"`
    LocalSort     bool   `json:"local_sort"`
    OnlyProxy     bool   `json:"only_proxy"`
    NoCache       bool   `json:"no_cache"`
    NoUpload      bool   `json:"no_upload"`
    NeedMs        bool   `json:"need_ms"`
    DefaultRoot   string `json:"default_root"`
    Alert         string `json:"alert"`
    OnlyIndices   bool   `json:"only_indices"`
    PreferProxy   bool   `json:"prefer_proxy"`
}

func NewStorageUseCase(config *config.Config) *StorageUseCase {
	return &StorageUseCase{
		config: config,
	}
}

func (s *StorageUseCase) GetDrivers() ([]string, error) {

	client := http.Client{Timeout: 10 * time.Second}

	// 构造HTTP GET请求，目标URL为OpenList的驱动列表接口
	req, err := http.NewRequest("GET", s.config.OpenList.BaseURL+"/api/admin/driver/names", nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头，指定User-Agent，并添加Authorization头以使用Bearer Token进行认证
	req.Header.Set("User-Agent", "OpenBridge/1.0")
	req.Header.Set("Authorization", s.config.OpenList.Token)

	// 发送HTTP请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
    var driverResponse DriverResponse
    if err := json.Unmarshal(body, &driverResponse); err != nil {
        return nil, err
    }

	return driverResponse.Data, nil
}

// GetDriverInfo 获取指定驱动的详细信息
func (s *StorageUseCase) GetDriverInfo(driverName string) (*InfoResponseData, error) {
    
	client := http.Client{Timeout: 10 * time.Second}

	// 构造HTTP GET请求，目标URL为OpenList的驱动列表接口
	req, err := http.NewRequest("GET", s.config.OpenList.BaseURL+"/api/admin/driver/info?driver="+driverName, nil)
	if err != nil {
		return nil, err
	}

	// 设置请求头，指定User-Agent，并添加Authorization头以使用Bearer Token进行认证
	req.Header.Set("User-Agent", "OpenBridge/1.0")
	req.Header.Set("Authorization", s.config.OpenList.Token)

	// 发送HTTP请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON响应
    var InfoResponse InfoResponse
    if err := json.Unmarshal(body, &InfoResponse); err != nil {
        return nil, err
    }

	return &InfoResponse.Data, nil
}