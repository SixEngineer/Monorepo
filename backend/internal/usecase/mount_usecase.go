package usecase

import (
	"context"
	"errors"
	"fmt"
	"openbridge/backend/internal/domain/entity"
	"openbridge/backend/internal/domain/interfaces"
	"openbridge/backend/internal/domain/providers"
	"openbridge/backend/internal/pkg/logger"
	"openbridge/backend/internal/repository"
	"openbridge/backend/internal/tool"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ErrMountInvalidMode           = errors.New("mount quota_mode is invalid")
	ErrMountProviderRequired      = errors.New("real mode requires provider_account_id")
	ErrMountParentRequired        = errors.New("inherit mode requires inherit_from_id")
	ErrMountParentNotReal         = errors.New("inherit parent must be real mode")
	ErrMountCircularInherit       = errors.New("inherit chain has cycle")
	ErrMountVirtualExceedsAllowed = errors.New("virtual_total exceeds allowed_max")
	ErrMountVirtualUsedInvalid    = errors.New("virtual_used must be <= virtual_total")
	ErrMountDisabled              = errors.New("mount is disabled")
)

type MountQuotaResult struct {
	MountID       uint             `json:"mount_id"`
	Mode          string           `json:"mode"`
	AllowedMax    int64            `json:"allowed_max"`
	Quota         entity.Quota     `json:"quota"`
	InheritChain  []uint           `json:"inherit_chain,omitempty"`
	VirtualConfig map[string]int64 `json:"virtual_config,omitempty"`
}

type MountUseCase struct {
	mountRepo        *repository.MountRepository
	providerRepo     *repository.ProviderRepository
	quotaRepo        *repository.QuotaRepository
	providerRegistry *tool.Registry
}

func NewMountUseCase(mountRepo *repository.MountRepository, providerRepo *repository.ProviderRepository, quotaRepo *repository.QuotaRepository, providerRegistry *tool.Registry) *MountUseCase {
	return &MountUseCase{
		mountRepo:        mountRepo,
		providerRepo:     providerRepo,
		quotaRepo:        quotaRepo,
		providerRegistry: providerRegistry,
	}
}

func (u *MountUseCase) CreateMount(ctx context.Context, mount entity.MountPoint) (*entity.MountPoint, error) {
	mode := entity.QuotaMode(strings.ToLower(strings.TrimSpace(mount.QuotaMode)))
	mount.QuotaMode = string(mode)

	if err := u.validateMountConfig(ctx, &mount, mode); err != nil {
		return nil, err
	}

	if mode == entity.QuotaModeVirtual && mount.VirtualTotal == 0 {
		mount.ReadOnly = true
	}

	if err := u.mountRepo.InsertMountPoint(&mount); err != nil {
		return nil, err
	}
	return &mount, nil
}

func (u *MountUseCase) GetMountQuota(ctx context.Context, mountID uint) (MountQuotaResult, error) {
	return u.resolveMountQuota(ctx, mountID, false)
}

func (u *MountUseCase) SyncMountQuota(ctx context.Context, mountID uint) (MountQuotaResult, error) {
	return u.resolveMountQuota(ctx, mountID, true)
}

