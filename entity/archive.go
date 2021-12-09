package entity

import "gorm.io/gorm"

type Archive struct {
	gorm.Model
	Email     string `gorm:"not null" json:"email"`
	ProductId string `gorm:"not null" json:"product_id"`
}
