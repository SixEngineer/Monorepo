package repository

import (
	"openbridge/backend/internal/domain/entity"

	"gorm.io/gorm"
)

type MountRepository struct {
	db *gorm.DB
}

func NewMountRepository(db *gorm.DB) *MountRepository {
	return &MountRepository{db: db}
}

func (repo *MountRepository) InsertMountPoint(mountPoint *entity.MountPoint) error {
	return repo.db.Create(mountPoint).Error
}

func (repo *MountRepository) GetMountPoint(id uint) (*entity.MountPoint, error) {
	var mount entity.MountPoint
	if err := repo.db.First(&mount, id).Error; err != nil {
		return nil, err
	}
	return &mount, nil
}
