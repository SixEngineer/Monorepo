package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"openbridge/backend/internal/config"
	"strings"
	"time"
)

type StorageUseCase struct {
	config *config.Config
}

// 1. 以下一个结构体用于 Get driver names (Admin) 接口的响应解析

// 1.0 DriverResponse 定义用于解析驱动列表响应的结构体
type DriverResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

// 2. 以下四个结构体用于 Get driver info (Admin) 接口的响应解析

// 2.0
type InfoResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    InfoResponseData `json:"data"`
}

// 2.1 data 字段的结构
type InfoResponseData struct {
	Common     []ConfigField `json:"common"`
	Additional []ConfigField `json:"additional"`
	Config     StorageConfig `json:"config"`
}

// 2.2 配置字段（common 和 additional 共用）
type ConfigField struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Default  string `json:"default"`
	Options  string `json:"options"`
	Required bool   `json:"required"`
	Help     string `json:"help"`
}

// 2.3 config 字段的结构
type StorageConfig struct {
	Name        string `json:"name"`
	LocalSort   bool   `json:"local_sort"`
	OnlyProxy   bool   `json:"only_proxy"`
	NoCache     bool   `json:"no_cache"`
	NoUpload    bool   `json:"no_upload"`
	NeedMs      bool   `json:"need_ms"`
	DefaultRoot string `json:"default_root"`
	Alert       string `json:"alert"`
	OnlyIndices bool   `json:"only_indices"`
	PreferProxy bool   `json:"prefer_proxy"`
}

// 3. 以下三个结构体用于 Get driver info (Admin) 接口的响应解析

// 3.0 最外层响应结构
type FileListResponse struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    FileListData `json:"data"`
}

// 3.1 data 字段的结构
type FileListData struct {
	Content  []FileItem `json:"content"`
	Total    int        `json:"total"`
	Readme   string     `json:"readme"`
	Header   string     `json:"header"`
	Write    bool       `json:"write"`
	Provider string     `json:"provider"`
}

// 3.2 文件/目录项结构
type FileItem struct {
	Name     string      `json:"name"`
	Size     int64       `json:"size"`
	IsDir    bool        `json:"is_dir"`
	Modified time.Time   `json:"modified"`
	Created  time.Time   `json:"created"`
	Sign     string      `json:"sign"`
	Thumb    string      `json:"thumb"`
	Type     int         `json:"type"`      // 1: 目录, 可能是其他值表示文件
	Hashinfo string      `json:"hashinfo"`  // 注意原始 JSON 中是字符串 "null"
	HashInfo interface{} `json:"hash_info"` // 原始 JSON 中是 null
}

// 4. 以下一个结构体用于 Get file info (Admin) 接口的响应解析

// 4.0 最外层响应结构
type FileInfoResponse struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    FileDetail `json:"data"`
}

// 4.1 data 字段的结构
type FileDetail struct {
	Name     string      `json:"name"`
	Size     int64       `json:"size"`
	IsDir    bool        `json:"is_dir"`
	Modified string      `json:"modified"` // 或 time.Time
	Created  string      `json:"created"`  // 或 time.Time
	Sign     string      `json:"sign"`
	Thumb    string      `json:"thumb"`
	Type     int         `json:"type"`
	Hashinfo string      `json:"hashinfo"`
	HashInfo interface{} `json:"hash_info"`
	RawURL   string      `json:"raw_url"`
	Readme   string      `json:"readme"`
	Header   string      `json:"header"`
	Provider string      `json:"provider"`
	Related  interface{} `json:"related"`
}

type DirectLinkResult struct {
	Path            string `json:"path"`
	Name            string `json:"name"`
	Size            int64  `json:"size"`
	Provider        string `json:"provider"`
	DirectLink      string `json:"direct_link"`
	IsOpenListProxy bool   `json:"is_openlist_proxy"`
}

func NewStorageUseCase(config *config.Config) *StorageUseCase {
	return &StorageUseCase{
		config: config,
	}
}

// GetDrivers 获取所有驱动名称
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

// GetFiles 获取指定目录下的文件列表
func (s *StorageUseCase) GetFiles(path string, page uint, pageSize uint) (*FileListData, error) {

	client := http.Client{Timeout: 10 * time.Second}

	// 准备请求体，包含路径和分页信息
	requestBody := map[string]interface{}{
		"path":     path,
		"page":     page,
		"per_page": pageSize,
	}

	// 将请求体编码为JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// 构造HTTP POST请求，目标URL为OpenList的文件列表接口
	req, err := http.NewRequest("POST", s.config.OpenList.BaseURL+"/api/fs/list", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	// 设置请求头，指定User-Agent，并添加Authorization头以使用Bearer Token进行认证
	req.Header.Set("User-Agent", "OpenBridge/1.0")
	req.Header.Set("Authorization", s.config.OpenList.Token)
	req.Header.Set("Content-Type", "application/json")

	// 发送HTTP请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	var fileListResponse FileListResponse
	if err := json.NewDecoder(resp.Body).Decode(&fileListResponse); err != nil {
		return nil, err
	}

	return &fileListResponse.Data, nil

}

// GetFileInfo 获取指定文件的信息
func (s *StorageUseCase) GetFileInfo(path string) (*FileDetail, error) {

	client := http.Client{Timeout: 10 * time.Second}

	// 准备请求体，包含路径和分页信息
	requestBody := map[string]interface{}{
		"path": path,
	}

	// 将请求体编码为JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	// 构造HTTP POST请求，目标URL为OpenList的文件列表接口
	req, err := http.NewRequest("POST", s.config.OpenList.BaseURL+"/api/fs/get", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	// 设置请求头，指定User-Agent，并添加Authorization头以使用Bearer Token进行认证
	req.Header.Set("User-Agent", "OpenBridge/1.0")
	req.Header.Set("Authorization", s.config.OpenList.Token)
	req.Header.Set("Content-Type", "application/json")

	// 发送HTTP请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	var fileInfoResponse FileInfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&fileInfoResponse); err != nil {
		return nil, err
	}

	return &fileInfoResponse.Data, nil
}

func (s *StorageUseCase) ResolveDirectLink(path string) (*DirectLinkResult, error) {
	detail, err := s.GetFileInfo(path)
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(detail.RawURL) == "" {
		return nil, errors.New("raw_url empty")
	}

	rawURL, err := url.Parse(detail.RawURL)
	if err != nil {
		return nil, err
	}

	baseURL, err := url.Parse(s.config.OpenList.BaseURL)
	if err != nil {
		return nil, err
	}
	if baseURL.Host == "" && baseURL.Path != "" {
		baseURL, err = url.Parse("http://" + s.config.OpenList.BaseURL)
		if err != nil {
			return nil, err
		}
	}

	isProxy := false
	if rawURL.Host != "" && baseURL.Host != "" {
		isProxy = strings.EqualFold(rawURL.Host, baseURL.Host)
	}

	return &DirectLinkResult{
		Path:            path,
		Name:            detail.Name,
		Size:            detail.Size,
		Provider:        detail.Provider,
		DirectLink:      detail.RawURL,
		IsOpenListProxy: isProxy,
	}, nil
}
