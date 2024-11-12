package repository

import (
	"yuemnoi-notification/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserDeviceRepository interface {
	CreateUserDevice(device model.UserDevice) error
	GetUserDevices(userId int) ([]model.UserDevice, error)
}

type UserDeviceRepositoryImpl struct {
	db *gorm.DB
}

func NewUserDeviceRepository(db *gorm.DB) UserDeviceRepository {
	return &UserDeviceRepositoryImpl{db}
}

func (i UserDeviceRepositoryImpl) CreateUserDevice(device model.UserDevice) error {
	if err := i.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "device_token"}}, // Specify columns for conflict
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),           // Only update the UpdatedAt field on conflict
	}).Create(&device).Error; err != nil {
		return err
	}
	return nil
}

func (i UserDeviceRepositoryImpl) GetUserDevices(userId int) ([]model.UserDevice, error) {
	var userDevices []model.UserDevice
	if err := i.db.Where("user_id = ?", userId).Find(&userDevices).Error; err != nil {
		return nil, err
	}

	return userDevices, nil
}