// resolveMountQuota 是一个用于解析挂载点配额的方法
// 它接收上下文、挂载点ID和是否同步远程的标志作为参数
// 返回解析结果和可能的错误
func (u *MountUseCase) resolveMountQuota(ctx context.Context, mountID uint, syncRemote bool) (MountQuotaResult, error) {
	// 从仓库获取挂载点信息
	mount, err := u.mountRepo.GetMountPoint(mountID)
	if err != nil {
		return MountQuotaResult{}, err
	}
	// 检查挂载点是否启用，未启用则返回错误
	if !mount.Enabled {
		return MountQuotaResult{}, ErrMountDisabled
	}

	// 根据模式解析配额，使用map防止循环引用
	result, err := u.resolveByMode(ctx, mount, syncRemote, map[uint]struct{}{})
	if err != nil {
		// 记录解析失败的日志
		logger.L().Error("mount quota resolve failed",
			zap.Uint("mount_id", mount.ID),
			zap.String("mode", mount.QuotaMode),
			zap.Error(err),
		)
		now := time.Now().UTC()
		// 插入配额快照记录失败状态
		_ = u.quotaRepo.InsertQuotaSnapshot(&entity.QuotaSnapshot{
			MountPointID:      mount.ID,
			ProviderAccountID: mount.ProviderAccountID,
			Provider:          mount.ProviderType,
			Mode:              mount.QuotaMode,
			SyncStatus:        "failed",
			ErrorMessage:      err.Error(),
			SyncedAt:          now,
		})
		return MountQuotaResult{}, err
	}

	now := time.Now().UTC()
	// 插入成功的配额快照记录
	if err := u.quotaRepo.InsertQuotaSnapshot(&entity.QuotaSnapshot{
		MountPointID:      mount.ID,
		ProviderAccountID: mount.ProviderAccountID,
		Provider:          mount.ProviderType,
		Mode:              mount.QuotaMode,
		Total:             result.Quota.Total,
		Used:              result.Quota.Used,
		Available:         result.Quota.Available,
		SyncStatus:        "success",
		SyncedAt:          now,
	}); err != nil {
		return MountQuotaResult{}, err
	}

	// 更新配额的更新时间并记录成功日志
	result.Quota.UpdatedAt = now
	logger.L().Info("mount quota resolved",
		zap.Uint("mount_id", mount.ID),
		zap.String("mode", mount.QuotaMode),
		zap.Int64("total", result.Quota.Total),
		zap.Int64("used", result.Quota.Used),
		zap.Int64("available", result.Quota.Available),
	)
	return result, nil
}

func (u *MountUseCase) resolveByMode(ctx context.Context, mount *entity.MountPoint, syncRemote bool, visited map[uint]struct{}) (MountQuotaResult, error) {
	if _, ok := visited[mount.ID]; ok {
		return MountQuotaResult{}, ErrMountCircularInherit
	}
	visited[mount.ID] = struct{}{}

	mode := entity.QuotaMode(strings.ToLower(mount.QuotaMode))
	switch mode {
	case entity.QuotaModeReal:
		quota, err := u.resolveRealQuota(ctx, mount, syncRemote)
		if err != nil {
			return MountQuotaResult{}, err
		}
		logger.L().Info("quota resolver mode",
			zap.String("mode", string(mode)),
			zap.Uint("mount_id", mount.ID),
		)
		return MountQuotaResult{
			MountID:    mount.ID,
			Mode:       mount.QuotaMode,
			AllowedMax: quota.Total,
			Quota:      quota,
		}, nil
	case entity.QuotaModeInherit:
		if mount.InheritFromID == nil {
			return MountQuotaResult{}, ErrMountParentRequired
		}
		parent, err := u.mountRepo.GetMountPoint(*mount.InheritFromID)
		if err != nil {
			return MountQuotaResult{}, err
		}
		if strings.ToLower(parent.QuotaMode) != string(entity.QuotaModeReal) {
			return MountQuotaResult{}, ErrMountParentNotReal
		}
		parentResult, err := u.resolveByMode(ctx, parent, syncRemote, visited)
		if err != nil {
			return MountQuotaResult{}, err
		}
		chain := []uint{parent.ID}
		logger.L().Info("quota resolver inherit chain",
			zap.Uint("mount_id", mount.ID),
			zap.Uint("parent_mount_id", parent.ID),
		)
		return MountQuotaResult{
			MountID:      mount.ID,
			Mode:         mount.QuotaMode,
			AllowedMax:   parentResult.AllowedMax,
			Quota:        parentResult.Quota,
			InheritChain: chain,
		}, nil
	case entity.QuotaModeVirtual:
		allowedMax, err := u.getAllowedMax(ctx, mount, syncRemote)
		if err != nil {
			return MountQuotaResult{}, err
		}
		if mount.VirtualUsed > mount.VirtualTotal {
			return MountQuotaResult{}, ErrMountVirtualUsedInvalid
		}
		if mount.VirtualTotal > allowedMax {
			logger.L().Warn("virtual quota validation failed",
				zap.Uint("mount_id", mount.ID),
				zap.Int64("virtual_total", mount.VirtualTotal),
				zap.Int64("allowed_max", allowedMax),
			)
			return MountQuotaResult{}, ErrMountVirtualExceedsAllowed
		}
		logger.L().Info("quota resolver mode",
			zap.String("mode", string(mode)),
			zap.Uint("mount_id", mount.ID),
			zap.Int64("virtual_total", mount.VirtualTotal),
			zap.Int64("virtual_used", mount.VirtualUsed),
			zap.Int64("allowed_max", allowedMax),
		)
		return MountQuotaResult{
			MountID:    mount.ID,
			Mode:       mount.QuotaMode,
			AllowedMax: allowedMax,
			Quota: entity.Quota{
				Provider:  mount.ProviderType,
				Total:     mount.VirtualTotal,
				Used:      mount.VirtualUsed,
				Available: mount.VirtualTotal - mount.VirtualUsed,
			},
			VirtualConfig: map[string]int64{
				"virtual_total": mount.VirtualTotal,
				"virtual_used":  mount.VirtualUsed,
			},
		}, nil
	default:
		return MountQuotaResult{}, ErrMountInvalidMode
	}
}

