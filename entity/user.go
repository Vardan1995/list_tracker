package entity

import "gorm.io/gorm"

// Product struct
type User struct {
	gorm.Model
	Username string `gorm:"not null" json:"username"`
	Email    string `gorm:"not null" json:"email"`
}
