package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string       `gorm:"not null" json:"username"`
	Email       string       `gorm:"not null;unique" json:"email"`
	Preferences []Preference `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"` // User has many Preferences
}
