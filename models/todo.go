package models

import "gorm.io/gorm"

type Todo struct {
    ID        uint   `gorm:"primaryKey"`
    Title     string `gorm:"type:varchar(255);not null"`
    Completed bool   `gorm:"default:false"`
}

func Migrate(db *gorm.DB) {
    db.AutoMigrate(&Todo{})
}
