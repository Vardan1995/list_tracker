package entity

import "gorm.io/gorm"

type Preference struct {
	gorm.Model
	UserId uint   `gorm:"not null" json:"user_id"`
	Link   string `gorm:"not null" json:"link"`
}
