//go:build linux
// +build linux

package usecase

import (
	"openbridge/backend/internal/domain/interfaces"
	"openbridge/backend/internal/domain/providers"
	"openbridge/backend/internal/repository"
)

// buildMountProviderByNetDiskLinux Linux 特定的 Provider 构建函数
func buildMountProviderByNetDiskLinux(netDisk string, providerRepo *repository.ProviderRepository) interfaces.Provider {
	switch netDisk {
	case "local_linux":
		return providers.NewLocalLinuxProvider(providerRepo)
	default:
		return nil
	}
}
