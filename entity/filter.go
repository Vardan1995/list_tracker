package entity

import "gorm.io/gorm"

type Filter struct {
	gorm.Model
	UserId uint   `gorm:"not null" json:"user_id"`
	Link   string `gorm:"not null" json:"link"`
}
