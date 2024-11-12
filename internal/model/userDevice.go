package model

import (
	"time"
)

type UserDevice struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UserID      uint      `gorm:"not null;index:idx_user_device,unique"`           // Add to unique index
	DeviceToken string    `gorm:"type:text;not null;index:idx_user_device,unique"` // Add to unique index
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

func (UserDevice) TableName() string {
	return "user_devices"
}