func (u *MountUseCase) resolveRealQuota(ctx context.Context, mount *entity.MountPoint, syncRemote bool) (entity.Quota, error) {
	if mount.ProviderAccountID == 0 {
		return entity.Quota{}, ErrMountProviderRequired
	}
	account, err := u.providerRepo.GetProviderAccount(mount.ProviderAccountID)
	if err != nil {
		return entity.Quota{}, err
	}

	if !syncRemote {
		if account.TotalQuota < account.UsedQuota || account.TotalQuota < 0 || account.UsedQuota < 0 {
			return entity.Quota{}, fmt.Errorf("invalid stored quota in provider account")
		}
		return entity.Quota{
			Provider:  mount.ProviderType,
			Total:     account.TotalQuota,
			Used:      account.UsedQuota,
			Available: account.AvailableQuota,
			UpdatedAt: account.UpdatedAt.UTC(),
		}, nil
	}

	providerInstance, err := u.resolveProvider(account)
	if err != nil {
		return entity.Quota{}, err
	}
	remoteQuota, err := providerInstance.GetQuota(ctx, account)
	if err != nil {
		return entity.Quota{}, err
	}
	logger.L().Info("provider quota fetched",
		zap.String("provider", account.NetDisk),
		zap.Uint("provider_account_id", account.ID),
		zap.Int64("total", remoteQuota.Total),
		zap.Int64("used", remoteQuota.Used),
		zap.Int64("available", remoteQuota.Available),
	)

	now := time.Now().UTC()
	if err := u.providerRepo.UpdateProviderQuota(account.ID, remoteQuota.Total, remoteQuota.Used, remoteQuota.Available, now); err != nil {
		return entity.Quota{}, err
	}
	remoteQuota.Provider = mount.ProviderType
	remoteQuota.UpdatedAt = now
	return remoteQuota, nil
}

