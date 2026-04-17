package usecase

import (
	"context"
	"errors"
	"openbridge/backend/internal/domain/entity"
	"openbridge/backend/internal/domain/interfaces"
	"openbridge/backend/internal/domain/providers"
	"openbridge/backend/internal/repository"
	"openbridge/backend/internal/tool"
	"time"

	"gorm.io/gorm"
)

var ErrProviderNotFound = errors.New("provider not found")

type QuotaUseCase struct {
	providerRepo     *repository.ProviderRepository
	quotaRepo        *repository.QuotaRepository
	providerRegistry *tool.Registry
}

// 构造函数
func NewQuotaUseCase(providerRepo *repository.ProviderRepository, quotaRepo *repository.QuotaRepository, providerRegistry *tool.Registry) *QuotaUseCase {
	return &QuotaUseCase{
		providerRepo:     providerRepo,
		quotaRepo:        quotaRepo,
		providerRegistry: providerRegistry,
	}
}

// 查询数据库中的当前配额，不触发远端调用
func (u *QuotaUseCase) QueryQuota(provider string) (entity.Quota, error) {
	account, err := u.providerRepo.GetProviderAccountByProvider(provider)
	if err != nil {
		return entity.Quota{}, err
	}

	updatedAt := account.UpdatedAt.UTC()
	if account.LastQuotaSyncAt != nil {
		updatedAt = account.LastQuotaSyncAt.UTC()
	}

	return entity.Quota{
		Provider:  provider,
		Total:     account.TotalQuota,
		Used:      account.UsedQuota,
		Available: account.AvailableQuota,
		UpdatedAt: updatedAt,
	}, nil
}

// 从远端同步配额并更新本地数据
func (u *QuotaUseCase) SyncQuota(ctx context.Context, provider string) (entity.Quota, error) {
	account, err := u.providerRepo.GetProviderAccountByProvider(provider)
	if err != nil {
		return entity.Quota{}, err
	}

	providerInstance, err := u.resolveProvider(account)
	if err != nil {
		return entity.Quota{}, err
	}

	remoteQuota, err := providerInstance.GetQuota(ctx, account)
	if err != nil {
		return entity.Quota{}, err
	}

	now := time.Now().UTC()
	if err := u.providerRepo.UpdateProviderQuota(account.ID, remoteQuota.Total, remoteQuota.Used, remoteQuota.Available, now); err != nil {
		return entity.Quota{}, err
	}

	snapshot := &entity.QuotaSnapshot{
		Provider:          provider,
		ProviderAccountID: account.ID,
		Total:             remoteQuota.Total,
		Used:              remoteQuota.Used,
		Available:         remoteQuota.Available,
		SyncedAt:          now,
	}
	if err := u.quotaRepo.InsertQuotaSnapshot(snapshot); err != nil {
		return entity.Quota{}, err
	}

	remoteQuota.Provider = provider
	remoteQuota.UpdatedAt = now
	return remoteQuota, nil
}

// resolveProvider 是一个方法，用于解析和获取提供者(Provider)实例
// 它接受一个 ProviderAccount 类型的指针作为参数，返回一个 Provider 接口和可能的错误
// @param account *entity.ProviderAccount - 提供者账户信息
// @return interfaces.Provider - 提供者接口实例
// @return error - 可能发生的错误
func (u *QuotaUseCase) resolveProvider(account *entity.ProviderAccount) (interfaces.Provider, error) {
	// 首先尝试从提供者注册表中获取已存在的提供者实例
	if providerInstance, ok := u.providerRegistry.Get(account.Name); ok {
		return providerInstance, nil
	}

	// 如果注册表中不存在，则根据网络磁盘类型构建新的提供者实例
	providerInstance := buildProviderByNetDisk(account.NetDisk)
	if providerInstance == nil {
		// 如果构建失败，返回提供者未找到的错误
		return nil, ErrProviderNotFound
	}

	// 将新构建的提供者实例注册到注册表中，以便后续使用
	// 使用下划线忽略返回值，因为我们不需要注册操作的返回结果
	_ = u.providerRegistry.Register(account.Name, providerInstance)
	return providerInstance, nil
}

func buildProviderByNetDisk(netDisk string) interfaces.Provider {
	switch netDisk {
	case "mock":
		return &providers.MockProvider{}
	default:
		return nil
	}
}

// 是否为数据库未找到错误
func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
