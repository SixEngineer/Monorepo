//go:build linux
// +build linux

package usecase

import (
	"errors"
	"openbridge/backend/internal/domain/entity"
	"openbridge/backend/internal/domain/providers"
)

// RegisterProviderLinux 注册 Linux 特定的 Provider
func (p *ProviderUseCase) RegisterProviderLinux(providerAccount entity.ProviderAccount) error {
	providerNetDisk := providerAccount.NetDisk

	switch providerNetDisk {
	case "local_linux":
		p.ProviderRegistry.Register(providerAccount.Name, providers.NewLocalLinuxProvider(p.ProviderRepo))
	default:
		return errors.New("provider netdisk type undefined")
	}

	return p.ProviderRepo.InsertProviderAccount(&providerAccount)
}

// UpdateProviderLinux 更新 Linux 特定的 Provider
func (p *ProviderUseCase) UpdateProviderLinux(providerAccount entity.ProviderAccount) error {
	providerAccountOld, err := p.ProviderRepo.GetProviderAccount(providerAccount.ID)
	if err != nil {
		return err
	}

	// 删除旧的Provider
	p.ProviderRegistry.Unregister(providerAccountOld.Name)

	// 注册新的Provider
	providerNetDisk := providerAccount.NetDisk

	switch providerNetDisk {
	case "local_linux":
		p.ProviderRegistry.Register(providerAccount.Name, providers.NewLocalLinuxProvider(p.ProviderRepo))
	default:
		return errors.New("provider not found")
	}

	return p.ProviderRepo.UpdateProviderAccount(&providerAccount)
}
