package repository

import (
	"openbridge/backend/internal/domain/entity"

	"gorm.io/gorm"
)

type ProviderRepository struct {
	db *gorm.DB
}

// 构造函数
func NewProviderRepository(db *gorm.DB) *ProviderRepository {
	return &ProviderRepository{db: db}
}

// 创建 ProviderAccount
func (repo *ProviderRepository) InsertProviderAccount(providerAccount *entity.ProviderAccount) error {
	return repo.db.Create(providerAccount).Error
}

// 删除 ProviderAccount
func (repo *ProviderRepository) DeleteProviderAccount(id uint) error {
	return repo.db.Delete(&entity.ProviderAccount{}, id).Error
}

// 更新 ProviderAccount
func (repo *ProviderRepository) UpdateProviderAccount(providerAccount *entity.ProviderAccount) error {
    return repo.db.Updates(providerAccount).Error
}

// 获取 ProviderAccount
func (repo *ProviderRepository) GetProviderAccount(id uint) (*entity.ProviderAccount, error) {
	var providerAccount entity.ProviderAccount
	err := repo.db.First(&providerAccount, id).Error
	if err != nil {
		return nil, err
	}
	return &providerAccount, nil
}

// 获取所有 ProviderAccount
func (repo *ProviderRepository) ListProviderAccounts() ([]entity.ProviderAccount, error) {
	var providerAccounts []entity.ProviderAccount
	err := repo.db.Find(&providerAccounts).Error
	if err != nil {
		return nil, err
	}
	return providerAccounts, nil
}