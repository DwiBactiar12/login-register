package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `gorm:"type:varchar(35);not null"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(16);not null"`
	HP       string `gorm:"type:varchar(35);not null"`
	Role     string `gorm:"type:varchar(35);not null"`
	Umur     int    `gorm:"typeint;not null"`
	Gender   string `gorm:"type:varchar(35);not null"`
	AdminID  uint   `gorm:"type:int;not null"`
}
type Admin struct {
	gorm.Model
	Nama     string `gorm:"type:varchar(35);not null"`
	Email    string `gorm:"type:varchar(100);not null;unique"`
	Password string `gorm:"type:varchar(16);not null"`
	HP       string `gorm:"type:varchar(35);not null"`
	Role     string `gorm:"type:varchar(35);not null"`
	Umur     int    `gorm:"typeint;not null"`
	Gender   string `gorm:"type:varchar(35);not null"`
	User     []User `gorm:"foreignkey:AdminID"`
}
