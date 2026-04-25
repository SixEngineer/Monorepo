package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"openbridge/backend/internal/domain/entity"
	"strings"
	"time"
)

const baiduQuotaURL = "https://pan.baidu.com/api/quota?checkfree=1"

type BaiduProvider struct {
	client *http.Client
}

type baiduQuotaResponse struct {
	Errno      int    `json:"errno"`
	Errmsg     string `json:"errmsg"`
	Total      int64  `json:"total"`
	Used       int64  `json:"used"`
	RequestID  int64  `json:"request_id"`
	GUIDInfo   int64  `json:"guid_info"`
	Expire     bool   `json:"expire"`
	Free       int64  `json:"free"`
	ActiveType int    `json:"active_type"`
}

func NewBaiduProvider() *BaiduProvider {
	return &BaiduProvider{
		client: &http.Client{Timeout: 10 * time.Second},
	}
}

func (p *BaiduProvider) Name() string {
	return "baidu"
}

// GetQuota 是百度提供商获取配额信息的方法
// 它接收一个上下文和提供商账户信息，返回配额信息和可能的错误
func (p *BaiduProvider) GetQuota(ctx context.Context, account *entity.ProviderAccount) (entity.Quota, error) {
	// 检查账户是否为空
	if account == nil {
		return entity.Quota{}, fmt.Errorf("baidu provider: account is nil")
	}
	// 检查访问令牌是否为空
	if strings.TrimSpace(account.AccessToken) == "" {
		return entity.Quota{}, fmt.Errorf("baidu provider: access token is empty")
	}

	// 创建新的HTTP请求，使用GET方法访问百度配额URL
	baiduQuotaURLWithToken := fmt.Sprintf("%s&access_token=%s", baiduQuotaURL, account.AccessToken)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baiduQuotaURLWithToken, nil)
	if err != nil {
		return entity.Quota{}, fmt.Errorf("baidu provider: create request failed: %w", err)
	}
	// 设置请求头
	req.Header.Set("User-Agent", "OpenBridge/1.0")
	// req.Header.Set("Authorization", "Bearer "+account.AccessToken)

	// 发送HTTP请求
	resp, err := p.client.Do(req)
	if err != nil {
		return entity.Quota{}, fmt.Errorf("baidu provider: request failed: %w", err)
	}
	defer resp.Body.Close() // 确保响应体被关闭

	// 检查响应状态码，处理未授权或禁止访问的情况
	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return entity.Quota{}, fmt.Errorf("baidu provider: token invalid or expired, status=%d", resp.StatusCode)
	}
	// 检查响应状态码是否为200，如果不是则返回错误
	if resp.StatusCode != http.StatusOK {
		return entity.Quota{}, fmt.Errorf("baidu provider: unexpected status=%d", resp.StatusCode)
	}

	// 解析响应体到百度配额响应结构体
	var payload baiduQuotaResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return entity.Quota{}, fmt.Errorf("baidu provider: parse response failed: %w", err)
	}

	// 检查API返回的错误码
	if payload.Errno != 0 {
		// 处理令牌无效或过期的情况
		if payload.Errno == -6 || payload.Errno == 111 {
			return entity.Quota{}, fmt.Errorf("baidu provider: token invalid or expired, errno=%d errmsg=%s", payload.Errno, payload.Errmsg)
		}
		// 处理其他API错误
		return entity.Quota{}, fmt.Errorf("baidu provider: api errno=%d errmsg=%s", payload.Errno, payload.Errmsg)
	}
	// 验证配额字段的合法性
	if payload.Total < 0 || payload.Used < 0 || payload.Used > payload.Total {
		return entity.Quota{}, fmt.Errorf("baidu provider: invalid quota fields total=%d used=%d", payload.Total, payload.Used)
	}

	// 创建当前UTC时间
	now := time.Now().UTC()

	// 转换为 MB
	payload.Total /= (1024 * 1024) 
	payload.Used /= (1024 * 1024)

	// 返回格式化后的配额信息
	return entity.Quota{
		Provider:  "baidu",
		Total:     payload.Total,
		Used:      payload.Used,
		Available: payload.Total - payload.Used,
		UpdatedAt: now,
	}, nil
}

func (p *BaiduProvider) GetDirectLink(ctx context.Context, fileID string, account *entity.ProviderAccount) (string, error) {
	// TODO: 实现获取百度网盘文件的直接链接的逻辑
	return "", fmt.Errorf("baidu provider: direct link not implemented")
}

func (p *BaiduProvider) RefreshToken(ctx context.Context, account *entity.ProviderAccount) error {
	// TODO: 实现百度提供商刷新令牌的逻辑
	return fmt.Errorf("baidu provider: refresh token not implemented")
}