func (u *MountUseCase) getAllowedMax(ctx context.Context, mount *entity.MountPoint, syncRemote bool) (int64, error) {
	account, err := u.providerRepo.GetProviderAccount(mount.ProviderAccountID)
	if err != nil {
		return 0, err
	}
	if !syncRemote && account.TotalQuota > 0 {
		return account.TotalQuota, nil
	}

	providerInstance, err := u.resolveProvider(account)
	if err != nil {
		return 0, err
	}
	quota, err := providerInstance.GetQuota(ctx, account)
	if err != nil {
		return 0, err
	}
	now := time.Now().UTC()
	if err := u.providerRepo.UpdateProviderQuota(account.ID, quota.Total, quota.Used, quota.Available, now); err != nil {
		return 0, err
	}
	logger.L().Info("virtual allowed_max evaluated",
		zap.Uint("mount_id", mount.ID),
		zap.Int64("allowed_max", quota.Total),
	)
	return quota.Total, nil
}

func (u *MountUseCase) validateMountConfig(ctx context.Context, mount *entity.MountPoint, mode entity.QuotaMode) error {
	switch mode {
	case entity.QuotaModeReal:
		if mount.ProviderAccountID == 0 {
			return ErrMountProviderRequired
		}
		account, err := u.providerRepo.GetProviderAccount(mount.ProviderAccountID)
		if err != nil {
			return err
		}
		mount.ProviderType = account.NetDisk
		mount.InheritFromID = nil
		mount.VirtualTotal = 0
		mount.VirtualUsed = 0
		return nil
	case entity.QuotaModeInherit:
		if mount.InheritFromID == nil {
			return ErrMountParentRequired
		}
		parent, err := u.mountRepo.GetMountPoint(*mount.InheritFromID)
		if err != nil {
			return err
		}
		if strings.ToLower(parent.QuotaMode) != string(entity.QuotaModeReal) {
			return ErrMountParentNotReal
		}
		if err := u.validateNoCycle(parent.ID, mount.ID); err != nil {
			return err
		}
		mount.ProviderAccountID = parent.ProviderAccountID
		mount.ProviderType = parent.ProviderType
		mount.VirtualTotal = 0
		mount.VirtualUsed = 0
		return nil
	case entity.QuotaModeVirtual:
		if mount.ProviderAccountID == 0 {
			return ErrMountProviderRequired
		}
		if mount.VirtualUsed > mount.VirtualTotal {
			return ErrMountVirtualUsedInvalid
		}
		account, err := u.providerRepo.GetProviderAccount(mount.ProviderAccountID)
		if err != nil {
			return err
		}
		mount.ProviderType = account.NetDisk
		allowedMax, err := u.getAllowedMax(ctx, mount, true)
		if err != nil {
			return err
		}
		if mount.VirtualTotal > allowedMax {
			return ErrMountVirtualExceedsAllowed
		}
		return nil
	default:
		return ErrMountInvalidMode
	}
}

func (u *MountUseCase) validateNoCycle(startID uint, candidateID uint) error {
	visited := map[uint]struct{}{}
	currentID := startID
	for currentID != 0 {
		if currentID == candidateID {
			return ErrMountCircularInherit
		}
		if _, ok := visited[currentID]; ok {
			return ErrMountCircularInherit
		}
		visited[currentID] = struct{}{}
		current, err := u.mountRepo.GetMountPoint(currentID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return err
		}
		if current.InheritFromID == nil {
			return nil
		}
		currentID = *current.InheritFromID
	}
	return nil
}

func (u *MountUseCase) resolveProvider(account *entity.ProviderAccount) (interfaces.Provider, error) {
	if providerInstance, ok := u.providerRegistry.Get(account.Name); ok {
		return providerInstance, nil
	}

	providerInstance := buildMountProviderByNetDisk(account.NetDisk)
	if providerInstance == nil {
		return nil, ErrProviderNotFound
	}
	_ = u.providerRegistry.Register(account.Name, providerInstance)
	return providerInstance, nil
}

func buildMountProviderByNetDisk(netDisk string) interfaces.Provider {
	switch netDisk {
	case "mock":
		return &providers.MockProvider{}
	case "baidu":
		return providers.NewBaiduProvider()
	default:
		return nil
	}
}
