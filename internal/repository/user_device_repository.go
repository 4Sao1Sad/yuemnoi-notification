package repository

import (
	"yuemnoi-notification/internal/model"

	"gorm.io/gorm"
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
	if err := i.db.Create(&device).Error; err != nil {
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
