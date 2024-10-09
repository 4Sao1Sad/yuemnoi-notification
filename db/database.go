package db

import (
	"fmt"
	"yuemnoi-notification/internal/config"
	"yuemnoi-notification/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgreSQL(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.Db.Host, cfg.Db.Username, cfg.Db.Password, cfg.Db.Database, cfg.Db.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Migration(db)
	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&model.UserDevice{})
}
